// Code generated by go-swagger; DO NOT EDIT.

package label

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

// NewGetLabelsConnectionParams creates a new GetLabelsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelsConnectionParams() *GetLabelsConnectionParams {
	return &GetLabelsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsConnectionParamsWithTimeout creates a new GetLabelsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetLabelsConnectionParamsWithTimeout(timeout time.Duration) *GetLabelsConnectionParams {
	return &GetLabelsConnectionParams{
		timeout: timeout,
	}
}

// NewGetLabelsConnectionParamsWithContext creates a new GetLabelsConnectionParams object
// with the ability to set a context for a request.
func NewGetLabelsConnectionParamsWithContext(ctx context.Context) *GetLabelsConnectionParams {
	return &GetLabelsConnectionParams{
		Context: ctx,
	}
}

// NewGetLabelsConnectionParamsWithHTTPClient creates a new GetLabelsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelsConnectionParamsWithHTTPClient(client *http.Client) *GetLabelsConnectionParams {
	return &GetLabelsConnectionParams{
		HTTPClient: client,
	}
}

/* GetLabelsConnectionParams contains all the parameters to send to the API endpoint
   for the get labels connection operation.

   Typically these are written to a http.Request.
*/
type GetLabelsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetLabelsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get labels connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsConnectionParams) WithDefaults() *GetLabelsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get labels connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetLabelsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get labels connection params
func (o *GetLabelsConnectionParams) WithTimeout(timeout time.Duration) *GetLabelsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels connection params
func (o *GetLabelsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels connection params
func (o *GetLabelsConnectionParams) WithContext(ctx context.Context) *GetLabelsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels connection params
func (o *GetLabelsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels connection params
func (o *GetLabelsConnectionParams) WithHTTPClient(client *http.Client) *GetLabelsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels connection params
func (o *GetLabelsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get labels connection params
func (o *GetLabelsConnectionParams) WithContentLanguage(contentLanguage *string) *GetLabelsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get labels connection params
func (o *GetLabelsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get labels connection params
func (o *GetLabelsConnectionParams) WithRequestBody(requestBody *models.GetLabelsConnectionRequestBody) *GetLabelsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get labels connection params
func (o *GetLabelsConnectionParams) SetRequestBody(requestBody *models.GetLabelsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
