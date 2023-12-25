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
// Source: ../client.go
//
// Generated by this command:
//
//	mockgen -destination client_mock.go -package mock_inboundnatrules -source ../client.go Client
//

// Package mock_inboundnatrules is a generated GoMock package.
package mock_inboundnatrules

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	gomock "go.uber.org/mock/gomock"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *Mockclient) List(arg0 context.Context, arg1, arg2 string) ([]armnetwork.InboundNatRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]armnetwork.InboundNatRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockclientMockRecorder) List(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Mockclient)(nil).List), arg0, arg1, arg2)
}
