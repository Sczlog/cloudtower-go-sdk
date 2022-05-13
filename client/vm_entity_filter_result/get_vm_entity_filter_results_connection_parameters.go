// Code generated by go-swagger; DO NOT EDIT.

package vm_entity_filter_result

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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewGetVMEntityFilterResultsConnectionParams creates a new GetVMEntityFilterResultsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMEntityFilterResultsConnectionParams() *GetVMEntityFilterResultsConnectionParams {
	return &GetVMEntityFilterResultsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMEntityFilterResultsConnectionParamsWithTimeout creates a new GetVMEntityFilterResultsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVMEntityFilterResultsConnectionParamsWithTimeout(timeout time.Duration) *GetVMEntityFilterResultsConnectionParams {
	return &GetVMEntityFilterResultsConnectionParams{
		timeout: timeout,
	}
}

// NewGetVMEntityFilterResultsConnectionParamsWithContext creates a new GetVMEntityFilterResultsConnectionParams object
// with the ability to set a context for a request.
func NewGetVMEntityFilterResultsConnectionParamsWithContext(ctx context.Context) *GetVMEntityFilterResultsConnectionParams {
	return &GetVMEntityFilterResultsConnectionParams{
		Context: ctx,
	}
}

// NewGetVMEntityFilterResultsConnectionParamsWithHTTPClient creates a new GetVMEntityFilterResultsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMEntityFilterResultsConnectionParamsWithHTTPClient(client *http.Client) *GetVMEntityFilterResultsConnectionParams {
	return &GetVMEntityFilterResultsConnectionParams{
		HTTPClient: client,
	}
}

/* GetVMEntityFilterResultsConnectionParams contains all the parameters to send to the API endpoint
   for the get Vm entity filter results connection operation.

   Typically these are written to a http.Request.
*/
type GetVMEntityFilterResultsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMEntityFilterResultsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm entity filter results connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMEntityFilterResultsConnectionParams) WithDefaults() *GetVMEntityFilterResultsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm entity filter results connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMEntityFilterResultsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMEntityFilterResultsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) WithTimeout(timeout time.Duration) *GetVMEntityFilterResultsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) WithContext(ctx context.Context) *GetVMEntityFilterResultsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) WithHTTPClient(client *http.Client) *GetVMEntityFilterResultsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) WithContentLanguage(contentLanguage *string) *GetVMEntityFilterResultsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) WithRequestBody(requestBody *models.GetVMEntityFilterResultsConnectionRequestBody) *GetVMEntityFilterResultsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm entity filter results connection params
func (o *GetVMEntityFilterResultsConnectionParams) SetRequestBody(requestBody *models.GetVMEntityFilterResultsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMEntityFilterResultsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
