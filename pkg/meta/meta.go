package meta

import (
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateObjectMeta function generates ObjectMeta struct as per the
// Conf struct passed.
func GenerateObjectMeta(c Conf) (om *metav1.ObjectMeta, err error) {
	var labels, annotations map[string]string
	var finalizers []string

	if c.GenLabelsFunc != nil {
		labels, err = c.GenLabelsFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate labels")
		}
	}

	if c.AppendLabels {
		err = mergo.Merge(&labels, c.Instance.GetLabels())
		if err != nil {
			return nil, errors.Wrap(err, "failed to append labels from owner object")
		}
	}

	if c.GenAnnotationsFunc != nil {
		annotations, err = c.GenAnnotationsFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate annotations")
		}
	}

	if c.GenFinalizersFunc != nil {
		finalizers, err = c.GenFinalizersFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate finalizers")
		}
	}

	om = &metav1.ObjectMeta{
		Name:        c.Name,
		Namespace:   c.Namespace,
		Labels:      labels,
		Annotations: annotations,
		Finalizers:  finalizers,
	}

	return om, nil
}
