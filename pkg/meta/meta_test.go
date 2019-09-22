package meta_test

import (
	"errors"
	"testing"

	"github.com/ankitrgadiya/operatorlib/pkg/interfaces"
	"github.com/ankitrgadiya/operatorlib/pkg/meta"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGenerateObjectMeta(t *testing.T) {
	t.Run("empty objectmeta", func(t *testing.T) {
		expected := &metav1.ObjectMeta{}

		result, err := meta.GenerateObjectMeta(meta.Conf{})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with name and namespace", func(t *testing.T) {
		expected := &metav1.ObjectMeta{Name: "test-object", Namespace: "test"}

		result, err := meta.GenerateObjectMeta(meta.Conf{
			Name: "test-object", Namespace: "test",
		})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with labels", func(t *testing.T) {
		expected := &metav1.ObjectMeta{
			Name:      "test-object",
			Namespace: "test",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		}

		result, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{
					"key1": "value1",
					"key2": "value2",
				}, nil
			},
		})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with labels fail", func(t *testing.T) {
		_, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("objectmeta with labels append", func(t *testing.T) {
		expected := &metav1.ObjectMeta{
			Name:      "test-object",
			Namespace: "test",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		}

		result, err := meta.GenerateObjectMeta(meta.Conf{
			Instance:  &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"key2": "value2"}}},
			Name:      "test-object",
			Namespace: "test",
			GenLabelsFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{"key1": "value1"}, nil
			},
			AppendLabels: true,
		})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with annotations", func(t *testing.T) {
		expected := &metav1.ObjectMeta{
			Name:      "test-object",
			Namespace: "test",
			Annotations: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		}

		result, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenAnnotationsFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{
					"key1": "value1",
					"key2": "value2",
				}, nil
			},
		})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with annotations fail", func(t *testing.T) {
		_, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenAnnotationsFunc: func(interfaces.Object) (map[string]string, error) {
				return map[string]string{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
	t.Run("objectmeta with finalizers", func(t *testing.T) {
		expected := &metav1.ObjectMeta{
			Name:      "test-object",
			Namespace: "test",
			Finalizers: []string{
				"finalizer1",
				"finalizer2",
			},
		}

		result, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenFinalizersFunc: func(interfaces.Object) ([]string, error) {
				return []string{
					"finalizer1",
					"finalizer2",
				}, nil
			},
		})
		assert.NoError(t, err)

		assert.Equal(t, expected, result)
	})
	t.Run("objectmeta with finalizers fail", func(t *testing.T) {
		_, err := meta.GenerateObjectMeta(meta.Conf{
			Name:      "test-object",
			Namespace: "test",
			GenFinalizersFunc: func(interfaces.Object) ([]string, error) {
				return []string{}, errors.New("test error")
			},
		})
		assert.Error(t, err)
	})
}
