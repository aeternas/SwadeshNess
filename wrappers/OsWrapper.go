package wrappers

import (
	"os"
)

type AnyOsWrapper interface {
	GetEnv(k string) string
	Open(n string) (*os.File, error)
}

type OsWrapper struct{}

func (w *OsWrapper) GetEnv(k string) string {
	return os.Getenv(k)
}

func (w *OsWrapper) Open(n string) (*os.File, error) {
	return os.Open(n)
}
