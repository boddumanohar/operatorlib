package configmap

import (
	"github.com/ankitrgadiya/operatorlib/pkg/meta"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateConfigMap generates ConfigMap object as per the `Conf` struct passed
func GenerateConfigMap(c Conf) (cm *corev1.ConfigMap, err error) {
	var om *metav1.ObjectMeta
	var data map[string]string
	var binData map[string][]byte

	om, err = meta.GenerateObjectMeta(meta.Conf{
		Instance:           c.Instance,
		Name:               c.Name,
		Namespace:          c.Namespace,
		GenLabelsFunc:      c.GenLabelsFunc,
		GenAnnotationsFunc: c.GenAnnotationsFunc,
		AppendLabels:       c.AppendLabels,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate objectmeta")
	}

	if c.GenDataFunc != nil {
		data, err = c.GenDataFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate configmap data")
		}
	}

	if c.GenBinaryDataFunc != nil {
		binData, err = c.GenBinaryDataFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate configmap binary data")
		}
	}

	cm = &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: *om,
		Data:       data,
		BinaryData: binData,
	}

	return cm, err
}
