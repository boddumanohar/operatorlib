package service_test

import (
	"errors"
	"testing"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces/mocks"
	"github.com/ankitrgadiya/operatorlib/pkg/service"

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

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "test-existing-service", Namespace: "test"},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: int32(80)}},
			Type:  "ClusterIP",
		},
	}
	r = mocks.NewMockReconcile(ctrl)

	c := fake.NewFakeClient([]runtime.Object{svc}...)
	s := scheme.Scheme
	s.AddKnownTypes(schema.GroupVersion{Group: "test", Version: "v1"}, i)

	r.EXPECT().GetClient().Return(c).AnyTimes()
	r.EXPECT().GetScheme().Return(s).AnyTimes()

	return i, r
}

func TestGenerateService(t *testing.T) {
	t.Run("generate empty service", func(t *testing.T) {
		expected := &corev1.Service{TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		}}

		result, err := service.GenerateService(service.Conf{})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("failed to generate objectmeta", func(t *testing.T) {
		result, err := service.GenerateService(service.Conf{
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("failed to generate service ports", func(t *testing.T) {
		result, err := service.GenerateService(service.Conf{
			GenServicePortsFunc: func(interfaces.Object) ([]corev1.ServicePort, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("failed to generate selectors", func(t *testing.T) {
		result, err := service.GenerateService(service.Conf{
			GenSelectorFunc: func(interfaces.Object) (map[string]string, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("generate cluster IP service with ports and selectors", func(t *testing.T) {
		expected := &corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			Spec: corev1.ServiceSpec{
				Type:     "ClusterIP",
				Ports:    []corev1.ServicePort{{Port: int32(80)}},
				Selector: map[string]string{"key": "value"},
			},
		}

		result, err := service.GenerateService(service.Conf{
			GenServicePortsFunc: func(interfaces.Object) ([]corev1.ServicePort, error) {
				return []corev1.ServicePort{{Port: int32(80)}}, nil
			},
			GenSelectorFunc: func(interfaces.Object) (map[string]string, error) { return map[string]string{"key": "value"}, nil },
			Type:            "ClusterIP",
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("generate node port service with ports and selectors", func(t *testing.T) {
		expected := &corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			Spec: corev1.ServiceSpec{
				Type:     "NodePort",
				Ports:    []corev1.ServicePort{{Port: int32(80), NodePort: int32(30001)}},
				Selector: map[string]string{"key": "value"},
			},
		}

		result, err := service.GenerateService(service.Conf{
			GenServicePortsFunc: func(interfaces.Object) ([]corev1.ServicePort, error) {
				return []corev1.ServicePort{{Port: int32(80), NodePort: int32(30001)}}, nil
			},
			GenSelectorFunc: func(interfaces.Object) (map[string]string, error) { return map[string]string{"key": "value"}, nil },
			Type:            "NodePort",
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}

func TestMaybeUpdate(t *testing.T) {
	t.Run("bad parameters", func(t *testing.T) {
		t.Run("bad original object", func(t *testing.T) {
			result, err := service.MaybeUpdate(&mocks.MockObject{}, &corev1.Service{})
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("bad new object", func(t *testing.T) {
			result, err := service.MaybeUpdate(&corev1.Service{}, &mocks.MockObject{})
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("bad objects", func(t *testing.T) {
			result, err := service.MaybeUpdate(&mocks.MockObject{}, &mocks.MockObject{})
			assert.Error(t, err)
			assert.False(t, result)
		})
	})
	t.Run("compare services", func(t *testing.T) {
		t.Run("empty services", func(t *testing.T) {
			result, err := service.MaybeUpdate(&corev1.Service{}, &corev1.Service{})
			assert.NoError(t, err)
			assert.False(t, result)
		})
		t.Run("different types", func(t *testing.T) {
			result, err := service.MaybeUpdate(
				&corev1.Service{Spec: corev1.ServiceSpec{Type: "ClusterIP"}},
				&corev1.Service{Spec: corev1.ServiceSpec{Type: "NodePort"}},
			)
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("different number of ports", func(t *testing.T) {
			existingService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(80)},
				},
			}}
			newService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(80)},
					{Port: int32(443)},
				},
			}}

			result, err := service.MaybeUpdate(existingService, newService)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingService, newService)
		})
		t.Run("differnet ports", func(t *testing.T) {
			existingService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(80)},
				},
			}}
			newService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(443)},
				},
			}}

			result, err := service.MaybeUpdate(existingService, newService)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingService, newService)
		})
		t.Run("differnet port names", func(t *testing.T) {
			existingService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(80), Name: "test1"},
				},
			}}
			newService := &corev1.Service{Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Port: int32(80), Name: "test2"},
				},
			}}

			result, err := service.MaybeUpdate(existingService, newService)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingService, newService)
		})
		t.Run("different selectors", func(t *testing.T) {
			existingService := &corev1.Service{Spec: corev1.ServiceSpec{
				Selector: map[string]string{"key": "value"},
			}}
			newService := &corev1.Service{Spec: corev1.ServiceSpec{
				Selector: map[string]string{"key": "new-value"},
			}}

			result, err := service.MaybeUpdate(existingService, newService)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingService, newService)
		})
	})
}

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("failed to generate", func(t *testing.T) {
		_, err := service.Create(service.Conf{GenSelectorFunc: func(interfaces.Object) (map[string]string, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to generate using custom generator function", func(t *testing.T) {
		_, err := service.Create(service.Conf{GenServiceFunc: func(service.Conf) (*corev1.Service, error) {
			return nil, errors.New("test error")
		}})
		assert.Error(t, err)
	})
	t.Run("failed to create", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := service.Create(service.Conf{
			Instance:  i,
			Reconcile: r,
			AfterCreateFunc: func(interfaces.Object, interfaces.Reconcile) (reconcile.Result, error) {
				return reconcile.Result{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("create service", func(t *testing.T) {
		i, r := mockSetup(controller)
		_, err := service.Create(service.Conf{
			Instance:  i,
			Reconcile: r,
		})
		assert.NoError(t, err)
	})
}
