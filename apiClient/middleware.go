package apiClient

type Middleware interface {
	AdaptRequest(r *Request)
	AdaptResponse(r *Response)
}
