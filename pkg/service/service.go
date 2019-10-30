package service

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

// GenerateService generates Service object as per the `Conf` struct
// passed. The function only supports a subset of all the Service
// features but should be good enough for most usecases. Support for
// ExternalService or LoadBalancer service is not there.
func GenerateService(c Conf) (s *corev1.Service, err error) {
	var om *metav1.ObjectMeta
	var ports []corev1.ServicePort
	var selectors map[string]string

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

	if c.GenServicePortsFunc != nil {
		ports, err = c.GenServicePortsFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate service ports")
		}
	}

	if c.GenSelectorFunc != nil {
		selectors, err = c.GenSelectorFunc(c.Instance)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate service selectors")
		}
	}

	s = &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: *om,
		Spec: corev1.ServiceSpec{
			Ports:    ports,
			Selector: selectors,
			Type:     corev1.ServiceType(c.Type),
		},
	}

	return s, nil
}

// MaybeUpdate is the implementation of operation.MaybeUpdateFunc for
// Service object. It compares two service objects and update the
// first one if required. Note however that this does not compare both
// services exhaustively since some fields can also be filled by the
// API Server. If those are also compared here that everytime this
// function is called it will always update/remove those fields. Also,
// service type is immutable and cannot be updated so it returns error
// if that is detected.
func MaybeUpdate(original interfaces.Object, new interfaces.Object) (bool, error) {
	os, ok := original.(*corev1.Service)
	if !ok {
		return false, errors.New("failed to assert the original object")
	}

	ns, ok := new.(*corev1.Service)
	if !ok {
		return false, errors.New("failed to assert the new object")
	}

	// Service Type is immutable field and so it cannot be
	// updated. Return error if it is different.
	if os.Spec.Type != ns.Spec.Type {
		return false, errors.New("type field of service object is different, hoever, it is immutable field which cannot be changed")
	}

	equal := func() bool {
		if len(os.Spec.Ports) != len(ns.Spec.Ports) {
			return false
		}

		for i := 0; i < len(os.Spec.Ports); i++ {
			// Compare Name and Port of the ServicePort Object
			if os.Spec.Ports[i].Name != ns.Spec.Ports[i].Name ||
				os.Spec.Ports[i].Port != ns.Spec.Ports[i].Port {
				return false
			}
		}

		return true
	}()

	// Check if Ports and Selectors are equal
	if equal && reflect.DeepEqual(os.Spec.Selector, ns.Spec.Selector) {
		return false, nil
	}

	// Update Selectors and Ports of the existing service
	os.Spec.Selector = ns.Spec.Selector
	os.Spec.Ports = ns.Spec.Ports

	return true, nil
}

// Create generates the Service as per the `Conf` struct passed and
// creates it in the cluster
func Create(c Conf) (reconcile.Result, error) {
	var s *corev1.Service
	var err error
	if c.GenServiceFunc != nil {
		s, err = c.GenServiceFunc(c)
	} else {
		s, err = GenerateService(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate service")
	}

	result, err := operation.Create(operation.Conf{
		Instance:        c.Instance,
		Reconcile:       c.Reconcile,
		Object:          s,
		OwnerReference:  c.OwnerReference,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create configmap")
	}
	return result, nil
}

// Update generates the Service as per the `Conf` struct passed and compares it with the in-cluster version. If required, it updates the in-cluster Service with the changes.
func Update(c Conf) (reconcile.Result, error) {
	var s *corev1.Service
	var err error
	if c.GenServiceFunc != nil {
		s, err = c.GenServiceFunc(c)
	} else {
		s, err = GenerateService(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate service")
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
		ExistingObject:  &corev1.Service{},
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to update service")
	}

	return result, nil
}

// CreateOrUpdate is a combination of `Create` and `Update`
// functions. It creates the Service object if it is not already in
// the cluster and updates the Service if one exists.
func CreateOrUpdate(c Conf) (reconcile.Result, error) {
	var s *corev1.Service
	var err error
	if c.GenServiceFunc != nil {
		s, err = c.GenServiceFunc(c)
	} else {
		s, err = GenerateService(c)
	}
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to generate service")
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
		ExistingObject:  &corev1.Service{},
		OwnerReference:  c.OwnerReference,
		MaybeUpdateFunc: maybeUpdateFunc,
		AfterUpdateFunc: c.AfterUpdateFunc,
		AfterCreateFunc: c.AfterCreateFunc,
	})
	if err != nil {
		return result, errors.Wrap(err, "failed to create or update service")
	}

	return result, nil
}
