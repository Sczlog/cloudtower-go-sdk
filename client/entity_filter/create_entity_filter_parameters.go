// Code generated by go-swagger; DO NOT EDIT.

package entity_filter

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

// NewCreateEntityFilterParams creates a new CreateEntityFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEntityFilterParams() *CreateEntityFilterParams {
	return &CreateEntityFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEntityFilterParamsWithTimeout creates a new CreateEntityFilterParams object
// with the ability to set a timeout on a request.
func NewCreateEntityFilterParamsWithTimeout(timeout time.Duration) *CreateEntityFilterParams {
	return &CreateEntityFilterParams{
		timeout: timeout,
	}
}

// NewCreateEntityFilterParamsWithContext creates a new CreateEntityFilterParams object
// with the ability to set a context for a request.
func NewCreateEntityFilterParamsWithContext(ctx context.Context) *CreateEntityFilterParams {
	return &CreateEntityFilterParams{
		Context: ctx,
	}
}

// NewCreateEntityFilterParamsWithHTTPClient creates a new CreateEntityFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEntityFilterParamsWithHTTPClient(client *http.Client) *CreateEntityFilterParams {
	return &CreateEntityFilterParams{
		HTTPClient: client,
	}
}

/* CreateEntityFilterParams contains all the parameters to send to the API endpoint
   for the create entity filter operation.

   Typically these are written to a http.Request.
*/
type CreateEntityFilterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.EntityFilterCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEntityFilterParams) WithDefaults() *CreateEntityFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEntityFilterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateEntityFilterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create entity filter params
func (o *CreateEntityFilterParams) WithTimeout(timeout time.Duration) *CreateEntityFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create entity filter params
func (o *CreateEntityFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create entity filter params
func (o *CreateEntityFilterParams) WithContext(ctx context.Context) *CreateEntityFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create entity filter params
func (o *CreateEntityFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create entity filter params
func (o *CreateEntityFilterParams) WithHTTPClient(client *http.Client) *CreateEntityFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create entity filter params
func (o *CreateEntityFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create entity filter params
func (o *CreateEntityFilterParams) WithContentLanguage(contentLanguage *string) *CreateEntityFilterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create entity filter params
func (o *CreateEntityFilterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create entity filter params
func (o *CreateEntityFilterParams) WithRequestBody(requestBody []*models.EntityFilterCreationParams) *CreateEntityFilterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create entity filter params
func (o *CreateEntityFilterParams) SetRequestBody(requestBody []*models.EntityFilterCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEntityFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
