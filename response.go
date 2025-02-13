package vismanet

import "net/http"

type Response struct {
	Http *http.Response
}

// NotFound determines if the response status is 404 Not Found
func (r *Response) NotFound() bool {
	return r.Http.StatusCode == http.StatusNotFound
}
