// Code generated by go-swagger; DO NOT EDIT.

package host

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

// NewEnterMaintenanceModeParams creates a new EnterMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnterMaintenanceModeParams() *EnterMaintenanceModeParams {
	return &EnterMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnterMaintenanceModeParamsWithTimeout creates a new EnterMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewEnterMaintenanceModeParamsWithTimeout(timeout time.Duration) *EnterMaintenanceModeParams {
	return &EnterMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewEnterMaintenanceModeParamsWithContext creates a new EnterMaintenanceModeParams object
// with the ability to set a context for a request.
func NewEnterMaintenanceModeParamsWithContext(ctx context.Context) *EnterMaintenanceModeParams {
	return &EnterMaintenanceModeParams{
		Context: ctx,
	}
}

// NewEnterMaintenanceModeParamsWithHTTPClient creates a new EnterMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnterMaintenanceModeParamsWithHTTPClient(client *http.Client) *EnterMaintenanceModeParams {
	return &EnterMaintenanceModeParams{
		HTTPClient: client,
	}
}

/* EnterMaintenanceModeParams contains all the parameters to send to the API endpoint
   for the enter maintenance mode operation.

   Typically these are written to a http.Request.
*/
type EnterMaintenanceModeParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.EnterMaintenanceModeParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enter maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterMaintenanceModeParams) WithDefaults() *EnterMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enter maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterMaintenanceModeParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := EnterMaintenanceModeParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) WithTimeout(timeout time.Duration) *EnterMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) WithContext(ctx context.Context) *EnterMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) WithHTTPClient(client *http.Client) *EnterMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) WithContentLanguage(contentLanguage *string) *EnterMaintenanceModeParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) WithRequestBody(requestBody *models.EnterMaintenanceModeParams) *EnterMaintenanceModeParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the enter maintenance mode params
func (o *EnterMaintenanceModeParams) SetRequestBody(requestBody *models.EnterMaintenanceModeParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *EnterMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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