package middlewares

import (
	"fmt"
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	Caching "github.com/aeternas/SwadeshNess/caching"
	Conf "github.com/aeternas/SwadeshNess/configuration"
	"log"
)

type cachingDefaultServerMiddleware struct {
	CW            *Caching.AnyCacheWrapper
	Configuration *Conf.Configuration
}

type CachingDefaultServerMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
	GetKey(r *apiClient.Request) string
}

func NewCachingDefaultServerMiddleware(c *Conf.Configuration) CachingDefaultServerMiddleware {
	cw := Caching.NewRedisCachingWrapper(c).(Caching.AnyCacheWrapper)
	return &cachingDefaultServerMiddleware{CW: &cw, Configuration: c}
}

func (c cachingDefaultServerMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	key := c.GetKey(r)
	cw := c.CW
	val, err := (*cw).GetCachedValue(key)
	if err != nil || len(val) == 0 {
		log.Printf("Cache retrieving error for key %s: %s", key, err)
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
		log.Println("Failed to save value to cache: ", r.Data)
	}
	return r
}

func (c cachingDefaultServerMiddleware) GetKey(r *apiClient.Request) string {
	version := fmt.Sprintf("&v=%s", c.Configuration.ConfigVersion)
	return r.NetRequest.URL.RawQuery + version
}
