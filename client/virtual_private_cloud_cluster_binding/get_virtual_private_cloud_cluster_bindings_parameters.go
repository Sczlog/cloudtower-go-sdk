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

// NewGetVirtualPrivateCloudClusterBindingsParams creates a new GetVirtualPrivateCloudClusterBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudClusterBindingsParams() *GetVirtualPrivateCloudClusterBindingsParams {
	return &GetVirtualPrivateCloudClusterBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsParamsWithTimeout creates a new GetVirtualPrivateCloudClusterBindingsParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudClusterBindingsParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudClusterBindingsParams {
	return &GetVirtualPrivateCloudClusterBindingsParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsParamsWithContext creates a new GetVirtualPrivateCloudClusterBindingsParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudClusterBindingsParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudClusterBindingsParams {
	return &GetVirtualPrivateCloudClusterBindingsParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudClusterBindingsParamsWithHTTPClient creates a new GetVirtualPrivateCloudClusterBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudClusterBindingsParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudClusterBindingsParams {
	return &GetVirtualPrivateCloudClusterBindingsParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudClusterBindingsParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud cluster bindings operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudClusterBindingsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudClusterBindingsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud cluster bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithDefaults() *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud cluster bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudClusterBindingsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudClusterBindingsRequestBody) *GetVirtualPrivateCloudClusterBindingsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud cluster bindings params
func (o *GetVirtualPrivateCloudClusterBindingsParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudClusterBindingsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudClusterBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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