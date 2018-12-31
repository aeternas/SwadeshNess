package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	Conf "github.com/aeternas/SwadeshNess/configuration"
	"log"
)

type cachingDefaultServerMiddleware struct {
	CW *Caching.AnyCacheWrapper
}

type CachingDefaultServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
	GetKey(r *apiClient.Request) string
}

func NewCachingDefaultServerMiddleware(c *Conf.Configuration) CachingDefaultServerMiddleware {
	cw := Caching.NewRedisCachingWrapper(c).(Caching.AnyCacheWrapper)
	return &cachingDefaultServerMiddleware{CW: &cw}
}

func (c cachingDefaultServerMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	key := c.GetKey(r)
	cw := c.CW
	val, err := (*cw).GetCachedValue(key)
	if err != nil || len(val) == 0 {
		log.Printf("Cache miss for %s", key)
		return r
	}
	bytes := []byte(val)
	log.Printf("Cache hit with %s ", val)
	r.Cached = true
	r.Data = bytes
	return r
}

func (c cachingDefaultServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	if r.Request.Cached {
		r.Data = r.Request.Data
		return r
	}
	key := c.GetKey(r.Request)
	cw := c.CW
	str := string(r.Data)
	log.Printf("Trying to save data %s for key %s", str, key)
	if err := (*cw).SaveCachedValue(key, str); err != nil {
		log.Printf("Failed to save value to cache: %s", r.Data)
	}
	return r
}

func (cachingDefaultServerMiddleware) GetKey(r *apiClient.Request) string {
	return r.NetRequest.URL.RawQuery
}
