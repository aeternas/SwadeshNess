package middlewares

import (
	Caching "github.com/aeternas/SwadeshNess/caching"
	"log"
	"net/http"
)

type cachingDefaultClientMiddleware struct {
	CachingWrapper *Caching.AnyCacheWrapper
}

type CachingDefaultClientMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewCachingDefaultClientMiddleware() CachingDefaultClientMiddleware {
	return &cachingDefaultClientMiddleware{}
}

func (cachingDefaultClientMiddleware) AdaptRequest(r *http.Request) *http.Request {
	log.Println(r)
	return r
}

func (cachingDefaultClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	log.Println(r)
	return r
}
