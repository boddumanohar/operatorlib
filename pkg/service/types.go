package service

import (
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/meta"
	"github.com/ankitrgadiya/operatorlib/pkg/operation"

	corev1 "k8s.io/api/core/v1"
)

// GenServiceFunc defines a function which generates Service
type GenServiceFunc func(Conf) (*corev1.Service, error)

// GenServicePortsFunc defines a function which generates slice of
// ServicePort for Service
type GenServicePortsFunc func(interfaces.Object) ([]corev1.ServicePort, error)

// GenSelectorFunc defines a function which generates selectors for
// the Service
type GenSelectorFunc func(interfaces.Object) (map[string]string, error)

// Conf is used to pass parameters to functions in this package to
// perform operations on Service objects.
type Conf struct {
	// Instance is the Owner object which manages the Service
	Instance interfaces.Object
	// Reconcile is the pointer to reconcile struct of owner object
	interfaces.Reconcile
	// Name of the Service
	Name string
	// Namespace of the Service
	Namespace string
	// GenLabelsFunc is used to generate labels for ObjectMeta
	meta.GenLabelsFunc
	// GenAnnotationsFunc is used to generate annotations for
	// ObjectMeta
	meta.GenAnnotationsFunc
	// GenFinalizers is used to generate finalizers for ObjectMeta
	meta.GenFinalizersFunc
	// AppendLabels is used to determine if labels from Owner object
	// are to be inherited
	AppendLabels bool
	// OwnerReference is used to determine if owner reference needs to
	// be set on Service before creating it in cluster
	OwnerReference bool
	// MaybeUpdateFunc defines an update function with custom logic
	// for Service update
	operation.MaybeUpdateFunc
	// AfterCreateFunc hook is called after creating the Service
	operation.AfterCreateFunc
	// AfterUpdateFunc hook is called after updating the Service
	operation.AfterUpdateFunc
	// AfterDeleteFunc hook is called after deleting the Service
	operation.AfterDeleteFunc
	// GenServicePortsFunc defines a function to generate ports for
	// the Service
	GenServicePortsFunc
	// GenSelectorFunc defines a function to generate selector for the
	// Service
	GenSelectorFunc
	// Type defines the type of Service object to be created
	Type string
}
