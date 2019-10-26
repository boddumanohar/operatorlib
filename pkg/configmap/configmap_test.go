package configmap_test

import (
	"errors"
	"testing"

	"github.com/ankitrgadiya/operatorlib/pkg/configmap"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces/mocks"

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

func TestGenerateConfigMap(t *testing.T) {
	t.Run("generate empty configmap", func(t *testing.T) {
		expected := &corev1.ConfigMap{TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		}}

		result, err := configmap.GenerateConfigMap(configmap.Conf{})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("failed to generate objectmeta", func(t *testing.T) {
		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("generate configmap with data", func(t *testing.T) {
		expected := &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			Data: map[string]string{"key": "value"},
		}

		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenDataFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key": "value"}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("failed to generate data", func(t *testing.T) {
		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenDataFunc: func(interfaces.Object) (map[string]string, error) {
				return nil, errors.New("test error")
			},
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("generate configmap with binary data", func(t *testing.T) {
		expected := &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			BinaryData: map[string][]byte{"key": []byte("value")},
		}

		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenBinaryDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return map[string][]byte{"key": []byte("value")}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("failed to generate binary data", func(t *testing.T) {
		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenBinaryDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return nil, errors.New("test error")
			},
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("generate configmap with data and binary data", func(t *testing.T) {
		expected := &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			Data:       map[string]string{"key": "string"},
			BinaryData: map[string][]byte{"key": []byte("bytes")},
		}

		result, err := configmap.GenerateConfigMap(configmap.Conf{
			GenDataFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key": "string"}, nil
			},
			GenBinaryDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return map[string][]byte{"key": []byte("bytes")}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func TestMaybeUpdate(t *testing.T) {
	t.Run("bad parameters", func(t *testing.T) {
		t.Run("bad existing object", func(t *testing.T) {
			result, err := configmap.MaybeUpdate(&mocks.MockObject{}, &corev1.ConfigMap{})
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("bad new object", func(t *testing.T) {
			result, err := configmap.MaybeUpdate(&corev1.ConfigMap{}, &mocks.MockObject{})
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("bad objects", func(t *testing.T) {
			result, err := configmap.MaybeUpdate(&mocks.MockObject{}, &mocks.MockObject{})
			assert.Error(t, err)
			assert.False(t, result)
		})
	})
	t.Run("compare configmaps", func(t *testing.T) {
		t.Run("empty configmaps", func(t *testing.T) {
			result, err := configmap.MaybeUpdate(&corev1.ConfigMap{}, &corev1.ConfigMap{})
			assert.NoError(t, err)
			assert.False(t, result)
		})
		t.Run("up-to-date configmaps", func(t *testing.T) {
			result, err := configmap.MaybeUpdate(
				&corev1.ConfigMap{Data: map[string]string{"key": "value"}},
				&corev1.ConfigMap{Data: map[string]string{"key": "value"}},
			)
			assert.NoError(t, err)
			assert.False(t, result)
		})
		t.Run("update date in configmaps", func(t *testing.T) {
			existingconfigmap := &corev1.ConfigMap{Data: map[string]string{"key": "value"}}
			newconfigmap := &corev1.ConfigMap{Data: map[string]string{"key": "new-value"}}

			result, err := configmap.MaybeUpdate(existingconfigmap, newconfigmap)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingconfigmap, newconfigmap)
		})
		t.Run("update binary date in configmaps", func(t *testing.T) {
			existingconfigmap := &corev1.ConfigMap{BinaryData: map[string][]byte{"key": []byte("value")}}
			newconfigmap := &corev1.ConfigMap{BinaryData: map[string][]byte{"key": []byte("new-value")}}

			result, err := configmap.MaybeUpdate(existingconfigmap, newconfigmap)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingconfigmap, newconfigmap)
		})
	})
}

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("failed to generate", func(t *testing.T) {
		_, err := configmap.Create(configmap.Conf{GenDataFunc: func(interfaces.Object) (map[string]string, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to generate using custom generator function", func(t *testing.T) {
		_, err := configmap.Create(configmap.Conf{GenConfigMapFunc: func(configmap.Conf) (*corev1.ConfigMap, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to create", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Create(configmap.Conf{
			Instance:  i,
			Reconcile: r,
			AfterCreateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("create configmap", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Create(configmap.Conf{
			Instance:  i,
			Reconcile: r,
		})
		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("failed to generate", func(t *testing.T) {
		_, err := configmap.Update(configmap.Conf{GenDataFunc: func(interfaces.Object) (map[string]string, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to generate using custom generator function", func(t *testing.T) {
		_, err := configmap.Update(configmap.Conf{GenConfigMapFunc: func(configmap.Conf) (*corev1.ConfigMap, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to update", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Update(configmap.Conf{
			Instance:  i,
			Reconcile: r,
			AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("custom maybeupdate function", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Update(configmap.Conf{
			Instance:        i,
			Reconcile:       r,
			MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return true, nil },
		})
		assert.NoError(t, err)
	})
	t.Run("update configmap", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Update(configmap.Conf{
			Instance:  i,
			Reconcile: r,
		})
		assert.NoError(t, err)
	})
}

func TestCreateOrUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("failed to generate", func(t *testing.T) {
		_, err := configmap.CreateOrUpdate(configmap.Conf{GenDataFunc: func(interfaces.Object) (map[string]string, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to generate using custom generator function", func(t *testing.T) {
		_, err := configmap.CreateOrUpdate(configmap.Conf{GenConfigMapFunc: func(configmap.Conf) (*corev1.ConfigMap, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to create", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.CreateOrUpdate(configmap.Conf{
			Instance:  i,
			Reconcile: r,
			AfterCreateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("failed to create", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.CreateOrUpdate(configmap.Conf{
			Name:      "test-existing-configmap",
			Namespace: "test",
			Instance:  i,
			Reconcile: r,
			AfterUpdateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("custom maybeupdate function", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.CreateOrUpdate(configmap.Conf{
			Instance:        i,
			Reconcile:       r,
			MaybeUpdateFunc: func(interfaces.Object, interfaces.Object) (bool, error) { return true, nil },
		})
		assert.NoError(t, err)
	})
	t.Run("create or update configmap", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.CreateOrUpdate(configmap.Conf{
			Instance:  i,
			Reconcile: r,
		})
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("failed to generate", func(t *testing.T) {
		_, err := configmap.Delete(configmap.Conf{GenDataFunc: func(interfaces.Object) (map[string]string, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to generate using custom generator function", func(t *testing.T) {
		_, err := configmap.Delete(configmap.Conf{GenConfigMapFunc: func(configmap.Conf) (*corev1.ConfigMap, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to delete", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Delete(configmap.Conf{
			Instance:  i,
			Reconcile: r,
			AfterDeleteFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("delete configmap", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := configmap.Delete(configmap.Conf{
			Instance:  i,
			Reconcile: r,
		})
		assert.NoError(t, err)
	})
}
