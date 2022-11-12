package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
)

type ClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}
