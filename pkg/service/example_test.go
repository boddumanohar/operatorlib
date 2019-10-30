package service_test

import (
	"log"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/service"

	corev1 "k8s.io/api/core/v1"
)

var ownerObject interfaces.Object
var ownerReconcile interfaces.Reconcile

func ExampleCreate() {
	result, err := service.Create(service.Conf{
		// Instance is the pointer to owner object under which
		// Service is being created.
		Instance: ownerObject,
		// OwnerReference can be used to tell if owner reference is
		// required to set on the Service object.
		OwnerReference: true,
		// Reconcile is the reconcile struct of the owner object which
		// implements the interfaces.Reconcile struct. For more
		// details check Reconcile interface documentation.
		Reconcile: ownerReconcile,
		// Name is the name of generated Service. There are several
		// options defines in service.Conf which can be used to
		// manipulate ObjectMeta of the generated object.
		Name: "cm-test",
		Type: "ClusterIP",
		// GenServicePortsFunc is the function that generates list of
		// service ports for the Service. This can be anonymous
		// function like this or can be defined somewhere else and
		// just pass the name of the function here. Check service.Conf
		// struct for other such funtions.
		GenServicePortsFunc: func(interfaces.Object) ([]corev1.ServicePort, error) {
			return []corev1.ServicePort{{Port: int32(80)}}, nil
		},
	})
	if err != nil {
		log.Fatal(result, err)
	}
}
