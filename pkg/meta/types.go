package meta

import (
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
)

// GenLabelsFunc defines a function which takes the Owner Object and
// generates map of string to string. This can be used to generate
// labels for the Object.
type GenLabelsFunc func(interfaces.Object) (map[string]string, error)

// GenAnnotationsFunc defines a function which takes the Owner Object
// and generates map of string to string. This can be used to generate
// Annotations for the Object.
type GenAnnotationsFunc func(interfaces.Object) (map[string]string, error)

// GenFinalizersFunc defines a function which takes the Owner Object
// and generates string slice This can be used to generate Finalizers
// for the Object.
type GenFinalizersFunc func(interfaces.Object) ([]string, error)

// Conf object is used to pass parameters to functions in Meta package
// and generate ObjectMeta struct.
type Conf struct {
	// Instance is the Owner Object which manages the Object.
	Instance interfaces.Object
	// Name defines the name for the Object. This is used to populate
	// the field of same name in ObjectMeta
	Name string
	// Namespace defines the namespace in which Object is/will be
	// present. This is used to populate the field of same name in
	// ObjectMeta
	Namespace string
	// GenLabelsFunc is used to generate labels for Object. The
	// generated string map populates the Labels in ObjectMeta.
	GenLabelsFunc
	// GenAnnotationsFunc is used to generate annotations for the
	// Object. The generated string map populates the Annotations in
	// ObjectMeta
	GenAnnotationsFunc
	// GenFinalizersFunc is used to generate finalizers for the
	// Object. The generated string slice populates the Finalizers in
	// ObjectMeta
	GenFinalizersFunc
}
