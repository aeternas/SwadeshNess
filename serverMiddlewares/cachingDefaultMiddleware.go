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
	log.Println(val)
	if err != nil {
		log.Printf("Failed to extract cached value %s", key)
		return r
	}
	bytes := []byte(val)
	log.Print(bytes)
	r.Cached = true
	r.Data = bytes
	return r
}

func (c cachingDefaultServerMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	key := c.GetKey(r.Request)
	cw := c.CW
	str := string(r.Data)
	if err := (*cw).SaveCachedValue(key, str); err != nil {
		log.Printf("Failed to save cached value %s", r.Data)
	}
	return r
}

func (cachingDefaultServerMiddleware) GetKey(r *apiClient.Request) string {
	log.Printf("RawQuery is:")
	log.Println(r.NetRequest.URL.RawQuery)
	return r.NetRequest.URL.RawQuery
}
