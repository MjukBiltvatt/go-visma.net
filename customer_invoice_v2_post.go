package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/omitempty"
	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerInvoiceV2Post() CustomerInvoiceV2Post {
	r := CustomerInvoiceV2Post{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerInvoiceV2Post struct {
	client      *Client
	queryParams *CustomerInvoiceV2PostQueryParams
	pathParams  *CustomerInvoiceV2PostPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceV2PostBody
}

func (r CustomerInvoiceV2Post) NewQueryParams() *CustomerInvoiceV2PostQueryParams {
	return &CustomerInvoiceV2PostQueryParams{}
}

type CustomerInvoiceV2PostQueryParams struct{}

func (p CustomerInvoiceV2PostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerInvoiceV2Post) QueryParams() *CustomerInvoiceV2PostQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceV2Post) NewPathParams() *CustomerInvoiceV2PostPathParams {
	return &CustomerInvoiceV2PostPathParams{}
}

type CustomerInvoiceV2PostPathParams struct {
}

func (p *CustomerInvoiceV2PostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerInvoiceV2Post) PathParams() *CustomerInvoiceV2PostPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceV2Post) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerInvoiceV2Post) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceV2Post) Method() string {
	return r.method
}

func (r CustomerInvoiceV2Post) NewRequestBody() CustomerInvoiceV2PostBody {
	return CustomerInvoiceV2PostBody{}
}

type CustomerInvoiceV2PostBody NewInvoice

func (r CustomerInvoiceV2PostBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *CustomerInvoiceV2Post) RequestBody() *CustomerInvoiceV2PostBody {
	return &r.requestBody
}

func (r *CustomerInvoiceV2Post) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerInvoiceV2Post) SetRequestBody(body CustomerInvoiceV2PostBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceV2Post) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v2/customerinvoice", r.PathParams())
	return &u
}

func (r *CustomerInvoiceV2Post) Do() (resp CustomerInvoiceV2PostResponse, err error) {
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

type CustomerInvoiceV2PostResponse struct {
	Response
	Body JournalTransactions
}
