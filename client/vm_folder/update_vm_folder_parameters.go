// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

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

// NewUpdateVMFolderParams creates a new UpdateVMFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMFolderParams() *UpdateVMFolderParams {
	return &UpdateVMFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMFolderParamsWithTimeout creates a new UpdateVMFolderParams object
// with the ability to set a timeout on a request.
func NewUpdateVMFolderParamsWithTimeout(timeout time.Duration) *UpdateVMFolderParams {
	return &UpdateVMFolderParams{
		timeout: timeout,
	}
}

// NewUpdateVMFolderParamsWithContext creates a new UpdateVMFolderParams object
// with the ability to set a context for a request.
func NewUpdateVMFolderParamsWithContext(ctx context.Context) *UpdateVMFolderParams {
	return &UpdateVMFolderParams{
		Context: ctx,
	}
}

// NewUpdateVMFolderParamsWithHTTPClient creates a new UpdateVMFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMFolderParamsWithHTTPClient(client *http.Client) *UpdateVMFolderParams {
	return &UpdateVMFolderParams{
		HTTPClient: client,
	}
}

/* UpdateVMFolderParams contains all the parameters to send to the API endpoint
   for the update Vm folder operation.

   Typically these are written to a http.Request.
*/
type UpdateVMFolderParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMFolderUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMFolderParams) WithDefaults() *UpdateVMFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMFolderParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMFolderParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm folder params
func (o *UpdateVMFolderParams) WithTimeout(timeout time.Duration) *UpdateVMFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm folder params
func (o *UpdateVMFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm folder params
func (o *UpdateVMFolderParams) WithContext(ctx context.Context) *UpdateVMFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm folder params
func (o *UpdateVMFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm folder params
func (o *UpdateVMFolderParams) WithHTTPClient(client *http.Client) *UpdateVMFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm folder params
func (o *UpdateVMFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm folder params
func (o *UpdateVMFolderParams) WithContentLanguage(contentLanguage *string) *UpdateVMFolderParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm folder params
func (o *UpdateVMFolderParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm folder params
func (o *UpdateVMFolderParams) WithRequestBody(requestBody *models.VMFolderUpdationParams) *UpdateVMFolderParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm folder params
func (o *UpdateVMFolderParams) SetRequestBody(requestBody *models.VMFolderUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
