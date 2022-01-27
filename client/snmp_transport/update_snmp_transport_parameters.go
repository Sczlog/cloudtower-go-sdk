// Code generated by go-swagger; DO NOT EDIT.

package snmp_transport

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

// NewUpdateSnmpTransportParams creates a new UpdateSnmpTransportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSnmpTransportParams() *UpdateSnmpTransportParams {
	return &UpdateSnmpTransportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSnmpTransportParamsWithTimeout creates a new UpdateSnmpTransportParams object
// with the ability to set a timeout on a request.
func NewUpdateSnmpTransportParamsWithTimeout(timeout time.Duration) *UpdateSnmpTransportParams {
	return &UpdateSnmpTransportParams{
		timeout: timeout,
	}
}

// NewUpdateSnmpTransportParamsWithContext creates a new UpdateSnmpTransportParams object
// with the ability to set a context for a request.
func NewUpdateSnmpTransportParamsWithContext(ctx context.Context) *UpdateSnmpTransportParams {
	return &UpdateSnmpTransportParams{
		Context: ctx,
	}
}

// NewUpdateSnmpTransportParamsWithHTTPClient creates a new UpdateSnmpTransportParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSnmpTransportParamsWithHTTPClient(client *http.Client) *UpdateSnmpTransportParams {
	return &UpdateSnmpTransportParams{
		HTTPClient: client,
	}
}

/* UpdateSnmpTransportParams contains all the parameters to send to the API endpoint
   for the update snmp transport operation.

   Typically these are written to a http.Request.
*/
type UpdateSnmpTransportParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnmpTransportUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update snmp transport params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSnmpTransportParams) WithDefaults() *UpdateSnmpTransportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update snmp transport params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSnmpTransportParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateSnmpTransportParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update snmp transport params
func (o *UpdateSnmpTransportParams) WithTimeout(timeout time.Duration) *UpdateSnmpTransportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update snmp transport params
func (o *UpdateSnmpTransportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update snmp transport params
func (o *UpdateSnmpTransportParams) WithContext(ctx context.Context) *UpdateSnmpTransportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update snmp transport params
func (o *UpdateSnmpTransportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update snmp transport params
func (o *UpdateSnmpTransportParams) WithHTTPClient(client *http.Client) *UpdateSnmpTransportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update snmp transport params
func (o *UpdateSnmpTransportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update snmp transport params
func (o *UpdateSnmpTransportParams) WithContentLanguage(contentLanguage *string) *UpdateSnmpTransportParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update snmp transport params
func (o *UpdateSnmpTransportParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update snmp transport params
func (o *UpdateSnmpTransportParams) WithRequestBody(requestBody *models.SnmpTransportUpdationParams) *UpdateSnmpTransportParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update snmp transport params
func (o *UpdateSnmpTransportParams) SetRequestBody(requestBody *models.SnmpTransportUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSnmpTransportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
