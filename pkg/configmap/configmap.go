package configmap

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

// MaybeUpdate implements MaybeUpdateFunc for Configmap object. It
// compares the two Configmaps being passed and update the first one
// if required.
func MaybeUpdate(original interfaces.Object, new interfaces.Object) (bool, error) {
	ocm, ok := original.(*corev1.ConfigMap)
	if !ok {
		return false, errors.New("failed to assert the original object")
	}

	ncm, ok := new.(*corev1.ConfigMap)
	if !ok {
		return false, errors.New("failed to assert the existing object")
	}

	result := reflect.DeepEqual(ocm.Data, ncm.Data) && reflect.DeepEqual(ocm.BinaryData, ncm.BinaryData)
	if result {
		return false, nil
	}

	ocm.Data = ncm.Data
	ocm.BinaryData = ncm.BinaryData

	return true, nil
}

// Create generates the ConfigMap as per the `Conf` struct passed and
// creates it in the cluster
func Create(c Conf) (reconcile.Result, error) {
	var cm *corev1.ConfigMap
	var err error
	if c.GenConfigMapFunc != nil {
		cm, err = c.GenConfigMapFunc(c)
	} else {
		cm, err = GenerateConfigMap(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate configmap")
	}

	result, err := operation.Create(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          cm,
		OwnerReference:  c.OwnerReference,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create configmap")
	}

	return result, nil
}

// Update generates the ConfigMap as per the `Conf` struct passed and
// compares it with the in-cluster version. If required, it updates
// the in-cluster ConfigMap with the changes. For comparing the
// ConfigMaps, it uses `MaybeUpdate` function by default but can also
// use `MaybeUpdateFunc` from `Conf` if passed.
func Update(c Conf) (reconcile.Result, error) {
	var cm *corev1.ConfigMap
	var err error
	if c.GenConfigMapFunc != nil {
		cm, err = c.GenConfigMapFunc(c)
	} else {
		cm, err = GenerateConfigMap(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate configmap")
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
		Object:          cm,
		ExistingObject:  &corev1.ConfigMap{},
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to update configmap")
	}

	return result, nil
}

// CreateOrUpdate is a combination of `Create` and `Update`
// functions. It creates the ConfigMap object if it is not already in
// the cluster and updates the ConfigMap if one exists.
func CreateOrUpdate(c Conf) (reconcile.Result, error) {
	var cm *corev1.ConfigMap
	var err error
	if c.GenConfigMapFunc != nil {
		cm, err = c.GenConfigMapFunc(c)
	} else {
		cm, err = GenerateConfigMap(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate configmap")
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
		Object:          cm,
		ExistingObject:  &corev1.ConfigMap{},
		OwnerReference:  c.OwnerReference,
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create or update configmap")
	}

	return result, nil
}

// Delete generates the ObjectMeta for ConfigMap as per the `Conf`
// struct passed and deletes it from the cluster
func Delete(c Conf) (reconcile.Result, error) {
	om, err := meta.GenerateObjectMeta(meta.Conf{
		Instance:           c.Instance,
		Name:               c.Name,
		Namespace:          c.Namespace,
		GenLabelsFunc:      c.GenLabelsFunc,
		GenAnnotationsFunc: c.GenAnnotationsFunc,
		AppendLabels:       c.AppendLabels,
	})
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate objectmeta for configmap")
	}

	result, err := operation.Delete(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          &corev1.ConfigMap{ObjectMeta: *om},
		AfterDeleteFunc: c.AfterDeleteFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to delete configmap")
	}

	return result, nil
}
