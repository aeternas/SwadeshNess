package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
)

type defaultClientMiddleware struct{}

type DefaultClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewDefaultClientMiddleware() DefaultClientMiddleware {
	return defaultClientMiddleware{}
}

func (defaultClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	r.NetRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func (defaultClientMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	return r
}
