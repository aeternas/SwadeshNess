package wrappers

import (
	"net/http"
)

type AnyResponseWriterWrapper interface {
	Write([]byte) (int, error)
}

type ResponseWriterWrapper struct {
	Writer *http.ResponseWriter
}
