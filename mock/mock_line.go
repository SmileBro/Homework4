// Code generated by MockGen. DO NOT EDIT.
// Source: ./distance/line/line.go

// Package mock is a generated GoMock package.
package mock

import (
	point "Homework4/distance/point"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLine is a mock of Line interface.
type MockLine struct {
	ctrl     *gomock.Controller
	recorder *MockLineMockRecorder
}

// MockLineMockRecorder is the mock recorder for MockLine.
type MockLineMockRecorder struct {
	mock *MockLine
}

// NewMockLine creates a new mock instance.
func NewMockLine(ctrl *gomock.Controller) *MockLine {
	mock := &MockLine{ctrl: ctrl}
	mock.recorder = &MockLineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLine) EXPECT() *MockLineMockRecorder {
	return m.recorder
}

// Coordinates mocks base method.
func (m *MockLine) Coordinates() [][]float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Coordinates")
	ret0, _ := ret[0].([][]float64)
	return ret0
}

// Coordinates indicates an expected call of Coordinates.
func (mr *MockLineMockRecorder) Coordinates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Coordinates", reflect.TypeOf((*MockLine)(nil).Coordinates))
}

// Distance mocks base method.
func (m *MockLine) Distance() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Distance")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Distance indicates an expected call of Distance.
func (mr *MockLineMockRecorder) Distance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distance", reflect.TypeOf((*MockLine)(nil).Distance))
}

// Point1 mocks base method.
func (m *MockLine) Point1() point.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Point1")
	ret0, _ := ret[0].(point.Point)
	return ret0
}

// Point1 indicates an expected call of Point1.
func (mr *MockLineMockRecorder) Point1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Point1", reflect.TypeOf((*MockLine)(nil).Point1))
}

// Point2 mocks base method.
func (m *MockLine) Point2() point.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Point2")
	ret0, _ := ret[0].(point.Point)
	return ret0
}

// Point2 indicates an expected call of Point2.
func (mr *MockLineMockRecorder) Point2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Point2", reflect.TypeOf((*MockLine)(nil).Point2))
}
