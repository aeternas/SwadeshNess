package caching

type AnyCacheWrapper interface {
	GetCachedValue(k string) (string, error)
	SaveCachedValue(k, v string) error
}
