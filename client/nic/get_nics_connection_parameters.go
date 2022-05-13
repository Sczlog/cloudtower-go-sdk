// Code generated by go-swagger; DO NOT EDIT.

package nic

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

// NewGetNicsConnectionParams creates a new GetNicsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNicsConnectionParams() *GetNicsConnectionParams {
	return &GetNicsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNicsConnectionParamsWithTimeout creates a new GetNicsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetNicsConnectionParamsWithTimeout(timeout time.Duration) *GetNicsConnectionParams {
	return &GetNicsConnectionParams{
		timeout: timeout,
	}
}

// NewGetNicsConnectionParamsWithContext creates a new GetNicsConnectionParams object
// with the ability to set a context for a request.
func NewGetNicsConnectionParamsWithContext(ctx context.Context) *GetNicsConnectionParams {
	return &GetNicsConnectionParams{
		Context: ctx,
	}
}

// NewGetNicsConnectionParamsWithHTTPClient creates a new GetNicsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNicsConnectionParamsWithHTTPClient(client *http.Client) *GetNicsConnectionParams {
	return &GetNicsConnectionParams{
		HTTPClient: client,
	}
}

/* GetNicsConnectionParams contains all the parameters to send to the API endpoint
   for the get nics connection operation.

   Typically these are written to a http.Request.
*/
type GetNicsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetNicsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nics connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNicsConnectionParams) WithDefaults() *GetNicsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nics connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNicsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetNicsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nics connection params
func (o *GetNicsConnectionParams) WithTimeout(timeout time.Duration) *GetNicsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nics connection params
func (o *GetNicsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nics connection params
func (o *GetNicsConnectionParams) WithContext(ctx context.Context) *GetNicsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nics connection params
func (o *GetNicsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nics connection params
func (o *GetNicsConnectionParams) WithHTTPClient(client *http.Client) *GetNicsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nics connection params
func (o *GetNicsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get nics connection params
func (o *GetNicsConnectionParams) WithContentLanguage(contentLanguage *string) *GetNicsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get nics connection params
func (o *GetNicsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get nics connection params
func (o *GetNicsConnectionParams) WithRequestBody(requestBody *models.GetNicsConnectionRequestBody) *GetNicsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nics connection params
func (o *GetNicsConnectionParams) SetRequestBody(requestBody *models.GetNicsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNicsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
