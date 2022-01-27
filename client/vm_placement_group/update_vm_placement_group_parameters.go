// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

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

// NewUpdateVMPlacementGroupParams creates a new UpdateVMPlacementGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMPlacementGroupParams() *UpdateVMPlacementGroupParams {
	return &UpdateVMPlacementGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMPlacementGroupParamsWithTimeout creates a new UpdateVMPlacementGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateVMPlacementGroupParamsWithTimeout(timeout time.Duration) *UpdateVMPlacementGroupParams {
	return &UpdateVMPlacementGroupParams{
		timeout: timeout,
	}
}

// NewUpdateVMPlacementGroupParamsWithContext creates a new UpdateVMPlacementGroupParams object
// with the ability to set a context for a request.
func NewUpdateVMPlacementGroupParamsWithContext(ctx context.Context) *UpdateVMPlacementGroupParams {
	return &UpdateVMPlacementGroupParams{
		Context: ctx,
	}
}

// NewUpdateVMPlacementGroupParamsWithHTTPClient creates a new UpdateVMPlacementGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMPlacementGroupParamsWithHTTPClient(client *http.Client) *UpdateVMPlacementGroupParams {
	return &UpdateVMPlacementGroupParams{
		HTTPClient: client,
	}
}

/* UpdateVMPlacementGroupParams contains all the parameters to send to the API endpoint
   for the update Vm placement group operation.

   Typically these are written to a http.Request.
*/
type UpdateVMPlacementGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMPlacementGroupUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMPlacementGroupParams) WithDefaults() *UpdateVMPlacementGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMPlacementGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMPlacementGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) WithTimeout(timeout time.Duration) *UpdateVMPlacementGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) WithContext(ctx context.Context) *UpdateVMPlacementGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) WithHTTPClient(client *http.Client) *UpdateVMPlacementGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) WithContentLanguage(contentLanguage *string) *UpdateVMPlacementGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) WithRequestBody(requestBody *models.VMPlacementGroupUpdationParams) *UpdateVMPlacementGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm placement group params
func (o *UpdateVMPlacementGroupParams) SetRequestBody(requestBody *models.VMPlacementGroupUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMPlacementGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
