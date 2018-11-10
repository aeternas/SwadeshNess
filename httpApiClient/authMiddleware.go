package httpApiClient

import (
	"log"
	"net/http"
	"net/url"
)

type authMiddleware struct {
	apiKey string
}

type AuthMiddleware interface {
	AdaptRequest(r *http.Request) *http.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewAuthMiddleware(apiKey string) AuthMiddleware {
	return &authMiddleware{apiKey: apiKey}
}

func (a *authMiddleware) AdaptRequest(r *http.Request) *http.Request {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("key", a.apiKey)
	u.RawQuery = q.Encode()
	r.URL = u
	return r
}

func (authMiddleware) AdaptResponse(r *http.Response) *http.Response {
	return r
}
