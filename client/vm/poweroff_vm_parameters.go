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

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// NewPoweroffVMParams creates a new PoweroffVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPoweroffVMParams() *PoweroffVMParams {
	return &PoweroffVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPoweroffVMParamsWithTimeout creates a new PoweroffVMParams object
// with the ability to set a timeout on a request.
func NewPoweroffVMParamsWithTimeout(timeout time.Duration) *PoweroffVMParams {
	return &PoweroffVMParams{
		timeout: timeout,
	}
}

// NewPoweroffVMParamsWithContext creates a new PoweroffVMParams object
// with the ability to set a context for a request.
func NewPoweroffVMParamsWithContext(ctx context.Context) *PoweroffVMParams {
	return &PoweroffVMParams{
		Context: ctx,
	}
}

// NewPoweroffVMParamsWithHTTPClient creates a new PoweroffVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewPoweroffVMParamsWithHTTPClient(client *http.Client) *PoweroffVMParams {
	return &PoweroffVMParams{
		HTTPClient: client,
	}
}

/* PoweroffVMParams contains all the parameters to send to the API endpoint
   for the poweroff Vm operation.

   Typically these are written to a http.Request.
*/
type PoweroffVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the poweroff Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoweroffVMParams) WithDefaults() *PoweroffVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the poweroff Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoweroffVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := PoweroffVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the poweroff Vm params
func (o *PoweroffVMParams) WithTimeout(timeout time.Duration) *PoweroffVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the poweroff Vm params
func (o *PoweroffVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the poweroff Vm params
func (o *PoweroffVMParams) WithContext(ctx context.Context) *PoweroffVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the poweroff Vm params
func (o *PoweroffVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the poweroff Vm params
func (o *PoweroffVMParams) WithHTTPClient(client *http.Client) *PoweroffVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the poweroff Vm params
func (o *PoweroffVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the poweroff Vm params
func (o *PoweroffVMParams) WithContentLanguage(contentLanguage *string) *PoweroffVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the poweroff Vm params
func (o *PoweroffVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the poweroff Vm params
func (o *PoweroffVMParams) WithRequestBody(requestBody *models.VMOperateParams) *PoweroffVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the poweroff Vm params
func (o *PoweroffVMParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PoweroffVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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