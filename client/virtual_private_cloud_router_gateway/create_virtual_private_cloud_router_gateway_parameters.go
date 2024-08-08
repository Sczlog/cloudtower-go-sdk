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

// NewCreateVirtualPrivateCloudRouterGatewayParams creates a new CreateVirtualPrivateCloudRouterGatewayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVirtualPrivateCloudRouterGatewayParams() *CreateVirtualPrivateCloudRouterGatewayParams {
	return &CreateVirtualPrivateCloudRouterGatewayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualPrivateCloudRouterGatewayParamsWithTimeout creates a new CreateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a timeout on a request.
func NewCreateVirtualPrivateCloudRouterGatewayParamsWithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudRouterGatewayParams {
	return &CreateVirtualPrivateCloudRouterGatewayParams{
		timeout: timeout,
	}
}

// NewCreateVirtualPrivateCloudRouterGatewayParamsWithContext creates a new CreateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a context for a request.
func NewCreateVirtualPrivateCloudRouterGatewayParamsWithContext(ctx context.Context) *CreateVirtualPrivateCloudRouterGatewayParams {
	return &CreateVirtualPrivateCloudRouterGatewayParams{
		Context: ctx,
	}
}

// NewCreateVirtualPrivateCloudRouterGatewayParamsWithHTTPClient creates a new CreateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVirtualPrivateCloudRouterGatewayParamsWithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudRouterGatewayParams {
	return &CreateVirtualPrivateCloudRouterGatewayParams{
		HTTPClient: client,
	}
}

/* CreateVirtualPrivateCloudRouterGatewayParams contains all the parameters to send to the API endpoint
   for the create virtual private cloud router gateway operation.

   Typically these are written to a http.Request.
*/
type CreateVirtualPrivateCloudRouterGatewayParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VirtualPrivateCloudRouterGatewayCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create virtual private cloud router gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithDefaults() *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create virtual private cloud router gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVirtualPrivateCloudRouterGatewayParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithContext(ctx context.Context) *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithContentLanguage(contentLanguage *string) *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WithRequestBody(requestBody []*models.VirtualPrivateCloudRouterGatewayCreationParams) *CreateVirtualPrivateCloudRouterGatewayParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create virtual private cloud router gateway params
func (o *CreateVirtualPrivateCloudRouterGatewayParams) SetRequestBody(requestBody []*models.VirtualPrivateCloudRouterGatewayCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualPrivateCloudRouterGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
