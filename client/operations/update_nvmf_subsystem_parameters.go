// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewUpdateNvmfSubsystemParams creates a new UpdateNvmfSubsystemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNvmfSubsystemParams() *UpdateNvmfSubsystemParams {
	return &UpdateNvmfSubsystemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNvmfSubsystemParamsWithTimeout creates a new UpdateNvmfSubsystemParams object
// with the ability to set a timeout on a request.
func NewUpdateNvmfSubsystemParamsWithTimeout(timeout time.Duration) *UpdateNvmfSubsystemParams {
	return &UpdateNvmfSubsystemParams{
		timeout: timeout,
	}
}

// NewUpdateNvmfSubsystemParamsWithContext creates a new UpdateNvmfSubsystemParams object
// with the ability to set a context for a request.
func NewUpdateNvmfSubsystemParamsWithContext(ctx context.Context) *UpdateNvmfSubsystemParams {
	return &UpdateNvmfSubsystemParams{
		Context: ctx,
	}
}

// NewUpdateNvmfSubsystemParamsWithHTTPClient creates a new UpdateNvmfSubsystemParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNvmfSubsystemParamsWithHTTPClient(client *http.Client) *UpdateNvmfSubsystemParams {
	return &UpdateNvmfSubsystemParams{
		HTTPClient: client,
	}
}

/* UpdateNvmfSubsystemParams contains all the parameters to send to the API endpoint
   for the update nvmf subsystem operation.

   Typically these are written to a http.Request.
*/
type UpdateNvmfSubsystemParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.NvmfSubsystemUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update nvmf subsystem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNvmfSubsystemParams) WithDefaults() *UpdateNvmfSubsystemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update nvmf subsystem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNvmfSubsystemParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateNvmfSubsystemParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) WithTimeout(timeout time.Duration) *UpdateNvmfSubsystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) WithContext(ctx context.Context) *UpdateNvmfSubsystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) WithHTTPClient(client *http.Client) *UpdateNvmfSubsystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) WithContentLanguage(contentLanguage *string) *UpdateNvmfSubsystemParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) WithRequestBody(requestBody *models.NvmfSubsystemUpdationParams) *UpdateNvmfSubsystemParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update nvmf subsystem params
func (o *UpdateNvmfSubsystemParams) SetRequestBody(requestBody *models.NvmfSubsystemUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNvmfSubsystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
