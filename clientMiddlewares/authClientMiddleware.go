package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"log"
	"net/http"
	"net/url"
)

type authClientMiddleware struct {
	apiKey string
}

type AuthClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *http.Response) *http.Response
}

func NewAuthClientMiddleware(apiKey string) AuthClientMiddleware {
	return &authClientMiddleware{apiKey: apiKey}
}

func (a *authClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	u, err := url.Parse(r.NetRequest.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("key", a.apiKey)
	u.RawQuery = q.Encode()
	r.NetRequest.URL = u
	return r
}

func (authClientMiddleware) AdaptResponse(r *http.Response) *http.Response {
	return r
}
