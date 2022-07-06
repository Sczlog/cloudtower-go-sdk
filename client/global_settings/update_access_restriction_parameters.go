// Code generated by go-swagger; DO NOT EDIT.

package global_settings

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

// NewUpdateAccessRestrictionParams creates a new UpdateAccessRestrictionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccessRestrictionParams() *UpdateAccessRestrictionParams {
	return &UpdateAccessRestrictionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccessRestrictionParamsWithTimeout creates a new UpdateAccessRestrictionParams object
// with the ability to set a timeout on a request.
func NewUpdateAccessRestrictionParamsWithTimeout(timeout time.Duration) *UpdateAccessRestrictionParams {
	return &UpdateAccessRestrictionParams{
		timeout: timeout,
	}
}

// NewUpdateAccessRestrictionParamsWithContext creates a new UpdateAccessRestrictionParams object
// with the ability to set a context for a request.
func NewUpdateAccessRestrictionParamsWithContext(ctx context.Context) *UpdateAccessRestrictionParams {
	return &UpdateAccessRestrictionParams{
		Context: ctx,
	}
}

// NewUpdateAccessRestrictionParamsWithHTTPClient creates a new UpdateAccessRestrictionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccessRestrictionParamsWithHTTPClient(client *http.Client) *UpdateAccessRestrictionParams {
	return &UpdateAccessRestrictionParams{
		HTTPClient: client,
	}
}

/* UpdateAccessRestrictionParams contains all the parameters to send to the API endpoint
   for the update access restriction operation.

   Typically these are written to a http.Request.
*/
type UpdateAccessRestrictionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.UpdateAccessRestrictionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update access restriction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccessRestrictionParams) WithDefaults() *UpdateAccessRestrictionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update access restriction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccessRestrictionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateAccessRestrictionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update access restriction params
func (o *UpdateAccessRestrictionParams) WithTimeout(timeout time.Duration) *UpdateAccessRestrictionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update access restriction params
func (o *UpdateAccessRestrictionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update access restriction params
func (o *UpdateAccessRestrictionParams) WithContext(ctx context.Context) *UpdateAccessRestrictionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update access restriction params
func (o *UpdateAccessRestrictionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update access restriction params
func (o *UpdateAccessRestrictionParams) WithHTTPClient(client *http.Client) *UpdateAccessRestrictionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update access restriction params
func (o *UpdateAccessRestrictionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update access restriction params
func (o *UpdateAccessRestrictionParams) WithContentLanguage(contentLanguage *string) *UpdateAccessRestrictionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update access restriction params
func (o *UpdateAccessRestrictionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update access restriction params
func (o *UpdateAccessRestrictionParams) WithRequestBody(requestBody *models.UpdateAccessRestrictionParams) *UpdateAccessRestrictionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update access restriction params
func (o *UpdateAccessRestrictionParams) SetRequestBody(requestBody *models.UpdateAccessRestrictionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccessRestrictionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
