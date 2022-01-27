// Code generated by go-swagger; DO NOT EDIT.

package vm_template

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

// NewCloneVMTemplateFromVMParams creates a new CloneVMTemplateFromVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneVMTemplateFromVMParams() *CloneVMTemplateFromVMParams {
	return &CloneVMTemplateFromVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneVMTemplateFromVMParamsWithTimeout creates a new CloneVMTemplateFromVMParams object
// with the ability to set a timeout on a request.
func NewCloneVMTemplateFromVMParamsWithTimeout(timeout time.Duration) *CloneVMTemplateFromVMParams {
	return &CloneVMTemplateFromVMParams{
		timeout: timeout,
	}
}

// NewCloneVMTemplateFromVMParamsWithContext creates a new CloneVMTemplateFromVMParams object
// with the ability to set a context for a request.
func NewCloneVMTemplateFromVMParamsWithContext(ctx context.Context) *CloneVMTemplateFromVMParams {
	return &CloneVMTemplateFromVMParams{
		Context: ctx,
	}
}

// NewCloneVMTemplateFromVMParamsWithHTTPClient creates a new CloneVMTemplateFromVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneVMTemplateFromVMParamsWithHTTPClient(client *http.Client) *CloneVMTemplateFromVMParams {
	return &CloneVMTemplateFromVMParams{
		HTTPClient: client,
	}
}

/* CloneVMTemplateFromVMParams contains all the parameters to send to the API endpoint
   for the clone Vm template from Vm operation.

   Typically these are written to a http.Request.
*/
type CloneVMTemplateFromVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VMTemplateCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMTemplateFromVMParams) WithDefaults() *CloneVMTemplateFromVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMTemplateFromVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CloneVMTemplateFromVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) WithTimeout(timeout time.Duration) *CloneVMTemplateFromVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) WithContext(ctx context.Context) *CloneVMTemplateFromVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) WithHTTPClient(client *http.Client) *CloneVMTemplateFromVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) WithContentLanguage(contentLanguage *string) *CloneVMTemplateFromVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) WithRequestBody(requestBody []*models.VMTemplateCreationParams) *CloneVMTemplateFromVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the clone Vm template from Vm params
func (o *CloneVMTemplateFromVMParams) SetRequestBody(requestBody []*models.VMTemplateCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CloneVMTemplateFromVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
