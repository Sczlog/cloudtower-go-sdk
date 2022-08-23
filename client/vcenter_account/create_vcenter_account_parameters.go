// Code generated by go-swagger; DO NOT EDIT.

package vcenter_account

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

// NewCreateVcenterAccountParams creates a new CreateVcenterAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVcenterAccountParams() *CreateVcenterAccountParams {
	return &CreateVcenterAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVcenterAccountParamsWithTimeout creates a new CreateVcenterAccountParams object
// with the ability to set a timeout on a request.
func NewCreateVcenterAccountParamsWithTimeout(timeout time.Duration) *CreateVcenterAccountParams {
	return &CreateVcenterAccountParams{
		timeout: timeout,
	}
}

// NewCreateVcenterAccountParamsWithContext creates a new CreateVcenterAccountParams object
// with the ability to set a context for a request.
func NewCreateVcenterAccountParamsWithContext(ctx context.Context) *CreateVcenterAccountParams {
	return &CreateVcenterAccountParams{
		Context: ctx,
	}
}

// NewCreateVcenterAccountParamsWithHTTPClient creates a new CreateVcenterAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVcenterAccountParamsWithHTTPClient(client *http.Client) *CreateVcenterAccountParams {
	return &CreateVcenterAccountParams{
		HTTPClient: client,
	}
}

/* CreateVcenterAccountParams contains all the parameters to send to the API endpoint
   for the create vcenter account operation.

   Typically these are written to a http.Request.
*/
type CreateVcenterAccountParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.CreateVcenterAccountParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vcenter account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVcenterAccountParams) WithDefaults() *CreateVcenterAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vcenter account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVcenterAccountParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVcenterAccountParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create vcenter account params
func (o *CreateVcenterAccountParams) WithTimeout(timeout time.Duration) *CreateVcenterAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vcenter account params
func (o *CreateVcenterAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vcenter account params
func (o *CreateVcenterAccountParams) WithContext(ctx context.Context) *CreateVcenterAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vcenter account params
func (o *CreateVcenterAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vcenter account params
func (o *CreateVcenterAccountParams) WithHTTPClient(client *http.Client) *CreateVcenterAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vcenter account params
func (o *CreateVcenterAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create vcenter account params
func (o *CreateVcenterAccountParams) WithContentLanguage(contentLanguage *string) *CreateVcenterAccountParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create vcenter account params
func (o *CreateVcenterAccountParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create vcenter account params
func (o *CreateVcenterAccountParams) WithRequestBody(requestBody *models.CreateVcenterAccountParams) *CreateVcenterAccountParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create vcenter account params
func (o *CreateVcenterAccountParams) SetRequestBody(requestBody *models.CreateVcenterAccountParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVcenterAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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