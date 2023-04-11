// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/golang/mock/sample (interfaces: Index,Embed,Embedded)

// Package user_test is a generated GoMock package.
package user_test

import (
	bufio "bufio"
	bytes "bytes"
	hash "hash"
	template "html/template"
	io "io"
	http "net/http"
	reflect "reflect"
	template0 "text/template"

	gomock "github.com/golang/mock/gomock"
	imp1 "github.com/golang/mock/sample/imp1"
	imp2 "github.com/golang/mock/sample/imp2"
	imp3 "github.com/golang/mock/sample/imp3"
	imp_four "github.com/golang/mock/sample/imp4"
)

// MockIndex is a mock of Index interface.
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex.
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance.
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// Anon mocks base method.
func (m *MockIndex) Anon(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Anon", arg0)
}

// Anon indicates an expected call of Anon.
func (mr *MockIndexMockRecorder) Anon(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Anon", reflect.TypeOf((*MockIndex)(nil).Anon), arg0)
}

// AnonStructReturn mocks base method.
func (m *MockIndex) AnonStructReturn() struct {
	ValueA string
	ValueB imp1.Imp1
	Nested struct {
		Value int
	}
} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnonStructReturn")
	ret0, _ := ret[0].(struct {
		ValueA string
		ValueB imp1.Imp1
		Nested struct {
			Value int
		}
	})
	return ret0
}

// AnonStructReturn indicates an expected call of AnonStructReturn.
func (mr *MockIndexMockRecorder) AnonStructReturn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnonStructReturn", reflect.TypeOf((*MockIndex)(nil).AnonStructReturn))
}

// Chan mocks base method.
func (m *MockIndex) Chan(arg0 chan int, arg1 chan<- hash.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Chan", arg0, arg1)
}

// Chan indicates an expected call of Chan.
func (mr *MockIndexMockRecorder) Chan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chan", reflect.TypeOf((*MockIndex)(nil).Chan), arg0, arg1)
}

// ConcreteRet mocks base method.
func (m *MockIndex) ConcreteRet() chan<- bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcreteRet")
	ret0, _ := ret[0].(chan<- bool)
	return ret0
}

// ConcreteRet indicates an expected call of ConcreteRet.
func (mr *MockIndexMockRecorder) ConcreteRet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcreteRet", reflect.TypeOf((*MockIndex)(nil).ConcreteRet))
}

// Ellip mocks base method.
func (m *MockIndex) Ellip(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Ellip", varargs...)
}

// Ellip indicates an expected call of Ellip.
func (mr *MockIndexMockRecorder) Ellip(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ellip", reflect.TypeOf((*MockIndex)(nil).Ellip), varargs...)
}

// EllipOnly mocks base method.
func (m *MockIndex) EllipOnly(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "EllipOnly", varargs...)
}

// EllipOnly indicates an expected call of EllipOnly.
func (mr *MockIndexMockRecorder) EllipOnly(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EllipOnly", reflect.TypeOf((*MockIndex)(nil).EllipOnly), arg0...)
}

// ForeignFour mocks base method.
func (m *MockIndex) ForeignFour(arg0 imp_four.Imp4) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForeignFour", arg0)
}

// ForeignFour indicates an expected call of ForeignFour.
func (mr *MockIndexMockRecorder) ForeignFour(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForeignFour", reflect.TypeOf((*MockIndex)(nil).ForeignFour), arg0)
}

// ForeignOne mocks base method.
func (m *MockIndex) ForeignOne(arg0 imp1.Imp1) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForeignOne", arg0)
}

// ForeignOne indicates an expected call of ForeignOne.
func (mr *MockIndexMockRecorder) ForeignOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForeignOne", reflect.TypeOf((*MockIndex)(nil).ForeignOne), arg0)
}

// ForeignThree mocks base method.
func (m *MockIndex) ForeignThree(arg0 imp3.Imp3) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForeignThree", arg0)
}

// ForeignThree indicates an expected call of ForeignThree.
func (mr *MockIndexMockRecorder) ForeignThree(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForeignThree", reflect.TypeOf((*MockIndex)(nil).ForeignThree), arg0)
}

// ForeignTwo mocks base method.
func (m *MockIndex) ForeignTwo(arg0 imp2.Imp2) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForeignTwo", arg0)
}

// ForeignTwo indicates an expected call of ForeignTwo.
func (mr *MockIndexMockRecorder) ForeignTwo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForeignTwo", reflect.TypeOf((*MockIndex)(nil).ForeignTwo), arg0)
}

// Func mocks base method.
func (m *MockIndex) Func(arg0 func(http.Request) (int, bool)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Func", arg0)
}

// Func indicates an expected call of Func.
func (mr *MockIndexMockRecorder) Func(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Func", reflect.TypeOf((*MockIndex)(nil).Func), arg0)
}

// Get mocks base method.
func (m *MockIndex) Get(arg0 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIndexMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIndex)(nil).Get), arg0)
}

// GetTwo mocks base method.
func (m *MockIndex) GetTwo(arg0, arg1 string) (interface{}, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwo", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// GetTwo indicates an expected call of GetTwo.
func (mr *MockIndexMockRecorder) GetTwo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwo", reflect.TypeOf((*MockIndex)(nil).GetTwo), arg0, arg1)
}

// Map mocks base method.
func (m *MockIndex) Map(arg0 map[int]hash.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Map", arg0)
}

// Map indicates an expected call of Map.
func (mr *MockIndexMockRecorder) Map(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockIndex)(nil).Map), arg0)
}

// NillableRet mocks base method.
func (m *MockIndex) NillableRet() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NillableRet")
	ret0, _ := ret[0].(error)
	return ret0
}

// NillableRet indicates an expected call of NillableRet.
func (mr *MockIndexMockRecorder) NillableRet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NillableRet", reflect.TypeOf((*MockIndex)(nil).NillableRet))
}

// Other mocks base method.
func (m *MockIndex) Other() hash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Other")
	ret0, _ := ret[0].(hash.Hash)
	return ret0
}

// Other indicates an expected call of Other.
func (mr *MockIndexMockRecorder) Other() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Other", reflect.TypeOf((*MockIndex)(nil).Other))
}

// Ptr mocks base method.
func (m *MockIndex) Ptr(arg0 *int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ptr", arg0)
}

// Ptr indicates an expected call of Ptr.
func (mr *MockIndexMockRecorder) Ptr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ptr", reflect.TypeOf((*MockIndex)(nil).Ptr), arg0)
}

// Put mocks base method.
func (m *MockIndex) Put(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put.
func (mr *MockIndexMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIndex)(nil).Put), arg0, arg1)
}

// Slice mocks base method.
func (m *MockIndex) Slice(arg0 []int, arg1 []byte) [3]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slice", arg0, arg1)
	ret0, _ := ret[0].([3]int)
	return ret0
}

// Slice indicates an expected call of Slice.
func (mr *MockIndexMockRecorder) Slice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slice", reflect.TypeOf((*MockIndex)(nil).Slice), arg0, arg1)
}

// Struct mocks base method.
func (m *MockIndex) Struct(arg0 struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Struct", arg0)
}

// Struct indicates an expected call of Struct.
func (mr *MockIndexMockRecorder) Struct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockIndex)(nil).Struct), arg0)
}

// StructChan mocks base method.
func (m *MockIndex) StructChan(arg0 chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StructChan", arg0)
}

// StructChan indicates an expected call of StructChan.
func (mr *MockIndexMockRecorder) StructChan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructChan", reflect.TypeOf((*MockIndex)(nil).StructChan), arg0)
}

// Summary mocks base method.
func (m *MockIndex) Summary(arg0 *bytes.Buffer, arg1 io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Summary", arg0, arg1)
}

// Summary indicates an expected call of Summary.
func (mr *MockIndexMockRecorder) Summary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Summary", reflect.TypeOf((*MockIndex)(nil).Summary), arg0, arg1)
}

// Templates mocks base method.
func (m *MockIndex) Templates(arg0 template.CSS, arg1 template0.FuncMap) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Templates", arg0, arg1)
}

// Templates indicates an expected call of Templates.
func (mr *MockIndexMockRecorder) Templates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Templates", reflect.TypeOf((*MockIndex)(nil).Templates), arg0, arg1)
}

// MockEmbed is a mock of Embed interface.
type MockEmbed struct {
	ctrl     *gomock.Controller
	recorder *MockEmbedMockRecorder
}

// MockEmbedMockRecorder is the mock recorder for MockEmbed.
type MockEmbedMockRecorder struct {
	mock *MockEmbed
}

// NewMockEmbed creates a new mock instance.
func NewMockEmbed(ctrl *gomock.Controller) *MockEmbed {
	mock := &MockEmbed{ctrl: ctrl}
	mock.recorder = &MockEmbedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbed) EXPECT() *MockEmbedMockRecorder {
	return m.recorder
}

// EmbeddedMethod mocks base method.
func (m *MockEmbed) EmbeddedMethod() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedMethod")
}

// EmbeddedMethod indicates an expected call of EmbeddedMethod.
func (mr *MockEmbedMockRecorder) EmbeddedMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedMethod", reflect.TypeOf((*MockEmbed)(nil).EmbeddedMethod))
}

// ForeignEmbeddedMethod mocks base method.
func (m *MockEmbed) ForeignEmbeddedMethod() *bufio.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForeignEmbeddedMethod")
	ret0, _ := ret[0].(*bufio.Reader)
	return ret0
}

// ForeignEmbeddedMethod indicates an expected call of ForeignEmbeddedMethod.
func (mr *MockEmbedMockRecorder) ForeignEmbeddedMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForeignEmbeddedMethod", reflect.TypeOf((*MockEmbed)(nil).ForeignEmbeddedMethod))
}

// ImplicitPackage mocks base method.
func (m *MockEmbed) ImplicitPackage(arg0 string, arg1 imp1.ImpT, arg2 []imp1.ImpT, arg3 *imp1.ImpT, arg4 chan imp1.ImpT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImplicitPackage", arg0, arg1, arg2, arg3, arg4)
}

// ImplicitPackage indicates an expected call of ImplicitPackage.
func (mr *MockEmbedMockRecorder) ImplicitPackage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImplicitPackage", reflect.TypeOf((*MockEmbed)(nil).ImplicitPackage), arg0, arg1, arg2, arg3, arg4)
}

// RegularMethod mocks base method.
func (m *MockEmbed) RegularMethod() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegularMethod")
}

// RegularMethod indicates an expected call of RegularMethod.
func (mr *MockEmbedMockRecorder) RegularMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegularMethod", reflect.TypeOf((*MockEmbed)(nil).RegularMethod))
}

// MockEmbedded is a mock of Embedded interface.
type MockEmbedded struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddedMockRecorder
}

// MockEmbeddedMockRecorder is the mock recorder for MockEmbedded.
type MockEmbeddedMockRecorder struct {
	mock *MockEmbedded
}

// NewMockEmbedded creates a new mock instance.
func NewMockEmbedded(ctrl *gomock.Controller) *MockEmbedded {
	mock := &MockEmbedded{ctrl: ctrl}
	mock.recorder = &MockEmbeddedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbedded) EXPECT() *MockEmbeddedMockRecorder {
	return m.recorder
}

// EmbeddedMethod mocks base method.
func (m *MockEmbedded) EmbeddedMethod() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedMethod")
}

// EmbeddedMethod indicates an expected call of EmbeddedMethod.
func (mr *MockEmbeddedMockRecorder) EmbeddedMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedMethod", reflect.TypeOf((*MockEmbedded)(nil).EmbeddedMethod))
}
