package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	"log"
)

type cachingDefaultServerMiddleware struct {
	CW *Caching.AnyCacheWrapper
}

type CachingDefaultServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewCachingDefaultServerMiddleware() CachingDefaultServerMiddleware {
	cw := Caching.NewRedisCachingWrapper().(Caching.AnyCacheWrapper)
	return &cachingDefaultServerMiddleware{CW: &cw}
}

func (c cachingDefaultServerMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println(r)
	return r
}

func (c cachingDefaultServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println(r)
	return r
}
