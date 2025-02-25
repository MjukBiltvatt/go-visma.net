package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerCreditNoteGetAll() CustomerCreditNoteGetAll {
	r := CustomerCreditNoteGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerCreditNoteGetAll struct {
	client      *Client
	queryParams *CustomerCreditNoteGetAllQueryParams
	pathParams  *CustomerCreditNoteGetAllPathParams
	method      string
	headers     http.Header
	requestBody CustomerCreditNoteGetAllBody
}

func (r CustomerCreditNoteGetAll) NewQueryParams() *CustomerCreditNoteGetAllQueryParams {
	return &CustomerCreditNoteGetAllQueryParams{}
}

type CustomerCreditNoteGetAllQueryParams struct {
	// The field is deprecated for specific customer document endpoints. It will
	// only be usable from customer document endpoint.
	DocumentType string `schema:"documentType,omitempty"`

	// Parameter for showing if invoice has been released or not.
	Released int `schema:"released,omitempty"`

	// The dunning level of the document
	DunningLevel int `schema:"dunningLevel,omitempty"`

	// The date of the closing of the financial period.
	ClosedFinancialPeriod string `schema:"closedFinancialPeriod,omitempty"`

	// The date and time for when the document last released a dunning letter.
	DunningLetterDateTime string `schema:"dunningLetterDateTime,omitempty"`

	// Set time/date as before (<), after (>), before and including (=<) OR
	// after and including (=>) to filter on time frame.
	DunningLetterDateTimeCondition string `schema:"dunningLetterDateTimeCondition,omitempty"`

	// The project with which the document is associated.
	Project string `schema:"project,omitempty"`

	// True if you want to see all dunning information regarding this document.
	ExpandApplications bool `schema:"expandApplications,omitempty"`

	ExpandDunningInformation bool `schema:"expandDunningInformation,omitempty"`

	// True if you want to see all attachments regarding this document.
	ExpandAttachments bool `schema:"expandAttachments,omitempty"`

	// True if you want to see all VAT details regarding this document.
	ExpandTaxDetails bool `schema:"expandTaxDetails,omitempty"`

	// True if you want to see all information regarding the invoice address for
	// this document.
	ExpandInvoiceAddress bool `schema:"expandInvoiceAddress,omitempty"`

	// The financial period to which the transactions recorded in the document
	// is posted. Format YYYYMM.
	FinancialPeriod string `schema:"financialPeriod,omitempty"`

	// The date when payment for the document is due, in accordance with the
	// credit terms.
	DocumentDueDate string `schema:"documentDueDate,omitempty"`

	// The status of the document. Use the dropdown to select status.
	Status string `schema:"status,omitempty"`

	// This field has been deprecated and will be removed in future versions.
	// Use pagenumber and pagesize for pagination purposes. Pagenumber and
	// pagesize does not work with NumberToRead and SkipRecords.
	NumberToRead int `schema:"numberToRead,omitempty"`

	// This field has been deprecated and will be removed in future versions.
	// Use pagenumber and pagesize for pagination purposes. Pagenumber and
	// pagesize does not work with NumberToRead and SkipRecords.
	SkipRecords int `schema:"skipRecords,omitempty"`

	// The top part > External reference > The external reference used in
	// AutoInvoice.
	ExternalReference string `schema:"externalReference,omitempty"`

	// The top part > Payment ref. > The reference number of the document, as
	// automatically generated by the system in accordance with the number
	// series assigned to cash sales in the Customer ledger preferences window…
	PaymentReference string `schema:"paymentReference,omitempty"`

	// The top part > External reference > The external reference used in
	// AutoInvoice.
	CustomerRefNumber string `schema:"customerRefNumber,omitempty"`

	// Greater than value. The item which is the object for this, varies from
	// API to API.
	GreaterThanValue string `schema:"greaterThanValue,omitempty"`

	// System generated value for last modification of transaction/record. Use
	// format: YYYY-MM-DD HH:MM (date and time) to filter from date to present.
	LastModifiedDateTime string `schema:"lastModifiedDateTime,omitempty"`

	// System retrieved information for state/condition.
	LastModifiedDateTimeCondition string `schema:"lastModifiedDateTimeCondition,omitempty"`

	// Creation date and time.
	CreatedDateTime string `schema:"createdDateTime,omitempty"`

	// System-retrieved information for state/condition
	CreatedDateTimeCondition string `schema:"createdDateTimeCondition,omitempty"`

	// Pagination parameter. Page number.
	PageNumber int `schema:"pageNumber,omitempty"`

	// Pagination parameter. Number of items to be collected. Please use a page
	// size lower or equal to the allowed max page size which is returned as
	// part of the metadata information. If requested page size is greater than
	// allowed max page size, request will be limited to max page size.
	PageSize int `schema:"pageSize,omitempty"`
}

func (p CustomerCreditNoteGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerCreditNoteGetAll) QueryParams() *CustomerCreditNoteGetAllQueryParams {
	return r.queryParams
}

func (r CustomerCreditNoteGetAll) NewPathParams() *CustomerCreditNoteGetAllPathParams {
	return &CustomerCreditNoteGetAllPathParams{}
}

type CustomerCreditNoteGetAllPathParams struct {
}

func (p *CustomerCreditNoteGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerCreditNoteGetAll) PathParams() *CustomerCreditNoteGetAllPathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerCreditNoteGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CustomerCreditNoteGetAll) Method() string {
	return r.method
}

func (r CustomerCreditNoteGetAll) NewRequestBody() CustomerCreditNoteGetAllBody {
	return CustomerCreditNoteGetAllBody{}
}

type CustomerCreditNoteGetAllBody struct {
}

func (r *CustomerCreditNoteGetAll) RequestBody() *CustomerCreditNoteGetAllBody {
	return nil
}

func (r *CustomerCreditNoteGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerCreditNoteGetAll) SetRequestBody(body CustomerCreditNoteGetAllBody) {
	r.requestBody = body
}

func (r *CustomerCreditNoteGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customerCreditNote", r.PathParams())
	return &u
}

func (r *CustomerCreditNoteGetAll) Do() (resp CustomerCreditNoteGetAllResponse, err error) {
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

type CustomerCreditNoteGetAllResponse struct {
	Response
	Body CreditNotes
}
