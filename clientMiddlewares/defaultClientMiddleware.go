package middlewares

import (
	"net/http"
)

type defaultClientMiddleware struct{}

type DefaultClientMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewDefaultClientMiddleware() DefaultClientMiddleware {
	return defaultClientMiddleware{}
}

func (defaultClientMiddleware) AdaptRequest(r *http.Request) *http.Request {
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (defaultClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	return r
}
