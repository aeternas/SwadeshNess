package apiClient

import (
	"net/http"
)

type Request struct {
	Data       []byte
	Cached     bool
	NetRequest *http.Request
}
