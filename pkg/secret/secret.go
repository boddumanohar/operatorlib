package secret

import (
	"reflect"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
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

	data, err = mergeData(data, stringData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to merge string data and data")
	}

	s = &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: *om,
		Data:       data,
		Type:       corev1.SecretType(c.Type),
	}
	return s, nil
}

// MaybeUpdate implements MaybeUpdateFunc for Secret object. It
// compares the two Secrets being passed and update the first one if
// required.
func MaybeUpdate(original interfaces.Object, new interfaces.Object) (bool, error) {
	os, ok := original.(*corev1.Secret)
	if !ok {
		return false, errors.New("failed to assert the original object")
	}

	ns, ok := new.(*corev1.Secret)
	if !ok {
		return false, errors.New("failed to assert the new object")
	}

	result := reflect.DeepEqual(os.Data, ns.Data)
	if result {
		return false, nil
	}

	os.Data = ns.Data

	return true, nil
}

// Create generates Secret as per the `Conf` struct passed and creates
// it in the cluster
func Create(c Conf) (reconcile.Result, error) {
	var s *corev1.Secret
	var err error
	if c.GenSecretFunc != nil {
		s, err = c.GenSecretFunc(c)
	} else {
		s, err = GenerateSecret(c)
	}
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

// Update generates the Secret as per the `Conf` struct passed and
// compares it with the in-cluster version. If required, it updates
// the in-cluster Secret with the changes. For comparing the Secrets,
// it uses `MaybeUpdate` function by default but can also use
// `MaybeUpdateFunc` from `Conf` if passed.
func Update(c Conf) (reconcile.Result, error) {
	var s *corev1.Secret
	var err error
	if c.GenSecretFunc != nil {
		s, err = c.GenSecretFunc(c)
	} else {
		s, err = GenerateSecret(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate secret")
	}

	var maybeUpdateFunc operation.MaybeUpdateFunc
	if c.MaybeUpdateFunc != nil {
		maybeUpdateFunc = c.MaybeUpdateFunc
	} else {
		maybeUpdateFunc = MaybeUpdate
	}

	result, err := operation.Update(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          s,
		ExistingObject:  &corev1.Secret{},
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to update secret")
	}

	return result, nil
}

// CreateOrUpdate is a combination of `Create` and `Update`
// functions. It creates the Secret object if it is not already in the
// cluster and updates the Secret if one exists.
func CreateOrUpdate(c Conf) (reconcile.Result, error) {
	var s *corev1.Secret
	var err error
	if c.GenSecretFunc != nil {
		s, err = c.GenSecretFunc(c)
	} else {
		s, err = GenerateSecret(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate secret")
	}

	var maybeUpdateFunc operation.MaybeUpdateFunc
	if c.MaybeUpdateFunc != nil {
		maybeUpdateFunc = c.MaybeUpdateFunc
	} else {
		maybeUpdateFunc = MaybeUpdate
	}

	result, err := operation.CreateOrUpdate(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          s,
		ExistingObject:  &corev1.Secret{},
		OwnerReference:  c.OwnerReference,
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create or update secret")
	}

	return result, nil
}

// Delete generates the ObjectMeta for Secret as per the `Conf` struct
// passed and deletes it from the cluster
func Delete(c Conf) (reconcile.Result, error) {
	om, err := meta.GenerateObjectMeta(meta.Conf{
		Instance:           c.Instance,
		Name:               c.Name,
		Namespace:          c.Namespace,
		GenLabelsFunc:      c.GenLabelsFunc,
		GenAnnotationsFunc: c.GenAnnotationsFunc,
		GenFinalizersFunc:  c.GenFinalizersFunc,
		AppendLabels:       c.AppendLabels,
	})
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate objectmeta for secret")
	}

	result, err := operation.Delete(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          &corev1.Secret{ObjectMeta: *om},
		AfterDeleteFunc: c.AfterDeleteFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to delete secret")
	}

	return result, nil
}
