// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"

	"github.com/stretchr/testify/mock"
)

// IDatabaseRead is an autogenerated mock type for the IDatabaseRead type
type IDatabaseRead struct {
	mock.Mock
}

// Find provides a mock function with given fields: entityPointer, where, table
func (_m *IDatabaseRead) Find(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse {
	ret := _m.Called(entityPointer, where, table)

	var r0 response.IResponse
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}, string) response.IResponse); ok {
		r0 = rf(entityPointer, where, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.IResponse)
		}
	}

	return r0
}

// FindPreload provides a mock function with given fields: entityPointer, where, preloads, table
func (_m *IDatabaseRead) FindPreload(entityPointer interface{}, where map[string]interface{}, preloads map[string][]interface{}, table string) response.IResponse {
	ret := _m.Called(entityPointer, where, preloads, table)

	var r0 response.IResponse
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}, map[string][]interface{}, string) response.IResponse); ok {
		r0 = rf(entityPointer, where, preloads, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.IResponse)
		}
	}

	return r0
}

// First provides a mock function with given fields: entityPointer, where, table
func (_m *IDatabaseRead) First(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse {
	ret := _m.Called(entityPointer, where, table)

	var r0 response.IResponse
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}, string) response.IResponse); ok {
		r0 = rf(entityPointer, where, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.IResponse)
		}
	}

	return r0
}

// IsAvailable provides a mock function with given fields:
func (_m *IDatabaseRead) IsAvailable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Raw provides a mock function with given fields: rawSQL, entityPointer, values
func (_m *IDatabaseRead) Raw(rawSQL string, entityPointer interface{}, values ...interface{}) response.IResponse {
	var _ca []interface{}
	_ca = append(_ca, rawSQL, entityPointer)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 response.IResponse
	if rf, ok := ret.Get(0).(func(string, interface{}, ...interface{}) response.IResponse); ok {
		r0 = rf(rawSQL, entityPointer, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.IResponse)
		}
	}

	return r0
}
