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

// NewUpdateVMHostOptionsParams creates a new UpdateVMHostOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMHostOptionsParams() *UpdateVMHostOptionsParams {
	return &UpdateVMHostOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMHostOptionsParamsWithTimeout creates a new UpdateVMHostOptionsParams object
// with the ability to set a timeout on a request.
func NewUpdateVMHostOptionsParamsWithTimeout(timeout time.Duration) *UpdateVMHostOptionsParams {
	return &UpdateVMHostOptionsParams{
		timeout: timeout,
	}
}

// NewUpdateVMHostOptionsParamsWithContext creates a new UpdateVMHostOptionsParams object
// with the ability to set a context for a request.
func NewUpdateVMHostOptionsParamsWithContext(ctx context.Context) *UpdateVMHostOptionsParams {
	return &UpdateVMHostOptionsParams{
		Context: ctx,
	}
}

// NewUpdateVMHostOptionsParamsWithHTTPClient creates a new UpdateVMHostOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMHostOptionsParamsWithHTTPClient(client *http.Client) *UpdateVMHostOptionsParams {
	return &UpdateVMHostOptionsParams{
		HTTPClient: client,
	}
}

/* UpdateVMHostOptionsParams contains all the parameters to send to the API endpoint
   for the update Vm host options operation.

   Typically these are written to a http.Request.
*/
type UpdateVMHostOptionsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateHostOptionsParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm host options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMHostOptionsParams) WithDefaults() *UpdateVMHostOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm host options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMHostOptionsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMHostOptionsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm host options params
func (o *UpdateVMHostOptionsParams) WithTimeout(timeout time.Duration) *UpdateVMHostOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm host options params
func (o *UpdateVMHostOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm host options params
func (o *UpdateVMHostOptionsParams) WithContext(ctx context.Context) *UpdateVMHostOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm host options params
func (o *UpdateVMHostOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm host options params
func (o *UpdateVMHostOptionsParams) WithHTTPClient(client *http.Client) *UpdateVMHostOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm host options params
func (o *UpdateVMHostOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm host options params
func (o *UpdateVMHostOptionsParams) WithContentLanguage(contentLanguage *string) *UpdateVMHostOptionsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm host options params
func (o *UpdateVMHostOptionsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm host options params
func (o *UpdateVMHostOptionsParams) WithRequestBody(requestBody *models.VMUpdateHostOptionsParams) *UpdateVMHostOptionsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm host options params
func (o *UpdateVMHostOptionsParams) SetRequestBody(requestBody *models.VMUpdateHostOptionsParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMHostOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
