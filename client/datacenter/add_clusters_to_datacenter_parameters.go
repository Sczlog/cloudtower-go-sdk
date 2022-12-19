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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewAddClustersToDatacenterParams creates a new AddClustersToDatacenterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddClustersToDatacenterParams() *AddClustersToDatacenterParams {
	return &AddClustersToDatacenterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddClustersToDatacenterParamsWithTimeout creates a new AddClustersToDatacenterParams object
// with the ability to set a timeout on a request.
func NewAddClustersToDatacenterParamsWithTimeout(timeout time.Duration) *AddClustersToDatacenterParams {
	return &AddClustersToDatacenterParams{
		timeout: timeout,
	}
}

// NewAddClustersToDatacenterParamsWithContext creates a new AddClustersToDatacenterParams object
// with the ability to set a context for a request.
func NewAddClustersToDatacenterParamsWithContext(ctx context.Context) *AddClustersToDatacenterParams {
	return &AddClustersToDatacenterParams{
		Context: ctx,
	}
}

// NewAddClustersToDatacenterParamsWithHTTPClient creates a new AddClustersToDatacenterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddClustersToDatacenterParamsWithHTTPClient(client *http.Client) *AddClustersToDatacenterParams {
	return &AddClustersToDatacenterParams{
		HTTPClient: client,
	}
}

/* AddClustersToDatacenterParams contains all the parameters to send to the API endpoint
   for the add clusters to datacenter operation.

   Typically these are written to a http.Request.
*/
type AddClustersToDatacenterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.AddClustersToDatacenterParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add clusters to datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClustersToDatacenterParams) WithDefaults() *AddClustersToDatacenterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add clusters to datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClustersToDatacenterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := AddClustersToDatacenterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) WithTimeout(timeout time.Duration) *AddClustersToDatacenterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) WithContext(ctx context.Context) *AddClustersToDatacenterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) WithHTTPClient(client *http.Client) *AddClustersToDatacenterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) WithContentLanguage(contentLanguage *string) *AddClustersToDatacenterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) WithRequestBody(requestBody []*models.AddClustersToDatacenterParams) *AddClustersToDatacenterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the add clusters to datacenter params
func (o *AddClustersToDatacenterParams) SetRequestBody(requestBody []*models.AddClustersToDatacenterParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *AddClustersToDatacenterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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