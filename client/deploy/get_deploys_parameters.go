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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewGetDeploysParams creates a new GetDeploysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploysParams() *GetDeploysParams {
	return &GetDeploysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploysParamsWithTimeout creates a new GetDeploysParams object
// with the ability to set a timeout on a request.
func NewGetDeploysParamsWithTimeout(timeout time.Duration) *GetDeploysParams {
	return &GetDeploysParams{
		timeout: timeout,
	}
}

// NewGetDeploysParamsWithContext creates a new GetDeploysParams object
// with the ability to set a context for a request.
func NewGetDeploysParamsWithContext(ctx context.Context) *GetDeploysParams {
	return &GetDeploysParams{
		Context: ctx,
	}
}

// NewGetDeploysParamsWithHTTPClient creates a new GetDeploysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploysParamsWithHTTPClient(client *http.Client) *GetDeploysParams {
	return &GetDeploysParams{
		HTTPClient: client,
	}
}

/* GetDeploysParams contains all the parameters to send to the API endpoint
   for the get deploys operation.

   Typically these are written to a http.Request.
*/
type GetDeploysParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetDeploysRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploysParams) WithDefaults() *GetDeploysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploysParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetDeploysParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deploys params
func (o *GetDeploysParams) WithTimeout(timeout time.Duration) *GetDeploysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deploys params
func (o *GetDeploysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deploys params
func (o *GetDeploysParams) WithContext(ctx context.Context) *GetDeploysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deploys params
func (o *GetDeploysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deploys params
func (o *GetDeploysParams) WithHTTPClient(client *http.Client) *GetDeploysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deploys params
func (o *GetDeploysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get deploys params
func (o *GetDeploysParams) WithContentLanguage(contentLanguage *string) *GetDeploysParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get deploys params
func (o *GetDeploysParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get deploys params
func (o *GetDeploysParams) WithRequestBody(requestBody *models.GetDeploysRequestBody) *GetDeploysParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get deploys params
func (o *GetDeploysParams) SetRequestBody(requestBody *models.GetDeploysRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
