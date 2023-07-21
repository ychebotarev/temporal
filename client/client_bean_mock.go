// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: client_bean.go

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/workflowservice/v1"
	v10 "go.temporal.io/server/api/adminservice/v1"
	v11 "go.temporal.io/server/api/historyservice/v1"
	v12 "go.temporal.io/server/api/matchingservice/v1"
	grpc "google.golang.org/grpc"
)

// MockBean is a mock of Bean interface.
type MockBean struct {
	ctrl     *gomock.Controller
	recorder *MockBeanMockRecorder
}

// MockBeanMockRecorder is the mock recorder for MockBean.
type MockBeanMockRecorder struct {
	mock *MockBean
}

// NewMockBean creates a new mock instance.
func NewMockBean(ctrl *gomock.Controller) *MockBean {
	mock := &MockBean{ctrl: ctrl}
	mock.recorder = &MockBeanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBean) EXPECT() *MockBeanMockRecorder {
	return m.recorder
}

// GetFrontendClient mocks base method.
func (m *MockBean) GetFrontendClient() v1.WorkflowServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendClient")
	ret0, _ := ret[0].(v1.WorkflowServiceClient)
	return ret0
}

// GetFrontendClient indicates an expected call of GetFrontendClient.
func (mr *MockBeanMockRecorder) GetFrontendClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendClient", reflect.TypeOf((*MockBean)(nil).GetFrontendClient))
}

// GetHistoryClient mocks base method.
func (m *MockBean) GetHistoryClient() v11.HistoryServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryClient")
	ret0, _ := ret[0].(v11.HistoryServiceClient)
	return ret0
}

// GetHistoryClient indicates an expected call of GetHistoryClient.
func (mr *MockBeanMockRecorder) GetHistoryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryClient", reflect.TypeOf((*MockBean)(nil).GetHistoryClient))
}

// GetMatchingClient mocks base method.
func (m *MockBean) GetMatchingClient(namespaceIDToName NamespaceIDToNameFunc) (v12.MatchingServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingClient", namespaceIDToName)
	ret0, _ := ret[0].(v12.MatchingServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchingClient indicates an expected call of GetMatchingClient.
func (mr *MockBeanMockRecorder) GetMatchingClient(namespaceIDToName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingClient", reflect.TypeOf((*MockBean)(nil).GetMatchingClient), namespaceIDToName)
}

// GetRemoteAdminClient mocks base method.
func (m *MockBean) GetRemoteAdminClient(arg0 string) (v10.AdminServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteAdminClient", arg0)
	ret0, _ := ret[0].(v10.AdminServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteAdminClient indicates an expected call of GetRemoteAdminClient.
func (mr *MockBeanMockRecorder) GetRemoteAdminClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteAdminClient", reflect.TypeOf((*MockBean)(nil).GetRemoteAdminClient), arg0)
}

// GetRemoteFrontendClient mocks base method.
func (m *MockBean) GetRemoteFrontendClient(arg0 string) (grpc.ClientConnInterface, v1.WorkflowServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteFrontendClient", arg0)
	ret0, _ := ret[0].(grpc.ClientConnInterface)
	ret1, _ := ret[1].(v1.WorkflowServiceClient)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRemoteFrontendClient indicates an expected call of GetRemoteFrontendClient.
func (mr *MockBeanMockRecorder) GetRemoteFrontendClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteFrontendClient", reflect.TypeOf((*MockBean)(nil).GetRemoteFrontendClient), arg0)
}

// SetRemoteAdminClient mocks base method.
func (m *MockBean) SetRemoteAdminClient(arg0 string, arg1 v10.AdminServiceClient) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteAdminClient", arg0, arg1)
}

// SetRemoteAdminClient indicates an expected call of SetRemoteAdminClient.
func (mr *MockBeanMockRecorder) SetRemoteAdminClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemoteAdminClient", reflect.TypeOf((*MockBean)(nil).SetRemoteAdminClient), arg0, arg1)
}