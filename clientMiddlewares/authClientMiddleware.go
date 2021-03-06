package middlewares

import (
	apiClient "github.com/aeternas/SwadeshNess/apiClient"
	"log"
	"net/url"
)

type authClientMiddleware struct {
	apiKey string
}

type AuthClientMiddleware interface {
	AdaptRequest(r *apiClient.Request) *apiClient.Request
	AdaptResponse(r *apiClient.Response) *apiClient.Response
}

func NewAuthClientMiddleware(apiKey string) AuthClientMiddleware {
	return &authClientMiddleware{apiKey: apiKey}
}

func (a *authClientMiddleware) AdaptRequest(r *apiClient.Request) *apiClient.Request {
	urlString, err := url.Parse(r.NetRequest.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	q := urlString.Query()
	q.Add("key", a.apiKey)
	urlString.RawQuery = q.Encode()
	r.NetRequest.URL = urlString
	return r
}

func (authClientMiddleware) AdaptResponse(r *apiClient.Response) *apiClient.Response {
	return r
}
