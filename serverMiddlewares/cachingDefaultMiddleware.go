package middlewares

import (
	"log"
	"net/http"
)

type cachingDefaultServerMiddleware struct{}

type CachingDefaultServerMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewCachingDefaultServerMiddleware() CachingDefaultServerMiddleware {
	return &cachingDefaultServerMiddleware{}
}

func (cachingDefaultServerMiddleware) AdaptRequest(r *http.Request) *http.Request {
	log.Println(r)
	return r
}

func (cachingDefaultServerMiddleware) AdaptResponse(r *http.Response) *http.Response {
	log.Println(r)
	return r
}
