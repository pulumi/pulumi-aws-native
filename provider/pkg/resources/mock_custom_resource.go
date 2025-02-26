// Code generated by MockGen. DO NOT EDIT.
// Source: custom.go
//
// Generated by this command:
//
//	mockgen -package resources -source custom.go -destination mock_custom_resource.go CustomResource
//

// Package resources is a generated GoMock package.
package resources

import (
	context "context"
	reflect "reflect"
	time "time"

	autonaming "github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	resource "github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomResource is a mock of CustomResource interface.
type MockCustomResource struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceMockRecorder
	isgomock struct{}
}

// MockCustomResourceMockRecorder is the mock recorder for MockCustomResource.
type MockCustomResourceMockRecorder struct {
	mock *MockCustomResource
}

// NewMockCustomResource creates a new mock instance.
func NewMockCustomResource(ctrl *gomock.Controller) *MockCustomResource {
	mock := &MockCustomResource{ctrl: ctrl}
	mock.recorder = &MockCustomResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomResource) EXPECT() *MockCustomResourceMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockCustomResource) Check(ctx context.Context, urn resource.URN, engineAutonaming autonaming.EngineAutoNamingConfig, inputs, state resource.PropertyMap, defaultTags map[string]string) (resource.PropertyMap, []ValidationFailure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, urn, engineAutonaming, inputs, state, defaultTags)
	ret0, _ := ret[0].(resource.PropertyMap)
	ret1, _ := ret[1].([]ValidationFailure)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Check indicates an expected call of Check.
func (mr *MockCustomResourceMockRecorder) Check(ctx, urn, engineAutonaming, inputs, state, defaultTags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockCustomResource)(nil).Check), ctx, urn, engineAutonaming, inputs, state, defaultTags)
}

// Create mocks base method.
func (m *MockCustomResource) Create(ctx context.Context, urn resource.URN, inputs resource.PropertyMap, timeout time.Duration) (*string, resource.PropertyMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, urn, inputs, timeout)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(resource.PropertyMap)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockCustomResourceMockRecorder) Create(ctx, urn, inputs, timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomResource)(nil).Create), ctx, urn, inputs, timeout)
}

// Delete mocks base method.
func (m *MockCustomResource) Delete(ctx context.Context, urn resource.URN, id string, inputs, state resource.PropertyMap, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, urn, id, inputs, state, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomResourceMockRecorder) Delete(ctx, urn, id, inputs, state, timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomResource)(nil).Delete), ctx, urn, id, inputs, state, timeout)
}

// PreviewCustomResourceOutputs mocks base method.
func (m *MockCustomResource) PreviewCustomResourceOutputs() resource.PropertyMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreviewCustomResourceOutputs")
	ret0, _ := ret[0].(resource.PropertyMap)
	return ret0
}

// PreviewCustomResourceOutputs indicates an expected call of PreviewCustomResourceOutputs.
func (mr *MockCustomResourceMockRecorder) PreviewCustomResourceOutputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreviewCustomResourceOutputs", reflect.TypeOf((*MockCustomResource)(nil).PreviewCustomResourceOutputs))
}

// Read mocks base method.
func (m *MockCustomResource) Read(ctx context.Context, urn resource.URN, id string, oldInputs, oldOutputs resource.PropertyMap) (resource.PropertyMap, resource.PropertyMap, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, urn, id, oldInputs, oldOutputs)
	ret0, _ := ret[0].(resource.PropertyMap)
	ret1, _ := ret[1].(resource.PropertyMap)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Read indicates an expected call of Read.
func (mr *MockCustomResourceMockRecorder) Read(ctx, urn, id, oldInputs, oldOutputs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCustomResource)(nil).Read), ctx, urn, id, oldInputs, oldOutputs)
}

// Update mocks base method.
func (m *MockCustomResource) Update(ctx context.Context, urn resource.URN, id string, inputs, oldInputs, state resource.PropertyMap, timeout time.Duration) (resource.PropertyMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, urn, id, inputs, oldInputs, state, timeout)
	ret0, _ := ret[0].(resource.PropertyMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomResourceMockRecorder) Update(ctx, urn, id, inputs, oldInputs, state, timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomResource)(nil).Update), ctx, urn, id, inputs, oldInputs, state, timeout)
}
