package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"log"
	"net/http"
)

type loggerClientMiddleware struct{}

type LoggerClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewLoggerClientMiddleware() LoggerClientMiddleware {
	return &loggerClientMiddleware{}
}

func (loggerClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println(r.NetRequest)
	return r
}

func (loggerClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	log.Println(r)
	return r
}
