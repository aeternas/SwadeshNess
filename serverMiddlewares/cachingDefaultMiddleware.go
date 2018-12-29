package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	"log"
)

type cachingDefaultServerMiddleware struct {
	CachingWrapper *Caching.AnyCacheWrapper
}

type CachingDefaultServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewCachingDefaultServerMiddleware() CachingDefaultServerMiddleware {
	return &cachingDefaultServerMiddleware{}
}

func (cachingDefaultServerMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println(r)
	return r
}

func (cachingDefaultServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println(r)
	return r
}
