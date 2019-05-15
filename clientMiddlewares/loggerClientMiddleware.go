package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"log"
)

type loggerClientMiddleware struct{}

type LoggerClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewLoggerClientMiddleware() LoggerClientMiddleware {
	return &loggerClientMiddleware{}
}

func (loggerClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println("LoggerClientMiddleware Request:\n")
	log.Println(r.NetRequest)
	return r
}

func (loggerClientMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println("LoggerClientMiddleware Response:\n")
	log.Println(r.NetResponse)
	return r
}
