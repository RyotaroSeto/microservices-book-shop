// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/book_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "gihyo/catalogue/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// GetBook mocks base method.
func (m *MockBookRepository) GetBook(ctx context.Context, id int) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBook", ctx, id)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBook indicates an expected call of GetBook.
func (mr *MockBookRepositoryMockRecorder) GetBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookRepository)(nil).GetBook), ctx, id)
}

// ListBooks mocks base method.
func (m *MockBookRepository) ListBooks(ctx context.Context) ([]*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooks", ctx)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooks indicates an expected call of ListBooks.
func (mr *MockBookRepositoryMockRecorder) ListBooks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooks", reflect.TypeOf((*MockBookRepository)(nil).ListBooks), ctx)
}
