package secret

import (
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/meta"
	"github.com/ankitrgadiya/operatorlib/pkg/operation"
)

// GenDataFunc defines a function which generates map of string to
// byte slice for `Data` field in Secret object
type GenDataFunc func(interfaces.Object) (map[string][]byte, error)

// GenStringDataFunc defines a function which generates string map for
// `StringData` field in Secret object
type GenStringDataFunc func(interfaces.Object) (map[string]string, error)

// Conf is used to pass parameters to functions in this package to
// perform operations on Secret objects.
type Conf struct {
	// Instance is the Owner object which manages the Secret
	Instance interfaces.Object
	// Reconcile is the pointer to reconcile struct of owner object
	interfaces.Reconcile
	// Name of the Secret
	Name string
	// Namespace of the Secret
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
	// be set on Secret before creating it in cluster
	OwnerReference bool
	// MaybeUpdateFunc defines an update function with custom logic
	// for Secret update
	operation.MaybeUpdateFunc
	// AfterCreateFunc hook is called after creating the Secret
	operation.AfterCreateFunc
	// AfterUpdateFunc hook is called after updating the Secret
	operation.AfterUpdateFunc
	// AfterDeleteFunc hook is called after deleting the Secret
	operation.AfterDeleteFunc
	// GenDataFunc defines a function to generate data for Secret
	GenDataFunc
	// GenBinaryDataFunc defines a function to generate binary data
	// for Secret
	GenStringDataFunc
	// Type defines the type of Secret
	Type string
}
