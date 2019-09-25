// /*
// Copyright 2019 OKN Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: okn/pkg/agent/openflow (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	net "net"
	openflow "okn/pkg/ovs/openflow"
	reflect "reflect"
)

// MockOFClient is a mock of Client interface
type MockOFClient struct {
	ctrl     *gomock.Controller
	recorder *MockOFClientMockRecorder
}

// MockOFClientMockRecorder is the mock recorder for MockOFClient
type MockOFClientMockRecorder struct {
	mock *MockOFClient
}

// NewMockOFClient creates a new mock instance
func NewMockOFClient(ctrl *gomock.Controller) *MockOFClient {
	mock := &MockOFClient{ctrl: ctrl}
	mock.recorder = &MockOFClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOFClient) EXPECT() *MockOFClientMockRecorder {
	return m.recorder
}

// GetFlowTableStatus mocks base method
func (m *MockOFClient) GetFlowTableStatus() []openflow.TableStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowTableStatus")
	ret0, _ := ret[0].([]openflow.TableStatus)
	return ret0
}

// GetFlowTableStatus indicates an expected call of GetFlowTableStatus
func (mr *MockOFClientMockRecorder) GetFlowTableStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowTableStatus", reflect.TypeOf((*MockOFClient)(nil).GetFlowTableStatus))
}

// Initialize mocks base method
func (m *MockOFClient) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockOFClientMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockOFClient)(nil).Initialize))
}

// InstallGatewayFlows mocks base method
func (m *MockOFClient) InstallGatewayFlows(arg0 net.IP, arg1 net.HardwareAddr, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallGatewayFlows", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallGatewayFlows indicates an expected call of InstallGatewayFlows
func (mr *MockOFClientMockRecorder) InstallGatewayFlows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallGatewayFlows", reflect.TypeOf((*MockOFClient)(nil).InstallGatewayFlows), arg0, arg1, arg2)
}

// InstallNodeFlows mocks base method
func (m *MockOFClient) InstallNodeFlows(arg0 string, arg1 net.HardwareAddr, arg2 net.IP, arg3 net.IPNet, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallNodeFlows", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallNodeFlows indicates an expected call of InstallNodeFlows
func (mr *MockOFClientMockRecorder) InstallNodeFlows(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallNodeFlows", reflect.TypeOf((*MockOFClient)(nil).InstallNodeFlows), arg0, arg1, arg2, arg3, arg4)
}

// InstallPodFlows mocks base method
func (m *MockOFClient) InstallPodFlows(arg0 string, arg1 net.IP, arg2, arg3 net.HardwareAddr, arg4 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPodFlows", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallPodFlows indicates an expected call of InstallPodFlows
func (mr *MockOFClientMockRecorder) InstallPodFlows(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPodFlows", reflect.TypeOf((*MockOFClient)(nil).InstallPodFlows), arg0, arg1, arg2, arg3, arg4)
}

// InstallServiceFlows mocks base method
func (m *MockOFClient) InstallServiceFlows(arg0 string, arg1 *net.IPNet, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallServiceFlows", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallServiceFlows indicates an expected call of InstallServiceFlows
func (mr *MockOFClientMockRecorder) InstallServiceFlows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallServiceFlows", reflect.TypeOf((*MockOFClient)(nil).InstallServiceFlows), arg0, arg1, arg2)
}

// InstallTunnelFlows mocks base method
func (m *MockOFClient) InstallTunnelFlows(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallTunnelFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallTunnelFlows indicates an expected call of InstallTunnelFlows
func (mr *MockOFClientMockRecorder) InstallTunnelFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallTunnelFlows", reflect.TypeOf((*MockOFClient)(nil).InstallTunnelFlows), arg0)
}

// UninstallNodeFlows mocks base method
func (m *MockOFClient) UninstallNodeFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallNodeFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallNodeFlows indicates an expected call of UninstallNodeFlows
func (mr *MockOFClientMockRecorder) UninstallNodeFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallNodeFlows", reflect.TypeOf((*MockOFClient)(nil).UninstallNodeFlows), arg0)
}

// UninstallPodFlows mocks base method
func (m *MockOFClient) UninstallPodFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPodFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallPodFlows indicates an expected call of UninstallPodFlows
func (mr *MockOFClientMockRecorder) UninstallPodFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPodFlows", reflect.TypeOf((*MockOFClient)(nil).UninstallPodFlows), arg0)
}

// UninstallServiceFlows mocks base method
func (m *MockOFClient) UninstallServiceFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallServiceFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallServiceFlows indicates an expected call of UninstallServiceFlows
func (mr *MockOFClientMockRecorder) UninstallServiceFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallServiceFlows", reflect.TypeOf((*MockOFClient)(nil).UninstallServiceFlows), arg0)
}
