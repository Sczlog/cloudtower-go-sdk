// Code generated by go-swagger; DO NOT EDIT.

package iscsi_target

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

// NewCreateIscsiTargetParams creates a new CreateIscsiTargetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIscsiTargetParams() *CreateIscsiTargetParams {
	return &CreateIscsiTargetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIscsiTargetParamsWithTimeout creates a new CreateIscsiTargetParams object
// with the ability to set a timeout on a request.
func NewCreateIscsiTargetParamsWithTimeout(timeout time.Duration) *CreateIscsiTargetParams {
	return &CreateIscsiTargetParams{
		timeout: timeout,
	}
}

// NewCreateIscsiTargetParamsWithContext creates a new CreateIscsiTargetParams object
// with the ability to set a context for a request.
func NewCreateIscsiTargetParamsWithContext(ctx context.Context) *CreateIscsiTargetParams {
	return &CreateIscsiTargetParams{
		Context: ctx,
	}
}

// NewCreateIscsiTargetParamsWithHTTPClient creates a new CreateIscsiTargetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIscsiTargetParamsWithHTTPClient(client *http.Client) *CreateIscsiTargetParams {
	return &CreateIscsiTargetParams{
		HTTPClient: client,
	}
}

/* CreateIscsiTargetParams contains all the parameters to send to the API endpoint
   for the create iscsi target operation.

   Typically these are written to a http.Request.
*/
type CreateIscsiTargetParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.IscsiTargetCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create iscsi target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIscsiTargetParams) WithDefaults() *CreateIscsiTargetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create iscsi target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIscsiTargetParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateIscsiTargetParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create iscsi target params
func (o *CreateIscsiTargetParams) WithTimeout(timeout time.Duration) *CreateIscsiTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create iscsi target params
func (o *CreateIscsiTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create iscsi target params
func (o *CreateIscsiTargetParams) WithContext(ctx context.Context) *CreateIscsiTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create iscsi target params
func (o *CreateIscsiTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create iscsi target params
func (o *CreateIscsiTargetParams) WithHTTPClient(client *http.Client) *CreateIscsiTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create iscsi target params
func (o *CreateIscsiTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create iscsi target params
func (o *CreateIscsiTargetParams) WithContentLanguage(contentLanguage *string) *CreateIscsiTargetParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create iscsi target params
func (o *CreateIscsiTargetParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create iscsi target params
func (o *CreateIscsiTargetParams) WithRequestBody(requestBody []*models.IscsiTargetCreationParams) *CreateIscsiTargetParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create iscsi target params
func (o *CreateIscsiTargetParams) SetRequestBody(requestBody []*models.IscsiTargetCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIscsiTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
