// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_route_table

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

// NewCreateVirtualPrivateCloudRouteTableParams creates a new CreateVirtualPrivateCloudRouteTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVirtualPrivateCloudRouteTableParams() *CreateVirtualPrivateCloudRouteTableParams {
	return &CreateVirtualPrivateCloudRouteTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualPrivateCloudRouteTableParamsWithTimeout creates a new CreateVirtualPrivateCloudRouteTableParams object
// with the ability to set a timeout on a request.
func NewCreateVirtualPrivateCloudRouteTableParamsWithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudRouteTableParams {
	return &CreateVirtualPrivateCloudRouteTableParams{
		timeout: timeout,
	}
}

// NewCreateVirtualPrivateCloudRouteTableParamsWithContext creates a new CreateVirtualPrivateCloudRouteTableParams object
// with the ability to set a context for a request.
func NewCreateVirtualPrivateCloudRouteTableParamsWithContext(ctx context.Context) *CreateVirtualPrivateCloudRouteTableParams {
	return &CreateVirtualPrivateCloudRouteTableParams{
		Context: ctx,
	}
}

// NewCreateVirtualPrivateCloudRouteTableParamsWithHTTPClient creates a new CreateVirtualPrivateCloudRouteTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVirtualPrivateCloudRouteTableParamsWithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudRouteTableParams {
	return &CreateVirtualPrivateCloudRouteTableParams{
		HTTPClient: client,
	}
}

/* CreateVirtualPrivateCloudRouteTableParams contains all the parameters to send to the API endpoint
   for the create virtual private cloud route table operation.

   Typically these are written to a http.Request.
*/
type CreateVirtualPrivateCloudRouteTableParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VirtualPrivateCloudRouteTableCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create virtual private cloud route table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudRouteTableParams) WithDefaults() *CreateVirtualPrivateCloudRouteTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create virtual private cloud route table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudRouteTableParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVirtualPrivateCloudRouteTableParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) WithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudRouteTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) WithContext(ctx context.Context) *CreateVirtualPrivateCloudRouteTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) WithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudRouteTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) WithContentLanguage(contentLanguage *string) *CreateVirtualPrivateCloudRouteTableParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) WithRequestBody(requestBody []*models.VirtualPrivateCloudRouteTableCreationParams) *CreateVirtualPrivateCloudRouteTableParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create virtual private cloud route table params
func (o *CreateVirtualPrivateCloudRouteTableParams) SetRequestBody(requestBody []*models.VirtualPrivateCloudRouteTableCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualPrivateCloudRouteTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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