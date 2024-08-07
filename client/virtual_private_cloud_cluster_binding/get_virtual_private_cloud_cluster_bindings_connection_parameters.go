// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_cluster_binding

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

// NewGetVirtualPrivateCloudClusterBindingsConnectionParams creates a new GetVirtualPrivateCloudClusterBindingsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudClusterBindingsConnectionParams() *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	return &GetVirtualPrivateCloudClusterBindingsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithTimeout creates a new GetVirtualPrivateCloudClusterBindingsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	return &GetVirtualPrivateCloudClusterBindingsConnectionParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithContext creates a new GetVirtualPrivateCloudClusterBindingsConnectionParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	return &GetVirtualPrivateCloudClusterBindingsConnectionParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithHTTPClient creates a new GetVirtualPrivateCloudClusterBindingsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudClusterBindingsConnectionParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	return &GetVirtualPrivateCloudClusterBindingsConnectionParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudClusterBindingsConnectionParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud cluster bindings connection operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudClusterBindingsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudClusterBindingsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud cluster bindings connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithDefaults() *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud cluster bindings connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudClusterBindingsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudClusterBindingsConnectionRequestBody) *GetVirtualPrivateCloudClusterBindingsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud cluster bindings connection params
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudClusterBindingsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudClusterBindingsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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