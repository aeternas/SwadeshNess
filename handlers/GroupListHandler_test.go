package handlers

import (
	"errors"
	"testing"
)

type MockResponseWriterWrapper struct {
	WriteWasCalled int
}

func (w *MockResponseWriterWrapper) Write([]byte) (int, error) {
	w.WriteWasCalled += 1
	return 0, errors.New("WriteError")
}

func TestGetGroupsGroupListHandler(t *testing.T) {
	_ = new(MockResponseWriterWrapper)
	t.SkipNow()
}
