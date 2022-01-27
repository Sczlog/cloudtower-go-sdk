// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewAddVMToFolderParams creates a new AddVMToFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVMToFolderParams() *AddVMToFolderParams {
	return &AddVMToFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVMToFolderParamsWithTimeout creates a new AddVMToFolderParams object
// with the ability to set a timeout on a request.
func NewAddVMToFolderParamsWithTimeout(timeout time.Duration) *AddVMToFolderParams {
	return &AddVMToFolderParams{
		timeout: timeout,
	}
}

// NewAddVMToFolderParamsWithContext creates a new AddVMToFolderParams object
// with the ability to set a context for a request.
func NewAddVMToFolderParamsWithContext(ctx context.Context) *AddVMToFolderParams {
	return &AddVMToFolderParams{
		Context: ctx,
	}
}

// NewAddVMToFolderParamsWithHTTPClient creates a new AddVMToFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVMToFolderParamsWithHTTPClient(client *http.Client) *AddVMToFolderParams {
	return &AddVMToFolderParams{
		HTTPClient: client,
	}
}

/* AddVMToFolderParams contains all the parameters to send to the API endpoint
   for the add Vm to folder operation.

   Typically these are written to a http.Request.
*/
type AddVMToFolderParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMAddFolderParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Vm to folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMToFolderParams) WithDefaults() *AddVMToFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Vm to folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMToFolderParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := AddVMToFolderParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add Vm to folder params
func (o *AddVMToFolderParams) WithTimeout(timeout time.Duration) *AddVMToFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Vm to folder params
func (o *AddVMToFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Vm to folder params
func (o *AddVMToFolderParams) WithContext(ctx context.Context) *AddVMToFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Vm to folder params
func (o *AddVMToFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Vm to folder params
func (o *AddVMToFolderParams) WithHTTPClient(client *http.Client) *AddVMToFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Vm to folder params
func (o *AddVMToFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the add Vm to folder params
func (o *AddVMToFolderParams) WithContentLanguage(contentLanguage *string) *AddVMToFolderParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the add Vm to folder params
func (o *AddVMToFolderParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the add Vm to folder params
func (o *AddVMToFolderParams) WithRequestBody(requestBody *models.VMAddFolderParams) *AddVMToFolderParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the add Vm to folder params
func (o *AddVMToFolderParams) SetRequestBody(requestBody *models.VMAddFolderParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *AddVMToFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
