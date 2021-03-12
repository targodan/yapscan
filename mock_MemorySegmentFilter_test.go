// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package yapscan

import (
	procio "github.com/fkie-cad/yapscan/procio"
	mock "github.com/stretchr/testify/mock"
)

// MockMemorySegmentFilter is an autogenerated mock type for the MemorySegmentFilter type
type MockMemorySegmentFilter struct {
	mock.Mock
}

// Filter provides a mock function with given fields: info
func (_m *MockMemorySegmentFilter) Filter(info *procio.MemorySegmentInfo) *FilterMatch {
	ret := _m.Called(info)

	var r0 *FilterMatch
	if rf, ok := ret.Get(0).(func(*procio.MemorySegmentInfo) *FilterMatch); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FilterMatch)
		}
	}

	return r0
}
