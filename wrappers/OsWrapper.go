package wrappers

import (
	"os"
)

type AnyOsWrapper interface {
	GetEnv(k string) string
}

type MockOsWrapper struct {
	GetEnvArgs      string
	GetEnvWasCalled int
	GetEnvStub      string
}

type OsWrapper struct{}

func (w *OsWrapper) GetEnv(k string) string {
	return os.Getenv(k)
}

func (w *MockOsWrapper) GetEnv(k string) string {
	w.GetEnvWasCalled += 1
	w.GetEnvArgs = k
	return w.GetEnvStub
}
