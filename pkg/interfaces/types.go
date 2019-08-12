package interfaces

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Object is the interface which all Kubernetes objects
// implements. This interface can be used to pass around any
// Kubernetes Object. This helps keep the functions more generic and
// less tied to the specific Objects.
type Object interface {
	// The object needs to implement Meta Object interface from API
	// machinery. This interface is used for various Client operations
	// on Kubernetes objects.
	metav1.Object
	// The object needs to implement Runtime Object interface from API
	// machinery.
	runtime.Object
}

// Reconcile is the interface for Reconcile object structs . This
// interface can be used to pass around Reconcile structs commonly
// used in Operators.
//
// Note however that by default Reconcile structs generated using
// Operator SDK do not implement this interface. Add following
// functions to implement this interface.
//
//     func (r *ReconcileObject) GetClient() client.Client { return r.client }
//     func (r *ReconcileObject) GetScheme() *runtime.Scheme { return r.scheme }
//
// The Reconcile object structs must implement this interface to use
// Operatorlib functions.
type Reconcile interface {
	// Getter function for reconcile client
	GetClient() client.Client
	// Getter function for reconcile Scheme
	GetScheme() *runtime.Scheme
}
