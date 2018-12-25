package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	"log"
)

type cachingDefaultClientMiddleware struct {
	CachingWrapper *Caching.AnyCacheWrapper
}

type CachingDefaultClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewCachingDefaultClientMiddleware() CachingDefaultClientMiddleware {
	return &cachingDefaultClientMiddleware{}
}

func (cachingDefaultClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	log.Println(r)
	return r
}

func (cachingDefaultClientMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println(r)
	return r
}
