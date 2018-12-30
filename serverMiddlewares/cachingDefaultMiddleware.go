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
	key := getKey(r)
	cw := c.CW
	val, err := (*cw).GetCachedValue(key)
	if err != nil {
		log.Printf("Failed to extract cached value %s", key)
		return r
	}
	bytes := []byte(val)
	r.Cached = true
	r.Data = bytes
	return r
}

func (c cachingDefaultServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	log.Println(r)
	return r
}

func getKey(r *apiClient.Request) string {
	return r.NetRequest.URL.RawQuery
}
