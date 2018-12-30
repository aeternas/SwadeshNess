package apiClient

import (
	"net/http"
)

type Response struct {
	Data        []byte
	NetResponse *http.Response
	Request     *Request
}
