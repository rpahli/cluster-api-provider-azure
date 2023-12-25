/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../tags.go
//
// Generated by this command:
//
//	mockgen -destination tags_mock.go -package mock_tags -source ../tags.go TagScope
//

// Package mock_tags is a generated GoMock package.
package mock_tags

import (
	reflect "reflect"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	gomock "go.uber.org/mock/gomock"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockTagScope is a mock of TagScope interface.
type MockTagScope struct {
	ctrl     *gomock.Controller
	recorder *MockTagScopeMockRecorder
}

// MockTagScopeMockRecorder is the mock recorder for MockTagScope.
type MockTagScopeMockRecorder struct {
	mock *MockTagScope
}

// NewMockTagScope creates a new mock instance.
func NewMockTagScope(ctrl *gomock.Controller) *MockTagScope {
	mock := &MockTagScope{ctrl: ctrl}
	mock.recorder = &MockTagScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagScope) EXPECT() *MockTagScopeMockRecorder {
	return m.recorder
}

// AnnotationJSON mocks base method.
func (m *MockTagScope) AnnotationJSON(arg0 string) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotationJSON", arg0)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnotationJSON indicates an expected call of AnnotationJSON.
func (mr *MockTagScopeMockRecorder) AnnotationJSON(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotationJSON", reflect.TypeOf((*MockTagScope)(nil).AnnotationJSON), arg0)
}

// BaseURI mocks base method.
func (m *MockTagScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockTagScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockTagScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockTagScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockTagScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockTagScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockTagScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockTagScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockTagScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockTagScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockTagScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockTagScope)(nil).CloudEnvironment))
}

// ClusterName mocks base method.
func (m *MockTagScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockTagScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockTagScope)(nil).ClusterName))
}

// HashKey mocks base method.
func (m *MockTagScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockTagScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockTagScope)(nil).HashKey))
}

// SubscriptionID mocks base method.
func (m *MockTagScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockTagScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockTagScope)(nil).SubscriptionID))
}

// TagsSpecs mocks base method.
func (m *MockTagScope) TagsSpecs() []azure.TagsSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagsSpecs")
	ret0, _ := ret[0].([]azure.TagsSpec)
	return ret0
}

// TagsSpecs indicates an expected call of TagsSpecs.
func (mr *MockTagScopeMockRecorder) TagsSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagsSpecs", reflect.TypeOf((*MockTagScope)(nil).TagsSpecs))
}

// TenantID mocks base method.
func (m *MockTagScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockTagScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockTagScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockTagScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockTagScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockTagScope)(nil).Token))
}

// UpdateAnnotationJSON mocks base method.
func (m *MockTagScope) UpdateAnnotationJSON(arg0 string, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnnotationJSON", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAnnotationJSON indicates an expected call of UpdateAnnotationJSON.
func (mr *MockTagScopeMockRecorder) UpdateAnnotationJSON(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnnotationJSON", reflect.TypeOf((*MockTagScope)(nil).UpdateAnnotationJSON), arg0, arg1)
}
