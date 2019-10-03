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

func TestMaybeUpdate(t *testing.T) {
	t.Run("bad parameters", func(t *testing.T) {
		t.Run("bad original object", func(t *testing.T) {
			result, err := secret.MaybeUpdate(&mocks.MockObject{}, &corev1.Secret{})
			assert.Error(t, err)
			assert.False(t, result)
		})
		t.Run("bad new object", func(t *testing.T) {
			result, err := secret.MaybeUpdate(&corev1.Secret{}, &mocks.MockObject{})
			assert.Error(t, err)
			assert.False(t, result)
		})
	})
	t.Run("compare secrets", func(t *testing.T) {
		t.Run("empty secrets", func(t *testing.T) {
			result, err := secret.MaybeUpdate(&corev1.Secret{}, &corev1.Secret{})
			assert.NoError(t, err)
			assert.False(t, result)
		})
		t.Run("up-to-date secrets", func(t *testing.T) {
			result, err := secret.MaybeUpdate(
				&corev1.Secret{Data: map[string][]byte{"key": []byte("value")}},
				&corev1.Secret{Data: map[string][]byte{"key": []byte("value")}},
			)
			assert.NoError(t, err)
			assert.False(t, result)
		})
		t.Run("update data in secret", func(t *testing.T) {
			existingSecret := &corev1.Secret{Data: map[string][]byte{"key": []byte("value")}}
			newSecret := &corev1.Secret{Data: map[string][]byte{"key": []byte("new-value")}}
			result, err := secret.MaybeUpdate(existingSecret, newSecret)
			assert.NoError(t, err)
			assert.True(t, result)
			assert.Equal(t, existingSecret, newSecret)
		})
	})
}
