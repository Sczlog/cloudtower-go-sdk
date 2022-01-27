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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewCreateVMFolderParams creates a new CreateVMFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMFolderParams() *CreateVMFolderParams {
	return &CreateVMFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMFolderParamsWithTimeout creates a new CreateVMFolderParams object
// with the ability to set a timeout on a request.
func NewCreateVMFolderParamsWithTimeout(timeout time.Duration) *CreateVMFolderParams {
	return &CreateVMFolderParams{
		timeout: timeout,
	}
}

// NewCreateVMFolderParamsWithContext creates a new CreateVMFolderParams object
// with the ability to set a context for a request.
func NewCreateVMFolderParamsWithContext(ctx context.Context) *CreateVMFolderParams {
	return &CreateVMFolderParams{
		Context: ctx,
	}
}

// NewCreateVMFolderParamsWithHTTPClient creates a new CreateVMFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMFolderParamsWithHTTPClient(client *http.Client) *CreateVMFolderParams {
	return &CreateVMFolderParams{
		HTTPClient: client,
	}
}

/* CreateVMFolderParams contains all the parameters to send to the API endpoint
   for the create Vm folder operation.

   Typically these are written to a http.Request.
*/
type CreateVMFolderParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VMFolderCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMFolderParams) WithDefaults() *CreateVMFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMFolderParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVMFolderParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create Vm folder params
func (o *CreateVMFolderParams) WithTimeout(timeout time.Duration) *CreateVMFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Vm folder params
func (o *CreateVMFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Vm folder params
func (o *CreateVMFolderParams) WithContext(ctx context.Context) *CreateVMFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Vm folder params
func (o *CreateVMFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Vm folder params
func (o *CreateVMFolderParams) WithHTTPClient(client *http.Client) *CreateVMFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Vm folder params
func (o *CreateVMFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create Vm folder params
func (o *CreateVMFolderParams) WithContentLanguage(contentLanguage *string) *CreateVMFolderParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create Vm folder params
func (o *CreateVMFolderParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create Vm folder params
func (o *CreateVMFolderParams) WithRequestBody(requestBody []*models.VMFolderCreationParams) *CreateVMFolderParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create Vm folder params
func (o *CreateVMFolderParams) SetRequestBody(requestBody []*models.VMFolderCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
