// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package stores is a generated GoMock package.
package stores

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "gofr.dev/examples/using-dynamodb/models"
	gofr "gofr.dev/pkg/gofr"
)

// MockPerson is a mock of Person interface.
type MockPerson struct {
	ctrl     *gomock.Controller
	recorder *MockPersonMockRecorder
}

// MockPersonMockRecorder is the mock recorder for MockPerson.
type MockPersonMockRecorder struct {
	mock *MockPerson
}

// NewMockPerson creates a new mock instance.
func NewMockPerson(ctrl *gomock.Controller) *MockPerson {
	mock := &MockPerson{ctrl: ctrl}
	mock.recorder = &MockPersonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerson) EXPECT() *MockPersonMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPerson) Create(c *gofr.Context, user models.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPersonMockRecorder) Create(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPerson)(nil).Create), c, user)
}

// Delete mocks base method.
func (m *MockPerson) Delete(c *gofr.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPersonMockRecorder) Delete(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPerson)(nil).Delete), c, id)
}

// Get mocks base method.
func (m *MockPerson) Get(c *gofr.Context, id string) (models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c, id)
	ret0, _ := ret[0].(models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPersonMockRecorder) Get(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPerson)(nil).Get), c, id)
}

// Update mocks base method.
func (m *MockPerson) Update(c *gofr.Context, user models.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPersonMockRecorder) Update(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPerson)(nil).Update), c, user)
}
