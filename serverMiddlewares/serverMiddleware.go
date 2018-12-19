package middlewares

import (
	"net/http"
)

type ServerMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponseWriter(w *http.ResponseWriter) *http.ResponseWriter
}
