// Code generated by MockGen. DO NOT EDIT.
// Source: sample/sample.go

// Package mock_sample is a generated GoMock package.
package mock_sample

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSample is a mock of Sample interface
type MockSample struct {
	ctrl     *gomock.Controller
	recorder *MockSampleMockRecorder
}

// MockSampleMockRecorder is the mock recorder for MockSample
type MockSampleMockRecorder struct {
	mock *MockSample
}

// NewMockSample creates a new mock instance
func NewMockSample(ctrl *gomock.Controller) *MockSample {
	mock := &MockSample{ctrl: ctrl}
	mock.recorder = &MockSampleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSample) EXPECT() *MockSampleMockRecorder {
	return m.recorder
}

// Method mocks base method
func (m *MockSample) Method(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Method indicates an expected call of Method
func (mr *MockSampleMockRecorder) Method(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockSample)(nil).Method), a, b)
}