package operation

import (
	"argp.in/go/operatorlib/pkg/interfaces"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MaybeUpdateFunc is the update function type which is used for
// updating objects. The function gives the user flexibility to define
// there own update functions and pass them. The function is supposed
// to update the first argument and return true if update is required
// or else false. The function can return error if one occurs during
// comparision.
type MaybeUpdateFunc func(interfaces.Object, interfaces.Object) (bool, error)

// HookFunc is the function type which is used by various hooks. The
// function of this type can be used to hook custom logic in
// pre-defined operations.
type HookFunc func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error)

// AfterCreateFunc is the hook called after creating the object.
type AfterCreateFunc HookFunc

// AfterUpdateFunc is the hook called after updating the object.
type AfterUpdateFunc HookFunc

// AfterDeleteFunc is the hook called after deleting the object.
type AfterDeleteFunc HookFunc

// Conf is the struct used by all Operation functions. This can be
// used to pass various parameters which can be used by the functions.
type Conf struct {
	// Instance is the pointer to the owner object which is performing
	// operation on the Object.
	Instance interfaces.Object
	// Reconcile is the pointer to reconcile struct of the owner
	// Object.
	Reconcile interfaces.Reconcile
	// Object is the pointer to the object on which operation is to be
	// performed.
	Object interfaces.Object
	// ExistingObject is the pointer to the empty struct of Object
	// which is used to fetch the existing object from cluster
	// TODO: Figure out a way to avoid this
	ExistingObject interfaces.Object
	// OwnerReference is the flag used to by Create operation to
	// determine if owner reference needs to be set on the created
	// object.
	OwnerReference bool
	// MaybeUpdateFunc is used by Update operation to determine if
	// Update is required and also update the object
	MaybeUpdateFunc
	// AfterCreateFunc hook is called after creating the Object
	AfterCreateFunc
	// AfterUpdateFunc hook is called after updating the Object
	AfterUpdateFunc
	// AfterDeleteFunc hook is called after deleting the Object
	AfterDeleteFunc
}
