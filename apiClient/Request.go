package apiClient

import (
	"net/http"
)

type Request struct {
	Data       interface{}
	Cached     bool
	NetRequest *http.Request
}
