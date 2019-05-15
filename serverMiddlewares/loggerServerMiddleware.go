package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"log"
)

type loggerServerMiddleware struct{}

type LoggerServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewLoggerServerMiddleware() LoggerServerMiddleware {
	return &loggerServerMiddleware{}
}

func (loggerServerMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println("Logged")
	log.Println(r.NetRequest)
	return r
}

func (loggerServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println("Logged")
	log.Println(r.NetResponse)
	return r
}
