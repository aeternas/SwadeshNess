package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"net/http"
)

type defaultClientMiddleware struct{}

type DefaultClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewDefaultClientMiddleware() DefaultClientMiddleware {
	return defaultClientMiddleware{}
}

func (defaultClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	r.NetRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (defaultClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	return r
}
