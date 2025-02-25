package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewFinancialperiodGet() FinancialperiodGet {
	r := FinancialperiodGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FinancialperiodGet struct {
	client      *Client
	queryParams *FinancialperiodGetQueryParams
	pathParams  *FinancialperiodGetPathParams
	method      string
	headers     http.Header
	requestBody FinancialperiodGetBody
}

func (r FinancialperiodGet) NewQueryParams() *FinancialperiodGetQueryParams {
	return &FinancialperiodGetQueryParams{}
}

type FinancialperiodGetQueryParams struct {
}

func (p FinancialperiodGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FinancialperiodGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r FinancialperiodGet) NewPathParams() *FinancialperiodGetPathParams {
	return &FinancialperiodGetPathParams{}
}

type FinancialperiodGetPathParams struct {
}

func (p *FinancialperiodGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FinancialperiodGet) PathParams() *FinancialperiodGetPathParams {
	return r.pathParams
}

func (r *FinancialperiodGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FinancialperiodGet) SetMethod(method string) {
	r.method = method
}

func (r *FinancialperiodGet) Method() string {
	return r.method
}

func (r FinancialperiodGet) NewRequestBody() FinancialperiodGetBody {
	return FinancialperiodGetBody{}
}

type FinancialperiodGetBody struct {
}

func (r *FinancialperiodGet) RequestBody() *FinancialperiodGetBody {
	return nil
}

func (r *FinancialperiodGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *FinancialperiodGet) SetRequestBody(body FinancialperiodGetBody) {
	r.requestBody = body
}

func (r *FinancialperiodGet) URL() *url.URL {
	u := r.client.GetEndpointURL("controller/api/v1/financialPeriod", r.PathParams())
	return &u
}

func (r *FinancialperiodGet) Do() (resp FinancialperiodGetResponse, err error) {
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

type FinancialperiodGetResponse struct {
	Response
	Body FinancialPeriods
}
