package configmap_test

import (
	"log"

	"github.com/ankitrgadiya/operatorlib/pkg/configmap"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
)

var ownerObject interfaces.Object
var ownerReconcile interfaces.Reconcile

func ExampleCreate() {
	result, err := configmap.Create(configmap.Conf{
		// Instance is the pointer to owner object under which
		// Configmap is being created.
		Instance: ownerObject,
		// OwnerReference can be used to tell if owner reference is
		// required to set on the configmap object.
		OwnerReference: true,
		// Reconcile is the reconcile struct of the owner object which
		// implements the interfaces.Reconcile struct. For more
		// details check Reconcile interface documentation.
		Reconcile: ownerReconcile,
		// Name is the name of generated Configmap. There are several
		// options defines in configmap.Conf which can be used to
		// manipulate ObjectMeta of the generated object.
		Name: "cm-test",
		// GenDataFunc is the function that generates Data to be put
		// into Configmap. This can be anonymous function like this or
		// can be defined somewhere else and just pass the name of the
		// function here. The configmap.Conf struct can also accept
		// GenBinaryDataFunc which generates map of string to byte
		// slice.
		GenDataFunc: func(interfaces.Object) (map[string]string, error) {
			return map[string]string{"key": "value"}, nil
		},
	})
	if err != nil {
		log.Fatal(result, err)
	}
}
