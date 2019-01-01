package middlewares

import (
	"fmt"
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	Conf "github.com/aeternas/SwadeshNess/configuration"
	"log"
)

type cachingDefaultClientMiddleware struct {
	CachingWrapper *Caching.AnyCacheWrapper
	Configuration  *Conf.Configuration
}

type CachingDefaultClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
	GetKey(r *apiClient.Request) string
}

func NewCachingDefaultClientMiddleware(c *Conf.Configuration) CachingDefaultClientMiddleware {
	cw := Caching.NewRedisCachingWrapper(c).(Caching.AnyCacheWrapper)
	return &cachingDefaultClientMiddleware{CachingWrapper: &cw, Configuration: c}
}

func (c cachingDefaultClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	key := c.GetKey(r)
	cw := c.CachingWrapper
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

func (c cachingDefaultClientMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	if r.Request.Cached {
		r.Data = r.Request.Data
		return r
	}
	key := c.GetKey(r.Request)
	cw := c.CachingWrapper
	str := string(r.Data)
	log.Printf("Trying to save data %s for key %s", str, key)
	if err := (*cw).SaveCachedValue(key, str); err != nil {
		log.Printf("Failed to save value to cache: %s", r.Data)
	}
	return r
}

func (c cachingDefaultClientMiddleware) GetKey(r *apiClient.Request) string {
	values := r.NetRequest.URL.Query()
	values.Del("key")
	encodedValues := values.Encode()
	version := fmt.Sprintf("&v=%s", c.Configuration.ConfigVersion)
	return encodedValues + version
}
