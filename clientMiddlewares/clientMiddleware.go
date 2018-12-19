package middlewares

import (
	"net/http"
)

type ClientMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}
