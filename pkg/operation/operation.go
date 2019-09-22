package operation

import (
	"context"
	"reflect"

	"argp.in/go/operatorlib/pkg/interfaces"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Create is a generic create function for any Kubernetes Object. This
// function can create the Object defined in Conf. This function is a
// lower level function and is supposed to be used by other Kubernetes
// Objects. However, this can also be used to create Custom Objects
// (or unsupported objects).
func Create(c Conf) (reconcile.Result, error) {
	return create(c)
}

func create(c Conf) (r reconcile.Result, err error) {
	client := c.Reconcile.GetClient()

	if c.OwnerReference {
		err = controllerutil.SetControllerReference(c.Instance, c.Object, c.Reconcile.GetScheme())
		if err != nil {
			return reconcile.Result{}, errors.Wrapf(err, "failed to set owner reference on the object")
		}
	}

	err = client.Create(context.TODO(), c.Object)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to create the object in cluster")
	}

	if c.AfterCreateFunc != nil {
		r, err = c.AfterCreateFunc(c.Instance, c.Reconcile)
		if err != nil {
			return r, errors.Wrap(err, "failed to run AfterCreate hook")
		}
	}

	return reconcile.Result{}, nil
}

// Update is a generic update function for any Kubernetes Object. This
// can be used to update the objects using MaybeUpdateFunc defined in
// Conf. This is a lower-level function which is supposed to be used
// by other Objects. This can also be used to implement custom update
// logic on any Kubernetes Object. It fetches the Object with same
// metadata as the specified object from the cluster and update them
// using MaybeUpdateFunc.
func Update(c Conf) (reconcile.Result, error) {
	return update(c)
}

func update(c Conf) (r reconcile.Result, err error) {
	client := c.Reconcile.GetClient()

	existingObject, ok := reflect.New(reflect.TypeOf(c.Object)).Interface().(interfaces.Object)
	if !ok {
		return reconcile.Result{}, errors.New("failed to create new instance of the object type")
	}

	err = client.Get(context.TODO(), types.NamespacedName{}, existingObject)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to get the existing object from cluster")
	}

	requireUpdate, err := c.MaybeUpdateFunc(existingObject, c.Object)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to update the object")
	}

	if requireUpdate {
		err = client.Update(context.TODO(), existingObject)
		if err != nil {
			return reconcile.Result{}, errors.Wrap(err, "failed to update the object in cluster")
		}
	}

	if c.AfterUpdateFunc != nil {
		r, err = c.AfterUpdateFunc(c.Instance, c.Reconcile)
		if err != nil {
			return r, errors.Wrap(err, "failed to run AfterUpdate hook")
		}
	}

	return reconcile.Result{}, nil
}

// CreateOrUpdate is the combination of Create and Update. It can be
// used to create any Kubernetes Object. It catches the
// "IsAlreadyExists" error and tries to update the object.
func CreateOrUpdate(c Conf) (r reconcile.Result, err error) {
	r, err = create(c)
	if err != nil && !kerrors.IsAlreadyExists(errors.Cause(err)) {
		return r, errors.Wrap(err, "adsadA")
	}

	if kerrors.IsAlreadyExists(errors.Cause(err)) {
		return update(c)
	}

	return r, nil
}

// Delete is a generic delete function for any Kubernetes Object. This
// function can be used to delete any Kubernetes Object defined in
// Conf. This is a lower-level function which is supposed to be used
// by other Kubernetes Objects. This can also be used to delete any
// Custom Objects (or unsupported Objects).
func Delete(c Conf) (r reconcile.Result, err error) {
	return delete(c)
}

func delete(c Conf) (r reconcile.Result, err error) {
	client := c.Reconcile.GetClient()

	err = client.Delete(context.TODO(), c.Object)
	if err != nil && !kerrors.IsNotFound(err) {
		return reconcile.Result{}, errors.Wrap(err, "failed to delete the object in cluster")
	}

	if c.AfterDeleteFunc != nil {
		r, err = c.AfterDeleteFunc(c.Instance, c.Reconcile)
		if err != nil {
			return r, errors.Wrap(err, "failed to run AfterDelete hook")
		}
	}

	return reconcile.Result{}, nil
}
