// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewRemoveVMCdRomParams creates a new RemoveVMCdRomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveVMCdRomParams() *RemoveVMCdRomParams {
	return &RemoveVMCdRomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveVMCdRomParamsWithTimeout creates a new RemoveVMCdRomParams object
// with the ability to set a timeout on a request.
func NewRemoveVMCdRomParamsWithTimeout(timeout time.Duration) *RemoveVMCdRomParams {
	return &RemoveVMCdRomParams{
		timeout: timeout,
	}
}

// NewRemoveVMCdRomParamsWithContext creates a new RemoveVMCdRomParams object
// with the ability to set a context for a request.
func NewRemoveVMCdRomParamsWithContext(ctx context.Context) *RemoveVMCdRomParams {
	return &RemoveVMCdRomParams{
		Context: ctx,
	}
}

// NewRemoveVMCdRomParamsWithHTTPClient creates a new RemoveVMCdRomParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveVMCdRomParamsWithHTTPClient(client *http.Client) *RemoveVMCdRomParams {
	return &RemoveVMCdRomParams{
		HTTPClient: client,
	}
}

/* RemoveVMCdRomParams contains all the parameters to send to the API endpoint
   for the remove Vm cd rom operation.

   Typically these are written to a http.Request.
*/
type RemoveVMCdRomParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMRemoveCdRomParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove Vm cd rom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMCdRomParams) WithDefaults() *RemoveVMCdRomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove Vm cd rom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMCdRomParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RemoveVMCdRomParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) WithTimeout(timeout time.Duration) *RemoveVMCdRomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) WithContext(ctx context.Context) *RemoveVMCdRomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) WithHTTPClient(client *http.Client) *RemoveVMCdRomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) WithContentLanguage(contentLanguage *string) *RemoveVMCdRomParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) WithRequestBody(requestBody *models.VMRemoveCdRomParams) *RemoveVMCdRomParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the remove Vm cd rom params
func (o *RemoveVMCdRomParams) SetRequestBody(requestBody *models.VMRemoveCdRomParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveVMCdRomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
