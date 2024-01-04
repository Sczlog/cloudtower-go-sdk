// Code generated by go-swagger; DO NOT EDIT.

package gpu_device

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

// NewUpdateGpuDeviceUsageParams creates a new UpdateGpuDeviceUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGpuDeviceUsageParams() *UpdateGpuDeviceUsageParams {
	return &UpdateGpuDeviceUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGpuDeviceUsageParamsWithTimeout creates a new UpdateGpuDeviceUsageParams object
// with the ability to set a timeout on a request.
func NewUpdateGpuDeviceUsageParamsWithTimeout(timeout time.Duration) *UpdateGpuDeviceUsageParams {
	return &UpdateGpuDeviceUsageParams{
		timeout: timeout,
	}
}

// NewUpdateGpuDeviceUsageParamsWithContext creates a new UpdateGpuDeviceUsageParams object
// with the ability to set a context for a request.
func NewUpdateGpuDeviceUsageParamsWithContext(ctx context.Context) *UpdateGpuDeviceUsageParams {
	return &UpdateGpuDeviceUsageParams{
		Context: ctx,
	}
}

// NewUpdateGpuDeviceUsageParamsWithHTTPClient creates a new UpdateGpuDeviceUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGpuDeviceUsageParamsWithHTTPClient(client *http.Client) *UpdateGpuDeviceUsageParams {
	return &UpdateGpuDeviceUsageParams{
		HTTPClient: client,
	}
}

/* UpdateGpuDeviceUsageParams contains all the parameters to send to the API endpoint
   for the update gpu device usage operation.

   Typically these are written to a http.Request.
*/
type UpdateGpuDeviceUsageParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GpuDeviceUsageUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update gpu device usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGpuDeviceUsageParams) WithDefaults() *UpdateGpuDeviceUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update gpu device usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGpuDeviceUsageParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateGpuDeviceUsageParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) WithTimeout(timeout time.Duration) *UpdateGpuDeviceUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) WithContext(ctx context.Context) *UpdateGpuDeviceUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) WithHTTPClient(client *http.Client) *UpdateGpuDeviceUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) WithContentLanguage(contentLanguage *string) *UpdateGpuDeviceUsageParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) WithRequestBody(requestBody *models.GpuDeviceUsageUpdationParams) *UpdateGpuDeviceUsageParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update gpu device usage params
func (o *UpdateGpuDeviceUsageParams) SetRequestBody(requestBody *models.GpuDeviceUsageUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGpuDeviceUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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