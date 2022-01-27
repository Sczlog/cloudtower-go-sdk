// Code generated by go-swagger; DO NOT EDIT.

package vlan

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

// NewCreateVMVlanParams creates a new CreateVMVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMVlanParams() *CreateVMVlanParams {
	return &CreateVMVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMVlanParamsWithTimeout creates a new CreateVMVlanParams object
// with the ability to set a timeout on a request.
func NewCreateVMVlanParamsWithTimeout(timeout time.Duration) *CreateVMVlanParams {
	return &CreateVMVlanParams{
		timeout: timeout,
	}
}

// NewCreateVMVlanParamsWithContext creates a new CreateVMVlanParams object
// with the ability to set a context for a request.
func NewCreateVMVlanParamsWithContext(ctx context.Context) *CreateVMVlanParams {
	return &CreateVMVlanParams{
		Context: ctx,
	}
}

// NewCreateVMVlanParamsWithHTTPClient creates a new CreateVMVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMVlanParamsWithHTTPClient(client *http.Client) *CreateVMVlanParams {
	return &CreateVMVlanParams{
		HTTPClient: client,
	}
}

/* CreateVMVlanParams contains all the parameters to send to the API endpoint
   for the create Vm vlan operation.

   Typically these are written to a http.Request.
*/
type CreateVMVlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VMVlanCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Vm vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMVlanParams) WithDefaults() *CreateVMVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Vm vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMVlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVMVlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create Vm vlan params
func (o *CreateVMVlanParams) WithTimeout(timeout time.Duration) *CreateVMVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Vm vlan params
func (o *CreateVMVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Vm vlan params
func (o *CreateVMVlanParams) WithContext(ctx context.Context) *CreateVMVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Vm vlan params
func (o *CreateVMVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Vm vlan params
func (o *CreateVMVlanParams) WithHTTPClient(client *http.Client) *CreateVMVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Vm vlan params
func (o *CreateVMVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create Vm vlan params
func (o *CreateVMVlanParams) WithContentLanguage(contentLanguage *string) *CreateVMVlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create Vm vlan params
func (o *CreateVMVlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create Vm vlan params
func (o *CreateVMVlanParams) WithRequestBody(requestBody []*models.VMVlanCreationParams) *CreateVMVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create Vm vlan params
func (o *CreateVMVlanParams) SetRequestBody(requestBody []*models.VMVlanCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
