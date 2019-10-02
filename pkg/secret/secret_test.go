package secret_test

import (
	"errors"
	"testing"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/interfaces/mocks"
	"github.com/ankitrgadiya/operatorlib/pkg/secret"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func mockSetup(ctrl *gomock.Controller) (i *mocks.MockObject, r *mocks.MockReconcile) {
	i = mocks.NewMockObject(ctrl)
	i.EXPECT().GetName().Return("test").AnyTimes()
	i.EXPECT().GetUID().Return(types.UID("199bd7a8-b72a-4411-b55e-91096769e58f")).AnyTimes()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: "test-existing-secret", Namespace: "test"},
		Data: map[string][]byte{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
		},
	}
	r = mocks.NewMockReconcile(ctrl)

	c := fake.NewFakeClient([]runtime.Object{secret}...)
	s := scheme.Scheme
	s.AddKnownTypes(schema.GroupVersion{Group: "test", Version: "v1"}, i)

	r.EXPECT().GetClient().Return(c).AnyTimes()
	r.EXPECT().GetScheme().Return(s).AnyTimes()

	return i, r
}

func TestGenerateSecret(t *testing.T) {
	t.Run("generate empty secret", func(t *testing.T) {
		expected := &corev1.Secret{TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		}}

		result, err := secret.GenerateSecret(secret.Conf{})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("failed to generate objectmeta", func(t *testing.T) {
		result, err := secret.GenerateSecret(secret.Conf{
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("failed to generate data", func(t *testing.T) {
		result, err := secret.GenerateSecret(secret.Conf{
			GenDataFunc: func(interfaces.Object) (map[string][]byte, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("failed to generate string data", func(t *testing.T) {
		result, err := secret.GenerateSecret(secret.Conf{
			GenStringDataFunc: func(interfaces.Object) (map[string]string, error) { return nil, errors.New("test error") },
		})
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("generate secret with only string data", func(t *testing.T) {
		expected := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
			Data:     map[string][]byte{"key": []byte("value")},
		}

		result, err := secret.GenerateSecret(secret.Conf{
			GenStringDataFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key": "value"}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("generate secret with only data", func(t *testing.T) {
		expected := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
			Data:     map[string][]byte{"key": []byte("value")},
		}

		result, err := secret.GenerateSecret(secret.Conf{
			GenDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return map[string][]byte{"key": []byte("value")}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("generate secret with data and string data", func(t *testing.T) {
		expected := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
			Data:     map[string][]byte{"key1": []byte("value1"), "key2": []byte("value2")},
		}

		result, err := secret.GenerateSecret(secret.Conf{
			GenDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return map[string][]byte{"key1": []byte("value1")}, nil
			},
			GenStringDataFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key2": "value2"}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("generate secret with data and string data with same keys", func(t *testing.T) {
		expected := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
			Data:     map[string][]byte{"key": []byte("value")},
		}

		result, err := secret.GenerateSecret(secret.Conf{
			GenDataFunc: func(interfaces.Object) (map[string][]byte, error) {
				return map[string][]byte{"key": []byte("value")}, nil
			},
			GenStringDataFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key": "new-value"}, nil
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
