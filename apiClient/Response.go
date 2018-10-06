package apiClient

import (
	"net/http"
)

type Response struct {
	NetResponse *http.Response
	Data        interface{}
}
