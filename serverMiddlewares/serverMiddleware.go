package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
)

type ServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(w *apiClient.Response) *apiClient.Response
}
