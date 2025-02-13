package vismanet

import "net/http"

type Response struct {
	Http *http.Response
}

// StatusCode returns the http status
func (r *Response) Status() string {
	return r.Http.Status
}

// StatusCode returns the http status code
func (r *Response) StatusCode() int {
	return r.Http.StatusCode
}

// IsStatus determines if the http status code is equal to the given code
func (r *Response) IsStatus(code int) bool {
	return r.StatusCode() == code
}
