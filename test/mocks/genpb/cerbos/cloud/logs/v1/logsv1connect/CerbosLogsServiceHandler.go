// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocklogsv1connect

import (
	context "context"

	connect "connectrpc.com/connect"

	logsv1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/logs/v1"

	mock "github.com/stretchr/testify/mock"
)

// CerbosLogsServiceHandler is an autogenerated mock type for the CerbosLogsServiceHandler type
type CerbosLogsServiceHandler struct {
	mock.Mock
}

type CerbosLogsServiceHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *CerbosLogsServiceHandler) EXPECT() *CerbosLogsServiceHandler_Expecter {
	return &CerbosLogsServiceHandler_Expecter{mock: &_m.Mock}
}

// Ingest provides a mock function with given fields: _a0, _a1
func (_m *CerbosLogsServiceHandler) Ingest(_a0 context.Context, _a1 *connect.Request[logsv1.IngestRequest]) (*connect.Response[logsv1.IngestResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Ingest")
	}

	var r0 *connect.Response[logsv1.IngestResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[logsv1.IngestRequest]) (*connect.Response[logsv1.IngestResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[logsv1.IngestRequest]) *connect.Response[logsv1.IngestResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[logsv1.IngestResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[logsv1.IngestRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CerbosLogsServiceHandler_Ingest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ingest'
type CerbosLogsServiceHandler_Ingest_Call struct {
	*mock.Call
}

// Ingest is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[logsv1.IngestRequest]
func (_e *CerbosLogsServiceHandler_Expecter) Ingest(_a0 interface{}, _a1 interface{}) *CerbosLogsServiceHandler_Ingest_Call {
	return &CerbosLogsServiceHandler_Ingest_Call{Call: _e.mock.On("Ingest", _a0, _a1)}
}

func (_c *CerbosLogsServiceHandler_Ingest_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[logsv1.IngestRequest])) *CerbosLogsServiceHandler_Ingest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[logsv1.IngestRequest]))
	})
	return _c
}

func (_c *CerbosLogsServiceHandler_Ingest_Call) Return(_a0 *connect.Response[logsv1.IngestResponse], _a1 error) *CerbosLogsServiceHandler_Ingest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CerbosLogsServiceHandler_Ingest_Call) RunAndReturn(run func(context.Context, *connect.Request[logsv1.IngestRequest]) (*connect.Response[logsv1.IngestResponse], error)) *CerbosLogsServiceHandler_Ingest_Call {
	_c.Call.Return(run)
	return _c
}

// NewCerbosLogsServiceHandler creates a new instance of CerbosLogsServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCerbosLogsServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *CerbosLogsServiceHandler {
	mock := &CerbosLogsServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
