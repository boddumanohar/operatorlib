package configmap

import (
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/meta"
	"github.com/ankitrgadiya/operatorlib/pkg/operation"
)

// GenDataFunc defines a function which generates data (string map)
// for Configmap.
type GenDataFunc func(interfaces.Object) (map[string]string, error)

// GenBinaryDataFunc defines a function which generates binary data
// (map of string to byte slice) for Configmap.
type GenBinaryDataFunc func(interfaces.Object) (map[string][]byte, error)

// Conf is used to pass parameters to functions in this package to
// perform operations on Configmap objects.
type Conf struct {
	// Instance is the Owner object which manages the Configmap
	Instance interfaces.Object
	// Reconcile is the pointer to reconcile struct of owner object
	interfaces.Reconcile
	// Name of the Configmap
	Name string
	// Namespace of the Configmap
	Namespace string
	// GenLalebsFunc is used to generate labels for ObjectMeta
	meta.GenLabelsFunc
	// GenAnnotationsFunc is used to generate annotations for ObjectMeta
	meta.GenAnnotationsFunc
	// GenFinalizers is used to generate finalizers for ObjectMeta
	meta.GenFinalizersFunc
	// AppendLabels is used to determine if labels from Owner object
	// are to be inherited
	AppendLabels bool
	// OwnerReference is used to determine if owner reference needs to
	// be set on Configmap before creating it in cluster
	OwnerReference bool
	// MaybeUpdateFunc defines an update function with custom logic
	// for Configmap update
	operation.MaybeUpdateFunc
	// AfterCreateFunc hook is called after creating the Configmap
	operation.AfterCreateFunc
	// AfterUpdateFunc hook is called after updating the Configmap
	operation.AfterUpdateFunc
	// AfterDeleteFunc hook is called after deleting the Configmap
	operation.AfterDeleteFunc
	// GenDataFunc defines a function to generate data for Configmap
	GenDataFunc
	// GenBinaryDataFunc defines a function to generate binary data
	// for Configmap
	GenBinaryDataFunc
}
