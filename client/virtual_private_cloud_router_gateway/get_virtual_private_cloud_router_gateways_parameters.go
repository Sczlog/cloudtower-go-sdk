// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_router_gateway

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

// NewGetVirtualPrivateCloudRouterGatewaysParams creates a new GetVirtualPrivateCloudRouterGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudRouterGatewaysParams() *GetVirtualPrivateCloudRouterGatewaysParams {
	return &GetVirtualPrivateCloudRouterGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudRouterGatewaysParamsWithTimeout creates a new GetVirtualPrivateCloudRouterGatewaysParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudRouterGatewaysParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudRouterGatewaysParams {
	return &GetVirtualPrivateCloudRouterGatewaysParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudRouterGatewaysParamsWithContext creates a new GetVirtualPrivateCloudRouterGatewaysParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudRouterGatewaysParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudRouterGatewaysParams {
	return &GetVirtualPrivateCloudRouterGatewaysParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudRouterGatewaysParamsWithHTTPClient creates a new GetVirtualPrivateCloudRouterGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudRouterGatewaysParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudRouterGatewaysParams {
	return &GetVirtualPrivateCloudRouterGatewaysParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudRouterGatewaysParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud router gateways operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudRouterGatewaysParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudRouterGatewaysRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud router gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithDefaults() *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud router gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudRouterGatewaysParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudRouterGatewaysRequestBody) *GetVirtualPrivateCloudRouterGatewaysParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud router gateways params
func (o *GetVirtualPrivateCloudRouterGatewaysParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudRouterGatewaysRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudRouterGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
