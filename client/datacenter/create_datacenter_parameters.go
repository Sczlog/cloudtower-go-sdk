// Code generated by go-swagger; DO NOT EDIT.

package datacenter

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

// NewCreateDatacenterParams creates a new CreateDatacenterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDatacenterParams() *CreateDatacenterParams {
	return &CreateDatacenterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatacenterParamsWithTimeout creates a new CreateDatacenterParams object
// with the ability to set a timeout on a request.
func NewCreateDatacenterParamsWithTimeout(timeout time.Duration) *CreateDatacenterParams {
	return &CreateDatacenterParams{
		timeout: timeout,
	}
}

// NewCreateDatacenterParamsWithContext creates a new CreateDatacenterParams object
// with the ability to set a context for a request.
func NewCreateDatacenterParamsWithContext(ctx context.Context) *CreateDatacenterParams {
	return &CreateDatacenterParams{
		Context: ctx,
	}
}

// NewCreateDatacenterParamsWithHTTPClient creates a new CreateDatacenterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDatacenterParamsWithHTTPClient(client *http.Client) *CreateDatacenterParams {
	return &CreateDatacenterParams{
		HTTPClient: client,
	}
}

/* CreateDatacenterParams contains all the parameters to send to the API endpoint
   for the create datacenter operation.

   Typically these are written to a http.Request.
*/
type CreateDatacenterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.DatacenterCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatacenterParams) WithDefaults() *CreateDatacenterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatacenterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateDatacenterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create datacenter params
func (o *CreateDatacenterParams) WithTimeout(timeout time.Duration) *CreateDatacenterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create datacenter params
func (o *CreateDatacenterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create datacenter params
func (o *CreateDatacenterParams) WithContext(ctx context.Context) *CreateDatacenterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create datacenter params
func (o *CreateDatacenterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create datacenter params
func (o *CreateDatacenterParams) WithHTTPClient(client *http.Client) *CreateDatacenterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create datacenter params
func (o *CreateDatacenterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create datacenter params
func (o *CreateDatacenterParams) WithContentLanguage(contentLanguage *string) *CreateDatacenterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create datacenter params
func (o *CreateDatacenterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create datacenter params
func (o *CreateDatacenterParams) WithRequestBody(requestBody []*models.DatacenterCreationParams) *CreateDatacenterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create datacenter params
func (o *CreateDatacenterParams) SetRequestBody(requestBody []*models.DatacenterCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatacenterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
