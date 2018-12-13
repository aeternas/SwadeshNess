package middlewares

import (
	"log"
	"net/http"
)

type loggerMiddleware struct{}

type LoggerMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewLoggerMiddleware() LoggerMiddleware {
	return &loggerMiddleware{}
}

func (loggerMiddleware) AdaptRequest(r *http.Request) *http.Request {
	log.Println(r)
	return r
}

func (loggerMiddleware) AdaptResponse(r *http.Response) *http.Response {
	log.Println(r)
	return r
}
