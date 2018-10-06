package apiClient

import (
	"net/http"
)

type Request struct {
	Data interface{}

	NetRequest *http.Request
}
