package wrappers

import (
	"os"
)

type FileOpened struct {
	F   *os.File
	Err error
}

type AnyOsWrapper interface {
	GetEnv(k string) string
	Open(n string) FileOpened
	GetEnvFallback(k, f string) string
	LookupEnv(key string) (string, bool)
}

type OsWrapper struct {
	NestedOsWrapper AnyOsWrapper
}

func New(w AnyOsWrapper) AnyOsWrapper {
	return &OsWrapper{NestedOsWrapper: w}
}

func (w *OsWrapper) GetEnv(k string) string {
	return os.Getenv(k)
}

func (w *OsWrapper) GetEnvFallback(k, f string) string {
	if w.NestedOsWrapper == nil {
		panic("OsWrapper was initialized improperly")
	}
	if value, ok := w.NestedOsWrapper.LookupEnv(k); ok {
		return value
	}
	return f
}

func (w *OsWrapper) Open(n string) FileOpened {
	file, err := os.Open(n)
	return FileOpened{F: file, Err: err}
}

func (w *OsWrapper) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

type MockOsWrapper struct {
	GetEnvArgs      []string
	GetEnvWasCalled int
	GetEnvStub      string

	GetEnvFallbackArgs      []string
	GetEnvFallbackWasCalled int
	GetEnvFallbackStub      string

	OpenArgs      string
	OpenWasCalled int
	OpenStub      FileOpened

	LookupEnvWasCalled int
	LookupEnvArgs      []string
	LookupEnvStub      string
	LookupEnvStatus    bool
}

func (w *MockOsWrapper) GetEnv(k string) string {
	w.GetEnvWasCalled += 1
	w.GetEnvArgs = append(w.GetEnvArgs, k)
	return w.GetEnvStub
}

func (w *MockOsWrapper) GetEnvFallback(k, f string) string {
	w.GetEnvFallbackWasCalled += 1
	w.GetEnvFallbackArgs = append(w.GetEnvFallbackArgs, k)
	w.GetEnvFallbackArgs = append(w.GetEnvFallbackArgs, f)
	return w.GetEnvFallbackStub
}

func (w *MockOsWrapper) Open(n string) FileOpened {
	w.OpenWasCalled += 1
	w.OpenArgs = n
	return w.OpenStub
}

func (w *MockOsWrapper) LookupEnv(key string) (string, bool) {
	w.LookupEnvWasCalled += 1
	w.LookupEnvArgs = append(w.LookupEnvArgs, key)
	return w.LookupEnvStub, w.LookupEnvStatus
}
