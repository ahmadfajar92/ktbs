package mock_shared

import (
	context "context"
	shared "notifications/src/shared"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecases is a mock of Usecases interface.
type MockUsecases struct {
	ctrl     *gomock.Controller
	recorder *MockUsecasesMockRecorder
}

// MockUsecasesMockRecorder is the mock recorder for MockUsecases.
type MockUsecasesMockRecorder struct {
	mock *MockUsecases
}

// NewMockUsecases creates a new mock instance.
func NewMockUsecases(ctrl *gomock.Controller) *MockUsecases {
	mock := &MockUsecases{ctrl: ctrl}
	mock.recorder = &MockUsecasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecases) EXPECT() *MockUsecasesMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockUsecases) Add(n string, u interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", n, u)
}

// Add indicates an expected call of Add.
func (mr *MockUsecasesMockRecorder) Add(n, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsecases)(nil).Add), n, u)
}

// Call mocks base method.
func (m *MockUsecases) Call(n string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", n)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Call indicates an expected call of Call.
func (mr *MockUsecasesMockRecorder) Call(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUsecases)(nil).Call), n)
}

// MockVendor is a mock of Vendor interface.
type MockVendor struct {
	ctrl     *gomock.Controller
	recorder *MockVendorMockRecorder
}

// MockVendorMockRecorder is the mock recorder for MockVendor.
type MockVendorMockRecorder struct {
	mock *MockVendor
}

// NewMockVendor creates a new mock instance.
func NewMockVendor(ctrl *gomock.Controller) *MockVendor {
	mock := &MockVendor{ctrl: ctrl}
	mock.recorder = &MockVendorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVendor) EXPECT() *MockVendorMockRecorder {
	return m.recorder
}

// IsDefault mocks base method.
func (m *MockVendor) IsDefault() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDefault")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDefault indicates an expected call of IsDefault.
func (mr *MockVendorMockRecorder) IsDefault() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDefault", reflect.TypeOf((*MockVendor)(nil).IsDefault))
}

// IsOn mocks base method.
func (m *MockVendor) IsOn() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOn")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOn indicates an expected call of IsOn.
func (mr *MockVendorMockRecorder) IsOn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOn", reflect.TypeOf((*MockVendor)(nil).IsOn))
}

// OnOff mocks base method.
func (m *MockVendor) OnOff(ctx context.Context, s bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnOff", ctx, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnOff indicates an expected call of OnOff.
func (mr *MockVendorMockRecorder) OnOff(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnOff", reflect.TypeOf((*MockVendor)(nil).OnOff), ctx, s)
}

// Send mocks base method.
func (m *MockVendor) Send(ctx context.Context, p interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, p)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockVendorMockRecorder) Send(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockVendor)(nil).Send), ctx, p)
}

// MockVendors is a mock of Vendors interface.
type MockVendors struct {
	ctrl     *gomock.Controller
	recorder *MockVendorsMockRecorder
}

// MockVendorsMockRecorder is the mock recorder for MockVendors.
type MockVendorsMockRecorder struct {
	mock *MockVendors
}

// NewMockVendors creates a new mock instance.
func NewMockVendors(ctrl *gomock.Controller) *MockVendors {
	mock := &MockVendors{ctrl: ctrl}
	mock.recorder = &MockVendorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVendors) EXPECT() *MockVendorsMockRecorder {
	return m.recorder
}

// ActiveVendor mocks base method.
func (m *MockVendors) ActiveVendor() shared.Vendor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveVendor")
	ret0, _ := ret[0].(shared.Vendor)
	return ret0
}

// ActiveVendor indicates an expected call of ActiveVendor.
func (mr *MockVendorsMockRecorder) ActiveVendor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveVendor", reflect.TypeOf((*MockVendors)(nil).ActiveVendor))
}

// All mocks base method.
func (m *MockVendors) All() map[string]shared.Vendor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(map[string]shared.Vendor)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockVendorsMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockVendors)(nil).All))
}

// Get mocks base method.
func (m *MockVendors) Get(n string) shared.Vendor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", n)
	ret0, _ := ret[0].(shared.Vendor)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockVendorsMockRecorder) Get(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVendors)(nil).Get), n)
}
