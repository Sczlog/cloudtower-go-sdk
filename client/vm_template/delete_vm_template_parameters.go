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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewDeleteVMTemplateParams creates a new DeleteVMTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMTemplateParams() *DeleteVMTemplateParams {
	return &DeleteVMTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMTemplateParamsWithTimeout creates a new DeleteVMTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteVMTemplateParamsWithTimeout(timeout time.Duration) *DeleteVMTemplateParams {
	return &DeleteVMTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteVMTemplateParamsWithContext creates a new DeleteVMTemplateParams object
// with the ability to set a context for a request.
func NewDeleteVMTemplateParamsWithContext(ctx context.Context) *DeleteVMTemplateParams {
	return &DeleteVMTemplateParams{
		Context: ctx,
	}
}

// NewDeleteVMTemplateParamsWithHTTPClient creates a new DeleteVMTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMTemplateParamsWithHTTPClient(client *http.Client) *DeleteVMTemplateParams {
	return &DeleteVMTemplateParams{
		HTTPClient: client,
	}
}

/* DeleteVMTemplateParams contains all the parameters to send to the API endpoint
   for the delete Vm template operation.

   Typically these are written to a http.Request.
*/
type DeleteVMTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMTemplateDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMTemplateParams) WithDefaults() *DeleteVMTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Vm template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVMTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete Vm template params
func (o *DeleteVMTemplateParams) WithTimeout(timeout time.Duration) *DeleteVMTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Vm template params
func (o *DeleteVMTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Vm template params
func (o *DeleteVMTemplateParams) WithContext(ctx context.Context) *DeleteVMTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Vm template params
func (o *DeleteVMTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Vm template params
func (o *DeleteVMTemplateParams) WithHTTPClient(client *http.Client) *DeleteVMTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Vm template params
func (o *DeleteVMTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete Vm template params
func (o *DeleteVMTemplateParams) WithContentLanguage(contentLanguage *string) *DeleteVMTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete Vm template params
func (o *DeleteVMTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete Vm template params
func (o *DeleteVMTemplateParams) WithRequestBody(requestBody *models.VMTemplateDeletionParams) *DeleteVMTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete Vm template params
func (o *DeleteVMTemplateParams) SetRequestBody(requestBody *models.VMTemplateDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
