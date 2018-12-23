package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"net/http"
)

type ClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *http.Response) *http.Response
}
