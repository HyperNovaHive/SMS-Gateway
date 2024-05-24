// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/sms/receiver.go
//
// Generated by this command:
//
//	mockgen -source=./internal/sms/receiver.go -package=sms -destination=./internal/sms/mock_receiver.go
//

// Package sms is a generated GoMock package.
package sms

import (
	reflect "reflect"

	gsm "github.com/warthog618/modem/gsm"
	gomock "go.uber.org/mock/gomock"
)

// MockReceiver is a mock of Receiver interface.
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver.
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance.
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// Listen mocks base method.
func (m *MockReceiver) Listen(handler func(gsm.Message)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Listen", handler)
}

// Listen indicates an expected call of Listen.
func (mr *MockReceiverMockRecorder) Listen(handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockReceiver)(nil).Listen), handler)
}
