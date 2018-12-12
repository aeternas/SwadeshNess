package caching

type AnyCacheWrapper interface {
	GetCachedValue(k string) string
	SaveCachedValue(k, v string)
}
