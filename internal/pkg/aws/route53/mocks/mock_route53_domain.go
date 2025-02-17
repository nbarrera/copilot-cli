// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/route53/route53domains.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	route53domains "github.com/aws/aws-sdk-go/service/route53domains"
	gomock "github.com/golang/mock/gomock"
)

// MockdomainAPI is a mock of domainAPI interface.
type MockdomainAPI struct {
	ctrl     *gomock.Controller
	recorder *MockdomainAPIMockRecorder
}

// MockdomainAPIMockRecorder is the mock recorder for MockdomainAPI.
type MockdomainAPIMockRecorder struct {
	mock *MockdomainAPI
}

// NewMockdomainAPI creates a new mock instance.
func NewMockdomainAPI(ctrl *gomock.Controller) *MockdomainAPI {
	mock := &MockdomainAPI{ctrl: ctrl}
	mock.recorder = &MockdomainAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdomainAPI) EXPECT() *MockdomainAPIMockRecorder {
	return m.recorder
}

// GetDomainDetail mocks base method.
func (m *MockdomainAPI) GetDomainDetail(input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainDetail", input)
	ret0, _ := ret[0].(*route53domains.GetDomainDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainDetail indicates an expected call of GetDomainDetail.
func (mr *MockdomainAPIMockRecorder) GetDomainDetail(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainDetail", reflect.TypeOf((*MockdomainAPI)(nil).GetDomainDetail), input)
}
