// Code generated by go-swagger; DO NOT EDIT.

package vm_template

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

// NewGetVMTemplatesParams creates a new GetVMTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMTemplatesParams() *GetVMTemplatesParams {
	return &GetVMTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMTemplatesParamsWithTimeout creates a new GetVMTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetVMTemplatesParamsWithTimeout(timeout time.Duration) *GetVMTemplatesParams {
	return &GetVMTemplatesParams{
		timeout: timeout,
	}
}

// NewGetVMTemplatesParamsWithContext creates a new GetVMTemplatesParams object
// with the ability to set a context for a request.
func NewGetVMTemplatesParamsWithContext(ctx context.Context) *GetVMTemplatesParams {
	return &GetVMTemplatesParams{
		Context: ctx,
	}
}

// NewGetVMTemplatesParamsWithHTTPClient creates a new GetVMTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMTemplatesParamsWithHTTPClient(client *http.Client) *GetVMTemplatesParams {
	return &GetVMTemplatesParams{
		HTTPClient: client,
	}
}

/* GetVMTemplatesParams contains all the parameters to send to the API endpoint
   for the get Vm templates operation.

   Typically these are written to a http.Request.
*/
type GetVMTemplatesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMTemplatesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMTemplatesParams) WithDefaults() *GetVMTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMTemplatesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMTemplatesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm templates params
func (o *GetVMTemplatesParams) WithTimeout(timeout time.Duration) *GetVMTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm templates params
func (o *GetVMTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm templates params
func (o *GetVMTemplatesParams) WithContext(ctx context.Context) *GetVMTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm templates params
func (o *GetVMTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm templates params
func (o *GetVMTemplatesParams) WithHTTPClient(client *http.Client) *GetVMTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm templates params
func (o *GetVMTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm templates params
func (o *GetVMTemplatesParams) WithContentLanguage(contentLanguage *string) *GetVMTemplatesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm templates params
func (o *GetVMTemplatesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm templates params
func (o *GetVMTemplatesParams) WithRequestBody(requestBody *models.GetVMTemplatesRequestBody) *GetVMTemplatesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm templates params
func (o *GetVMTemplatesParams) SetRequestBody(requestBody *models.GetVMTemplatesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
