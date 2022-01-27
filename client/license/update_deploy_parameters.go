// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewUpdateDeployParams creates a new UpdateDeployParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeployParams() *UpdateDeployParams {
	return &UpdateDeployParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeployParamsWithTimeout creates a new UpdateDeployParams object
// with the ability to set a timeout on a request.
func NewUpdateDeployParamsWithTimeout(timeout time.Duration) *UpdateDeployParams {
	return &UpdateDeployParams{
		timeout: timeout,
	}
}

// NewUpdateDeployParamsWithContext creates a new UpdateDeployParams object
// with the ability to set a context for a request.
func NewUpdateDeployParamsWithContext(ctx context.Context) *UpdateDeployParams {
	return &UpdateDeployParams{
		Context: ctx,
	}
}

// NewUpdateDeployParamsWithHTTPClient creates a new UpdateDeployParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeployParamsWithHTTPClient(client *http.Client) *UpdateDeployParams {
	return &UpdateDeployParams{
		HTTPClient: client,
	}
}

/* UpdateDeployParams contains all the parameters to send to the API endpoint
   for the update deploy operation.

   Typically these are written to a http.Request.
*/
type UpdateDeployParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.LicenseUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeployParams) WithDefaults() *UpdateDeployParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeployParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateDeployParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update deploy params
func (o *UpdateDeployParams) WithTimeout(timeout time.Duration) *UpdateDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update deploy params
func (o *UpdateDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update deploy params
func (o *UpdateDeployParams) WithContext(ctx context.Context) *UpdateDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update deploy params
func (o *UpdateDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update deploy params
func (o *UpdateDeployParams) WithHTTPClient(client *http.Client) *UpdateDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update deploy params
func (o *UpdateDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update deploy params
func (o *UpdateDeployParams) WithContentLanguage(contentLanguage *string) *UpdateDeployParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update deploy params
func (o *UpdateDeployParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update deploy params
func (o *UpdateDeployParams) WithRequestBody(requestBody *models.LicenseUpdationParams) *UpdateDeployParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update deploy params
func (o *UpdateDeployParams) SetRequestBody(requestBody *models.LicenseUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
