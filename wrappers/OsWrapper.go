package wrappers

import (
	"os"
)

type AnyOsWrapper interface {
	GetEnv(k string) string
}

type OsWrapper struct{}

func (w *OsWrapper) GetEnv(k string) string {
	return os.Getenv(k)
}
