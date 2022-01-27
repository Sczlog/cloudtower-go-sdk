// Code generated by go-swagger; DO NOT EDIT.

package zone

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

// NewGetZonesConnectionParams creates a new GetZonesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZonesConnectionParams() *GetZonesConnectionParams {
	return &GetZonesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZonesConnectionParamsWithTimeout creates a new GetZonesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetZonesConnectionParamsWithTimeout(timeout time.Duration) *GetZonesConnectionParams {
	return &GetZonesConnectionParams{
		timeout: timeout,
	}
}

// NewGetZonesConnectionParamsWithContext creates a new GetZonesConnectionParams object
// with the ability to set a context for a request.
func NewGetZonesConnectionParamsWithContext(ctx context.Context) *GetZonesConnectionParams {
	return &GetZonesConnectionParams{
		Context: ctx,
	}
}

// NewGetZonesConnectionParamsWithHTTPClient creates a new GetZonesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZonesConnectionParamsWithHTTPClient(client *http.Client) *GetZonesConnectionParams {
	return &GetZonesConnectionParams{
		HTTPClient: client,
	}
}

/* GetZonesConnectionParams contains all the parameters to send to the API endpoint
   for the get zones connection operation.

   Typically these are written to a http.Request.
*/
type GetZonesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetZonesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zones connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesConnectionParams) WithDefaults() *GetZonesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zones connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetZonesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get zones connection params
func (o *GetZonesConnectionParams) WithTimeout(timeout time.Duration) *GetZonesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zones connection params
func (o *GetZonesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zones connection params
func (o *GetZonesConnectionParams) WithContext(ctx context.Context) *GetZonesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zones connection params
func (o *GetZonesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zones connection params
func (o *GetZonesConnectionParams) WithHTTPClient(client *http.Client) *GetZonesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zones connection params
func (o *GetZonesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get zones connection params
func (o *GetZonesConnectionParams) WithContentLanguage(contentLanguage *string) *GetZonesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get zones connection params
func (o *GetZonesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get zones connection params
func (o *GetZonesConnectionParams) WithRequestBody(requestBody *models.GetZonesConnectionRequestBody) *GetZonesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get zones connection params
func (o *GetZonesConnectionParams) SetRequestBody(requestBody *models.GetZonesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetZonesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
