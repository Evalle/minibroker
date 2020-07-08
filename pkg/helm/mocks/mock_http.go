// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/minibroker/pkg/helm (interfaces: HTTPGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHTTPGetter is a mock of HTTPGetter interface
type MockHTTPGetter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPGetterMockRecorder
}

// MockHTTPGetterMockRecorder is the mock recorder for MockHTTPGetter
type MockHTTPGetterMockRecorder struct {
	mock *MockHTTPGetter
}

// NewMockHTTPGetter creates a new mock instance
func NewMockHTTPGetter(ctrl *gomock.Controller) *MockHTTPGetter {
	mock := &MockHTTPGetter{ctrl: ctrl}
	mock.recorder = &MockHTTPGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPGetter) EXPECT() *MockHTTPGetterMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockHTTPGetter) Get(arg0 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHTTPGetterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPGetter)(nil).Get), arg0)
}
