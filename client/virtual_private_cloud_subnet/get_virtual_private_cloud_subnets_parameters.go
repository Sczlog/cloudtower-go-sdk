// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_subnet

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

// NewGetVirtualPrivateCloudSubnetsParams creates a new GetVirtualPrivateCloudSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudSubnetsParams() *GetVirtualPrivateCloudSubnetsParams {
	return &GetVirtualPrivateCloudSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudSubnetsParamsWithTimeout creates a new GetVirtualPrivateCloudSubnetsParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudSubnetsParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudSubnetsParams {
	return &GetVirtualPrivateCloudSubnetsParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudSubnetsParamsWithContext creates a new GetVirtualPrivateCloudSubnetsParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudSubnetsParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudSubnetsParams {
	return &GetVirtualPrivateCloudSubnetsParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudSubnetsParamsWithHTTPClient creates a new GetVirtualPrivateCloudSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudSubnetsParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudSubnetsParams {
	return &GetVirtualPrivateCloudSubnetsParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudSubnetsParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud subnets operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudSubnetsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudSubnetsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudSubnetsParams) WithDefaults() *GetVirtualPrivateCloudSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudSubnetsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudSubnetsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudSubnetsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudSubnetsRequestBody) *GetVirtualPrivateCloudSubnetsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud subnets params
func (o *GetVirtualPrivateCloudSubnetsParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudSubnetsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
