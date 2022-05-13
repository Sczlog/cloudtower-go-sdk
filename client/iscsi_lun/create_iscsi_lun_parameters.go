// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun

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

// NewCreateIscsiLunParams creates a new CreateIscsiLunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIscsiLunParams() *CreateIscsiLunParams {
	return &CreateIscsiLunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIscsiLunParamsWithTimeout creates a new CreateIscsiLunParams object
// with the ability to set a timeout on a request.
func NewCreateIscsiLunParamsWithTimeout(timeout time.Duration) *CreateIscsiLunParams {
	return &CreateIscsiLunParams{
		timeout: timeout,
	}
}

// NewCreateIscsiLunParamsWithContext creates a new CreateIscsiLunParams object
// with the ability to set a context for a request.
func NewCreateIscsiLunParamsWithContext(ctx context.Context) *CreateIscsiLunParams {
	return &CreateIscsiLunParams{
		Context: ctx,
	}
}

// NewCreateIscsiLunParamsWithHTTPClient creates a new CreateIscsiLunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIscsiLunParamsWithHTTPClient(client *http.Client) *CreateIscsiLunParams {
	return &CreateIscsiLunParams{
		HTTPClient: client,
	}
}

/* CreateIscsiLunParams contains all the parameters to send to the API endpoint
   for the create iscsi lun operation.

   Typically these are written to a http.Request.
*/
type CreateIscsiLunParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.IscsiLunCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIscsiLunParams) WithDefaults() *CreateIscsiLunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIscsiLunParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateIscsiLunParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create iscsi lun params
func (o *CreateIscsiLunParams) WithTimeout(timeout time.Duration) *CreateIscsiLunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create iscsi lun params
func (o *CreateIscsiLunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create iscsi lun params
func (o *CreateIscsiLunParams) WithContext(ctx context.Context) *CreateIscsiLunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create iscsi lun params
func (o *CreateIscsiLunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create iscsi lun params
func (o *CreateIscsiLunParams) WithHTTPClient(client *http.Client) *CreateIscsiLunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create iscsi lun params
func (o *CreateIscsiLunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create iscsi lun params
func (o *CreateIscsiLunParams) WithContentLanguage(contentLanguage *string) *CreateIscsiLunParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create iscsi lun params
func (o *CreateIscsiLunParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create iscsi lun params
func (o *CreateIscsiLunParams) WithRequestBody(requestBody []*models.IscsiLunCreationParams) *CreateIscsiLunParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create iscsi lun params
func (o *CreateIscsiLunParams) SetRequestBody(requestBody []*models.IscsiLunCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIscsiLunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
