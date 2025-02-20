// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// ApprovalService is an autogenerated mock type for the approvalService type
type ApprovalService struct {
	mock.Mock
}

// AdvanceApproval provides a mock function with given fields: _a0, _a1
func (_m *ApprovalService) AdvanceApproval(_a0 context.Context, _a1 *domain.Appeal) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Appeal) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
