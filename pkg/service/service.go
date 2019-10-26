package service

import (
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
