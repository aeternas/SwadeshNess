package middlewares

import (
	"log"
	"net/http"
)

type loggerClientMiddleware struct{}

type LoggerClientMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewLoggerClientMiddleware() LoggerClientMiddleware {
	return &loggerClientMiddleware{}
}

func (loggerClientMiddleware) AdaptRequest(r *http.Request) *http.Request {
	log.Println(r)
	return r
}

func (loggerClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	log.Println(r)
	return r
}
