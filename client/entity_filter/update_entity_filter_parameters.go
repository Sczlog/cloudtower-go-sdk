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

// NewUpdateEntityFilterParams creates a new UpdateEntityFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEntityFilterParams() *UpdateEntityFilterParams {
	return &UpdateEntityFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEntityFilterParamsWithTimeout creates a new UpdateEntityFilterParams object
// with the ability to set a timeout on a request.
func NewUpdateEntityFilterParamsWithTimeout(timeout time.Duration) *UpdateEntityFilterParams {
	return &UpdateEntityFilterParams{
		timeout: timeout,
	}
}

// NewUpdateEntityFilterParamsWithContext creates a new UpdateEntityFilterParams object
// with the ability to set a context for a request.
func NewUpdateEntityFilterParamsWithContext(ctx context.Context) *UpdateEntityFilterParams {
	return &UpdateEntityFilterParams{
		Context: ctx,
	}
}

// NewUpdateEntityFilterParamsWithHTTPClient creates a new UpdateEntityFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEntityFilterParamsWithHTTPClient(client *http.Client) *UpdateEntityFilterParams {
	return &UpdateEntityFilterParams{
		HTTPClient: client,
	}
}

/* UpdateEntityFilterParams contains all the parameters to send to the API endpoint
   for the update entity filter operation.

   Typically these are written to a http.Request.
*/
type UpdateEntityFilterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.EntityFilterUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEntityFilterParams) WithDefaults() *UpdateEntityFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEntityFilterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateEntityFilterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update entity filter params
func (o *UpdateEntityFilterParams) WithTimeout(timeout time.Duration) *UpdateEntityFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update entity filter params
func (o *UpdateEntityFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update entity filter params
func (o *UpdateEntityFilterParams) WithContext(ctx context.Context) *UpdateEntityFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update entity filter params
func (o *UpdateEntityFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update entity filter params
func (o *UpdateEntityFilterParams) WithHTTPClient(client *http.Client) *UpdateEntityFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update entity filter params
func (o *UpdateEntityFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update entity filter params
func (o *UpdateEntityFilterParams) WithContentLanguage(contentLanguage *string) *UpdateEntityFilterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update entity filter params
func (o *UpdateEntityFilterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update entity filter params
func (o *UpdateEntityFilterParams) WithRequestBody(requestBody *models.EntityFilterUpdationParams) *UpdateEntityFilterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update entity filter params
func (o *UpdateEntityFilterParams) SetRequestBody(requestBody *models.EntityFilterUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEntityFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
