package secret

import (
	"github.com/ankitrgadiya/operatorlib/pkg/meta"
	"github.com/ankitrgadiya/operatorlib/pkg/operation"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// GenerateSecret generates Secret object as per the `Conf` struct
// passed. However, this does one special thing while generating
// Secret. StringData is merged into Data because of how Secrets are
// handled by Kubernetes. API Server never returns StringData field
// and converts it into Data. One might thing that is fine since API
// Server is already doing the conversion but the way this library
// manages Update will break if the generated object, Secret do not
// match the one in cluster. Merging the StringData ensures that if no
// genuine change is made then generated Secret will match the one in
// cluster.
func GenerateSecret(c Conf) (s *corev1.Secret, err error) {
	var om *metav1.ObjectMeta
	var data map[string][]byte
	var stringData map[string]string

	om, err = meta.GenerateObjectMeta(meta.Conf{
		Instance:           c.Instance,
		Name:               c.Name,
		Namespace:          c.Namespace,
		GenLabelsFunc:      c.GenLabelsFunc,
		GenAnnotationsFunc: c.GenAnnotationsFunc,
		GenFinalizersFunc:  c.GenFinalizersFunc,
		AppendLabels:       c.AppendLabels,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate objectmeta")
	}

	if c.GenDataFunc != nil {
		data, err = c.GenDataFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate data")
		}
	}

	if c.GenStringDataFunc != nil {
		stringData, err = c.GenStringDataFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate string data")
		}
	}

	err = mergeData(data, stringData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to merge string data and data")
	}

	s = &corev1.Secret{
		ObjectMeta: *om,
		Data:       data,
		StringData: stringData,
		Type:       corev1.SecretType(c.Type),
	}
	return s, nil
}

// Create generates Secret as per the `Conf` struct passed and creates
// it in the cluster
func Create(c Conf) (reconcile.Result, error) {
	s, err := GenerateSecret(c)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate secret")
	}

	result, err := operation.Create(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          s,
		OwnerReference:  c.OwnerReference,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create secret")
	}
	return result, nil
}
