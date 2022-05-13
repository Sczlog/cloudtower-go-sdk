// Code generated by go-swagger; DO NOT EDIT.

package deploy

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

// NewGetDeploysConnectionParams creates a new GetDeploysConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploysConnectionParams() *GetDeploysConnectionParams {
	return &GetDeploysConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploysConnectionParamsWithTimeout creates a new GetDeploysConnectionParams object
// with the ability to set a timeout on a request.
func NewGetDeploysConnectionParamsWithTimeout(timeout time.Duration) *GetDeploysConnectionParams {
	return &GetDeploysConnectionParams{
		timeout: timeout,
	}
}

// NewGetDeploysConnectionParamsWithContext creates a new GetDeploysConnectionParams object
// with the ability to set a context for a request.
func NewGetDeploysConnectionParamsWithContext(ctx context.Context) *GetDeploysConnectionParams {
	return &GetDeploysConnectionParams{
		Context: ctx,
	}
}

// NewGetDeploysConnectionParamsWithHTTPClient creates a new GetDeploysConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploysConnectionParamsWithHTTPClient(client *http.Client) *GetDeploysConnectionParams {
	return &GetDeploysConnectionParams{
		HTTPClient: client,
	}
}

/* GetDeploysConnectionParams contains all the parameters to send to the API endpoint
   for the get deploys connection operation.

   Typically these are written to a http.Request.
*/
type GetDeploysConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetDeploysConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deploys connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploysConnectionParams) WithDefaults() *GetDeploysConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deploys connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploysConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetDeploysConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deploys connection params
func (o *GetDeploysConnectionParams) WithTimeout(timeout time.Duration) *GetDeploysConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deploys connection params
func (o *GetDeploysConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deploys connection params
func (o *GetDeploysConnectionParams) WithContext(ctx context.Context) *GetDeploysConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deploys connection params
func (o *GetDeploysConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deploys connection params
func (o *GetDeploysConnectionParams) WithHTTPClient(client *http.Client) *GetDeploysConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deploys connection params
func (o *GetDeploysConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get deploys connection params
func (o *GetDeploysConnectionParams) WithContentLanguage(contentLanguage *string) *GetDeploysConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get deploys connection params
func (o *GetDeploysConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get deploys connection params
func (o *GetDeploysConnectionParams) WithRequestBody(requestBody *models.GetDeploysConnectionRequestBody) *GetDeploysConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get deploys connection params
func (o *GetDeploysConnectionParams) SetRequestBody(requestBody *models.GetDeploysConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploysConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
