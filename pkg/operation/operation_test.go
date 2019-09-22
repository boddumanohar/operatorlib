package operation_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces/mocks"
	"github.com/ankitrgadiya/operatorlib/pkg/operation"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func mockSetup(ctrl *gomock.Controller) (i *mocks.MockObject, r *mocks.MockReconcile) {
	i = mocks.NewMockObject(ctrl)
	i.EXPECT().GetName().Return("test").AnyTimes()
	i.EXPECT().GetUID().Return(types.UID("199bd7a8-b72a-4411-b55e-91096769e58f")).AnyTimes()

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}
	r = mocks.NewMockReconcile(ctrl)

	c := fake.NewFakeClient([]runtime.Object{cm}...)
	s := scheme.Scheme
	s.AddKnownTypes(schema.GroupVersion{Group: "test", Version: "v1"}, i)

	r.EXPECT().GetClient().Return(c).AnyTimes()
	r.EXPECT().GetScheme().Return(s).AnyTimes()

	return i, r
}

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("bad parameters", func(t *testing.T) {
		t.Run("empty instance and reconciler", func(t *testing.T) {
			// instance, reconcile := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Create(operation.Conf{}) }, "")
		})
		t.Run("empty instance", func(t *testing.T) {
			_, r := mockSetup(controller)
			assert.NotPanics(t, func() { _, _ = operation.Create(operation.Conf{Reconcile: r}) }, "")
		})
		t.Run("empty reconcile", func(t *testing.T) {
			i, _ := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Create(operation.Conf{Instance: i}) }, "")
		})
	})
	t.Run("mock instance", func(t *testing.T) {
		t.Run("create empty configmap", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"}}

			_, err := operation.Create(operation.Conf{Reconcile: r, Instance: i, Object: object})
			assert.NoError(t, err)

			createdObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, createdObject)
			assert.NoError(t, err)

			assert.Equal(t, object, createdObject)
		})
		t.Run("create configmap with data", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			_, err := operation.Create(operation.Conf{Reconcile: r, Instance: i, Object: object})
			assert.NoError(t, err)

			createdObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, createdObject)
			assert.NoError(t, err)

			assert.Equal(t, object, createdObject)
		})
		t.Run("create configmap with owner reference", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			_, err := operation.Create(operation.Conf{Reconcile: r, Instance: i, Object: object, OwnerReference: true})
			assert.NoError(t, err)

			createdObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, createdObject)
			assert.NoError(t, err)

			assert.Equal(t, object, createdObject)
		})
		t.Run("create configmap and after create function", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			_, err := operation.Create(operation.Conf{
				Reconcile: r,
				Instance:  i,
				Object:    object,
				AfterCreateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, nil
				},
			})
			assert.NoError(t, err)

			createdObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, createdObject)
			assert.NoError(t, err)

			assert.Equal(t, object, createdObject)
		})
		t.Run("create configmap and after create function fails", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			_, err := operation.Create(operation.Conf{
				Reconcile: r,
				Instance:  i,
				Object:    object,
				AfterCreateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, errors.New("test error")
				},
			})
			assert.Error(t, err)

			createdObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, createdObject)
			assert.NoError(t, err)

			assert.Equal(t, object, createdObject)
		})
		t.Run("create configmap fails when configmap already exists", func(t *testing.T) {
			i, r := mockSetup(controller)

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			}

			_, err := operation.Create(operation.Conf{Reconcile: r, Instance: i, Object: object})
			assert.Error(t, err)
		})
	})
}

func TestCreateOrUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("bad parameters", func(t *testing.T) {
		t.Run("empty instance and reconciler", func(t *testing.T) {
			assert.Panics(t, func() { _, _ = operation.CreateOrUpdate(operation.Conf{}) }, "")
		})
		t.Run("empty instance", func(t *testing.T) {
			_, r := mockSetup(controller)
			assert.NotPanics(t, func() { _, _ = operation.CreateOrUpdate(operation.Conf{Reconcile: r}) }, "")
		})
		t.Run("empty reconcile", func(t *testing.T) {
			i, _ := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.CreateOrUpdate(operation.Conf{Instance: i}) }, "")
		})
	})
	t.Run("mock instance", func(t *testing.T) {
		t.Run("create configmap", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"}}

			_, err := operation.CreateOrUpdate(operation.Conf{Instance: i, Reconcile: r, Object: object, ExistingObject: &corev1.ConfigMap{}})
			assert.NoError(t, err)

			result := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, result)
			assert.NoError(t, err)
			assert.Equal(t, object, result)
		})
		t.Run("update configmap", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "new-value1",
					"key2": "new-value2",
				},
			}

			_, err := operation.CreateOrUpdate(operation.Conf{
				Instance:       i,
				Reconcile:      r,
				Object:         object,
				ExistingObject: &corev1.ConfigMap{},
				MaybeUpdateFunc: func(e interfaces.Object, o interfaces.Object) (bool, error) {
					ecm := e.(*corev1.ConfigMap)
					ocm := o.(*corev1.ConfigMap)

					ecm.Data = ocm.Data

					return true, nil
				},
			})
			assert.NoError(t, err)

			result := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-existing-configmap", Namespace: "test"}, result)
			assert.NoError(t, err)
			assert.Equal(t, object.Data, result.Data)
		})
	})
}

func TestUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("bad parameters", func(t *testing.T) {
		t.Run("empty instance and reconciler", func(t *testing.T) {
			assert.Panics(t, func() { _, _ = operation.Update(operation.Conf{}) }, "")
		})
		t.Run("empty instance", func(t *testing.T) {
			_, r := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Update(operation.Conf{Reconcile: r}) }, "")
		})
		t.Run("empty reconcile", func(t *testing.T) {
			i, _ := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Update(operation.Conf{Instance: i}) }, "")
		})
		t.Run("empty object", func(t *testing.T) {
			i, r := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Update(operation.Conf{Instance: i, Reconcile: r}) }, "")
		})
	})
	t.Run("mock instance", func(t *testing.T) {
		t.Run("update configmap which does not exist", func(t *testing.T) {
			i, r := mockSetup(controller)
			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"}}

			_, err := operation.Update(operation.Conf{
				Instance:       i,
				Reconcile:      r,
				Object:         object,
				ExistingObject: &corev1.ConfigMap{},
			})
			assert.Error(t, err)
		})
		t.Run("update configmap maybeupdate function returns error", func(t *testing.T) {
			i, r := mockSetup(controller)
			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Update(operation.Conf{
				Instance:        i,
				Reconcile:       r,
				Object:          object,
				ExistingObject:  &corev1.ConfigMap{},
				MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return false, errors.New("test error") },
			})
			assert.Error(t, err)
		})
		t.Run("update configmap which is up-to-date", func(t *testing.T) {
			i, r := mockSetup(controller)
			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Update(operation.Conf{
				Instance:        i,
				Reconcile:       r,
				Object:          object,
				ExistingObject:  &corev1.ConfigMap{},
				MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return false, nil },
			})
			assert.NoError(t, err)
		})
		t.Run("update configmap which is up-to-date with after update function fails", func(t *testing.T) {
			i, r := mockSetup(controller)
			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Update(operation.Conf{
				Instance:        i,
				Reconcile:       r,
				Object:          object,
				ExistingObject:  &corev1.ConfigMap{},
				MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return false, nil },
				AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, errors.New("test error")
				},
			})
			assert.Error(t, err)
		})
		t.Run("update configmap which is up-to-date with after updatei function func", func(t *testing.T) {
			i, r := mockSetup(controller)
			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Update(operation.Conf{
				Instance:        i,
				Reconcile:       r,
				Object:          object,
				ExistingObject:  &corev1.ConfigMap{},
				MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return false, nil },
				AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, nil
				},
			})
			assert.NoError(t, err)
		})
		t.Run("update configmap data", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "new-value1",
					"key2": "new-value2",
				},
			}

			_, err := operation.Update(operation.Conf{
				Instance:       i,
				Reconcile:      r,
				Object:         object,
				ExistingObject: &corev1.ConfigMap{},
				MaybeUpdateFunc: func(e interfaces.Object, o interfaces.Object) (bool, error) {
					ecm := e.(*corev1.ConfigMap)
					ocm := o.(*corev1.ConfigMap)

					ecm.Data = ocm.Data

					return true, nil
				},
			})
			assert.NoError(t, err)

			result := &corev1.ConfigMap{}
			err = client.Get(context.Background(), types.NamespacedName{Name: "test-existing-configmap", Namespace: "test"}, result)
			assert.NoError(t, err)
			assert.Equal(t, result.Data, object.Data)
		})
		t.Run("update configmap data with after update function fails", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "new-value1",
					"key2": "new-value2",
				},
			}

			_, err := operation.Update(operation.Conf{
				Instance:       i,
				Reconcile:      r,
				Object:         object,
				ExistingObject: &corev1.ConfigMap{},
				MaybeUpdateFunc: func(e interfaces.Object, o interfaces.Object) (bool, error) {
					ecm := e.(*corev1.ConfigMap)
					ocm := o.(*corev1.ConfigMap)

					ecm.Data = ocm.Data

					return true, nil
				},
				AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, errors.New("test error")
				},
			})
			assert.Error(t, err)

			result := &corev1.ConfigMap{}
			err = client.Get(context.Background(), types.NamespacedName{Name: "test-existing-configmap", Namespace: "test"}, result)
			assert.NoError(t, err)
			assert.Equal(t, result.Data, object.Data)
		})
		t.Run("update configmap data with after update function", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"},
				Data: map[string]string{
					"key1": "new-value1",
					"key2": "new-value2",
				},
			}

			_, err := operation.Update(operation.Conf{
				Instance:       i,
				Reconcile:      r,
				Object:         object,
				ExistingObject: &corev1.ConfigMap{},
				MaybeUpdateFunc: func(e interfaces.Object, o interfaces.Object) (bool, error) {
					ecm := e.(*corev1.ConfigMap)
					ocm := o.(*corev1.ConfigMap)

					ecm.Data = ocm.Data

					return true, nil
				},
				AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, nil
				},
			})
			assert.NoError(t, err)

			result := &corev1.ConfigMap{}
			err = client.Get(context.Background(), types.NamespacedName{Name: "test-existing-configmap", Namespace: "test"}, result)
			assert.NoError(t, err)
			assert.Equal(t, result.Data, object.Data)
		})
	})
}

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("bad parameters", func(t *testing.T) {
		t.Run("empty instance and reconciler", func(t *testing.T) {
			assert.Panics(t, func() { _, _ = operation.Delete(operation.Conf{}) }, "")
		})
		t.Run("empty instance", func(t *testing.T) {
			_, r := mockSetup(controller)
			assert.NotPanics(t, func() { _, _ = operation.Delete(operation.Conf{Reconcile: r}) }, "")
		})
		t.Run("empty reconcile", func(t *testing.T) {
			i, _ := mockSetup(controller)
			assert.Panics(t, func() { _, _ = operation.Delete(operation.Conf{Instance: i}) }, "")
		})
	})
	t.Run("mock instance", func(t *testing.T) {
		t.Run("delete existing configmap", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Delete(operation.Conf{Instance: i, Reconcile: r, Object: object})
			assert.NoError(t, err)

			deletedObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, deletedObject)
			assert.Error(t, err)
		})
		t.Run("delete existing configmap that do not exist", func(t *testing.T) {
			i, r := mockSetup(controller)

			object := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Name: "test-configmap", Namespace: "test"},
			}

			_, err := operation.Delete(operation.Conf{Instance: i, Reconcile: r, Object: object})
			assert.NoError(t, err)
		})
		t.Run("delete configmap and after delete function", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Delete(operation.Conf{
				Instance:  i,
				Reconcile: r,
				Object:    object,
				AfterDeleteFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, nil
				},
			})
			assert.NoError(t, err)

			deletedObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, deletedObject)
			assert.Error(t, err)
		})
		t.Run("delete configmap and after delete function fails", func(t *testing.T) {
			i, r := mockSetup(controller)
			client := r.GetClient()

			object := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: "test-existing-configmap", Namespace: "test"}}

			_, err := operation.Delete(operation.Conf{
				Instance:  i,
				Reconcile: r,
				Object:    object,
				AfterDeleteFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
					return reconcile.Result{}, errors.New("test error")
				},
			})
			assert.Error(t, err)

			deletedObject := &corev1.ConfigMap{}
			err = client.Get(context.TODO(), types.NamespacedName{Name: "test-configmap", Namespace: "test"}, deletedObject)
			assert.Error(t, err)
		})
	})
}
