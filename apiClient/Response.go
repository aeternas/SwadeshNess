package apiClient

import (
	"net/http"
)

type Response struct {
	Data        interface{}
	NetResponse *http.Response
}
