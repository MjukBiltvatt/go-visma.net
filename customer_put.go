package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerPut() CustomerPut {
	r := CustomerPut{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerPut struct {
	client      *Client
	queryParams *CustomerPutQueryParams
	pathParams  *CustomerPutPathParams
	method      string
	headers     http.Header
	requestBody CustomerPostBody
}

func (r CustomerPut) NewQueryParams() *CustomerPutQueryParams {
	return &CustomerPutQueryParams{}
}

type CustomerPutQueryParams struct{}

func (p CustomerPutQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerPut) QueryParams() *CustomerPutQueryParams {
	return r.queryParams
}

func (r CustomerPut) NewPathParams() *CustomerPutPathParams {
	return &CustomerPutPathParams{}
}

type CustomerPutPathParams struct {
	CustomerCD string `schema:"customer_cd"`
}

func (p *CustomerPutPathParams) Params() map[string]string {
	return map[string]string{
		"customer_cd": p.CustomerCD,
	}
}

func (r *CustomerPut) PathParams() *CustomerPutPathParams {
	return r.pathParams
}

func (r *CustomerPut) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerPut) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPut) Method() string {
	return r.method
}

func (r CustomerPut) NewRequestBody() CustomerPostBody {
	return CustomerPostBody{}
}

func (r *CustomerPut) RequestBody() *CustomerPostBody {
	return &r.requestBody
}

func (r *CustomerPut) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerPut) SetRequestBody(body CustomerPostBody) {
	r.requestBody = body
}

func (r *CustomerPut) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customer/{{.customer_cd}}", r.PathParams())
	return &u
}

func (r *CustomerPut) Do() (resp CustomerPutResponse, err error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return resp, err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return resp, err
	}

	resp.Http, err = r.client.Do(req, &resp.Body)
	return resp, err
}

type CustomerPutResponse struct {
	Response
	Body struct{}
}
