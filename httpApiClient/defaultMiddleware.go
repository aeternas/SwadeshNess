package httpApiClient

import (
	. "github.com/aeternas/SwadeshNess/apiClient"
	"net/http"
)

type defaultMiddleware struct {
	Middleware
}

type DefaultMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewDefaultMiddleware() DefaultMiddleware {
	return defaultMiddleware{}
}

func (defaultMiddleware) AdaptRequest(r *http.Request) *http.Request {
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (defaultMiddleware) AdaptResponse(r *http.Response) *http.Response {
	return r
}
