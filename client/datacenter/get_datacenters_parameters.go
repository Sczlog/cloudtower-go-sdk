// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewGetDatacentersParams creates a new GetDatacentersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatacentersParams() *GetDatacentersParams {
	return &GetDatacentersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatacentersParamsWithTimeout creates a new GetDatacentersParams object
// with the ability to set a timeout on a request.
func NewGetDatacentersParamsWithTimeout(timeout time.Duration) *GetDatacentersParams {
	return &GetDatacentersParams{
		timeout: timeout,
	}
}

// NewGetDatacentersParamsWithContext creates a new GetDatacentersParams object
// with the ability to set a context for a request.
func NewGetDatacentersParamsWithContext(ctx context.Context) *GetDatacentersParams {
	return &GetDatacentersParams{
		Context: ctx,
	}
}

// NewGetDatacentersParamsWithHTTPClient creates a new GetDatacentersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatacentersParamsWithHTTPClient(client *http.Client) *GetDatacentersParams {
	return &GetDatacentersParams{
		HTTPClient: client,
	}
}

/* GetDatacentersParams contains all the parameters to send to the API endpoint
   for the get datacenters operation.

   Typically these are written to a http.Request.
*/
type GetDatacentersParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetDatacentersRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datacenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatacentersParams) WithDefaults() *GetDatacentersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datacenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatacentersParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetDatacentersParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get datacenters params
func (o *GetDatacentersParams) WithTimeout(timeout time.Duration) *GetDatacentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datacenters params
func (o *GetDatacentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datacenters params
func (o *GetDatacentersParams) WithContext(ctx context.Context) *GetDatacentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datacenters params
func (o *GetDatacentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datacenters params
func (o *GetDatacentersParams) WithHTTPClient(client *http.Client) *GetDatacentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datacenters params
func (o *GetDatacentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get datacenters params
func (o *GetDatacentersParams) WithContentLanguage(contentLanguage *string) *GetDatacentersParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get datacenters params
func (o *GetDatacentersParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get datacenters params
func (o *GetDatacentersParams) WithRequestBody(requestBody *models.GetDatacentersRequestBody) *GetDatacentersParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get datacenters params
func (o *GetDatacentersParams) SetRequestBody(requestBody *models.GetDatacentersRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatacentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
