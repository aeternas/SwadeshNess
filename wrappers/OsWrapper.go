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
}

type OsWrapper struct{}

func (w *OsWrapper) GetEnv(k string) string {
	return os.Getenv(k)
}

func (w *OsWrapper) Open(n string) FileOpened {
	file, err := os.Open(n)
	return FileOpened{F: file, Err: err}
}
