// Code generated by MockGen. DO NOT EDIT.
// Source: argp.in/go/operatorlib/pkg/interfaces (interfaces: Object)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockObject is a mock of Object interface
type MockObject struct {
	ctrl     *gomock.Controller
	recorder *MockObjectMockRecorder
}

// MockObjectMockRecorder is the mock recorder for MockObject
type MockObjectMockRecorder struct {
	mock *MockObject
}

// NewMockObject creates a new mock instance
func NewMockObject(ctrl *gomock.Controller) *MockObject {
	mock := &MockObject{ctrl: ctrl}
	mock.recorder = &MockObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObject) EXPECT() *MockObjectMockRecorder {
	return m.recorder
}

// DeepCopyObject mocks base method
func (m *MockObject) DeepCopyObject() runtime.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopyObject")
	ret0, _ := ret[0].(runtime.Object)
	return ret0
}

// DeepCopyObject indicates an expected call of DeepCopyObject
func (mr *MockObjectMockRecorder) DeepCopyObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyObject", reflect.TypeOf((*MockObject)(nil).DeepCopyObject))
}

// GetAnnotations mocks base method
func (m *MockObject) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations
func (mr *MockObjectMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockObject)(nil).GetAnnotations))
}

// GetClusterName mocks base method
func (m *MockObject) GetClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterName indicates an expected call of GetClusterName
func (mr *MockObjectMockRecorder) GetClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterName", reflect.TypeOf((*MockObject)(nil).GetClusterName))
}

// GetCreationTimestamp mocks base method
func (m *MockObject) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp
func (mr *MockObjectMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).GetCreationTimestamp))
}

// GetDeletionGracePeriodSeconds mocks base method
func (m *MockObject) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds
func (mr *MockObjectMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).GetDeletionGracePeriodSeconds))
}

// GetDeletionTimestamp mocks base method
func (m *MockObject) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp
func (mr *MockObjectMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).GetDeletionTimestamp))
}

// GetFinalizers mocks base method
func (m *MockObject) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers
func (mr *MockObjectMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockObject)(nil).GetFinalizers))
}

// GetGenerateName mocks base method
func (m *MockObject) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName
func (mr *MockObjectMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockObject)(nil).GetGenerateName))
}

// GetGeneration mocks base method
func (m *MockObject) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration
func (mr *MockObjectMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockObject)(nil).GetGeneration))
}

// GetInitializers mocks base method
func (m *MockObject) GetInitializers() *v1.Initializers {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitializers")
	ret0, _ := ret[0].(*v1.Initializers)
	return ret0
}

// GetInitializers indicates an expected call of GetInitializers
func (mr *MockObjectMockRecorder) GetInitializers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitializers", reflect.TypeOf((*MockObject)(nil).GetInitializers))
}

// GetLabels mocks base method
func (m *MockObject) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels
func (mr *MockObjectMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockObject)(nil).GetLabels))
}

// GetManagedFields mocks base method
func (m *MockObject) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields
func (mr *MockObjectMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockObject)(nil).GetManagedFields))
}

// GetName mocks base method
func (m *MockObject) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockObjectMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockObject)(nil).GetName))
}

// GetNamespace mocks base method
func (m *MockObject) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockObjectMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockObject)(nil).GetNamespace))
}

// GetObjectKind mocks base method
func (m *MockObject) GetObjectKind() schema.ObjectKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectKind")
	ret0, _ := ret[0].(schema.ObjectKind)
	return ret0
}

// GetObjectKind indicates an expected call of GetObjectKind
func (mr *MockObjectMockRecorder) GetObjectKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectKind", reflect.TypeOf((*MockObject)(nil).GetObjectKind))
}

// GetOwnerReferences mocks base method
func (m *MockObject) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences
func (mr *MockObjectMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockObject)(nil).GetOwnerReferences))
}

// GetResourceVersion mocks base method
func (m *MockObject) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion
func (mr *MockObjectMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockObject)(nil).GetResourceVersion))
}

// GetSelfLink mocks base method
func (m *MockObject) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink
func (mr *MockObjectMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockObject)(nil).GetSelfLink))
}

// GetUID mocks base method
func (m *MockObject) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID
func (mr *MockObjectMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockObject)(nil).GetUID))
}

// SetAnnotations mocks base method
func (m *MockObject) SetAnnotations(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", arg0)
}

// SetAnnotations indicates an expected call of SetAnnotations
func (mr *MockObjectMockRecorder) SetAnnotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockObject)(nil).SetAnnotations), arg0)
}

// SetClusterName mocks base method
func (m *MockObject) SetClusterName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterName", arg0)
}

// SetClusterName indicates an expected call of SetClusterName
func (mr *MockObjectMockRecorder) SetClusterName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterName", reflect.TypeOf((*MockObject)(nil).SetClusterName), arg0)
}

// SetCreationTimestamp mocks base method
func (m *MockObject) SetCreationTimestamp(arg0 v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", arg0)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp
func (mr *MockObjectMockRecorder) SetCreationTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).SetCreationTimestamp), arg0)
}

// SetDeletionGracePeriodSeconds mocks base method
func (m *MockObject) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds
func (mr *MockObjectMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// SetDeletionTimestamp mocks base method
func (m *MockObject) SetDeletionTimestamp(arg0 *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", arg0)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp
func (mr *MockObjectMockRecorder) SetDeletionTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).SetDeletionTimestamp), arg0)
}

// SetFinalizers mocks base method
func (m *MockObject) SetFinalizers(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", arg0)
}

// SetFinalizers indicates an expected call of SetFinalizers
func (mr *MockObjectMockRecorder) SetFinalizers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockObject)(nil).SetFinalizers), arg0)
}

// SetGenerateName mocks base method
func (m *MockObject) SetGenerateName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", arg0)
}

// SetGenerateName indicates an expected call of SetGenerateName
func (mr *MockObjectMockRecorder) SetGenerateName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockObject)(nil).SetGenerateName), arg0)
}

// SetGeneration mocks base method
func (m *MockObject) SetGeneration(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", arg0)
}

// SetGeneration indicates an expected call of SetGeneration
func (mr *MockObjectMockRecorder) SetGeneration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockObject)(nil).SetGeneration), arg0)
}

// SetInitializers mocks base method
func (m *MockObject) SetInitializers(arg0 *v1.Initializers) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInitializers", arg0)
}

// SetInitializers indicates an expected call of SetInitializers
func (mr *MockObjectMockRecorder) SetInitializers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitializers", reflect.TypeOf((*MockObject)(nil).SetInitializers), arg0)
}

// SetLabels mocks base method
func (m *MockObject) SetLabels(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", arg0)
}

// SetLabels indicates an expected call of SetLabels
func (mr *MockObjectMockRecorder) SetLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockObject)(nil).SetLabels), arg0)
}

// SetManagedFields mocks base method
func (m *MockObject) SetManagedFields(arg0 []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", arg0)
}

// SetManagedFields indicates an expected call of SetManagedFields
func (mr *MockObjectMockRecorder) SetManagedFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockObject)(nil).SetManagedFields), arg0)
}

// SetName mocks base method
func (m *MockObject) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName
func (mr *MockObjectMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockObject)(nil).SetName), arg0)
}

// SetNamespace mocks base method
func (m *MockObject) SetNamespace(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", arg0)
}

// SetNamespace indicates an expected call of SetNamespace
func (mr *MockObjectMockRecorder) SetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockObject)(nil).SetNamespace), arg0)
}

// SetOwnerReferences mocks base method
func (m *MockObject) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences
func (mr *MockObjectMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockObject)(nil).SetOwnerReferences), arg0)
}

// SetResourceVersion mocks base method
func (m *MockObject) SetResourceVersion(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", arg0)
}

// SetResourceVersion indicates an expected call of SetResourceVersion
func (mr *MockObjectMockRecorder) SetResourceVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockObject)(nil).SetResourceVersion), arg0)
}

// SetSelfLink mocks base method
func (m *MockObject) SetSelfLink(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", arg0)
}

// SetSelfLink indicates an expected call of SetSelfLink
func (mr *MockObjectMockRecorder) SetSelfLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockObject)(nil).SetSelfLink), arg0)
}

// SetUID mocks base method
func (m *MockObject) SetUID(arg0 types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", arg0)
}

// SetUID indicates an expected call of SetUID
func (mr *MockObjectMockRecorder) SetUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockObject)(nil).SetUID), arg0)
}
