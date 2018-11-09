package apiClient

import (
	"net/http"
)

type Middleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}
