// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

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

// NewCreateVirtualPrivateCloudParams creates a new CreateVirtualPrivateCloudParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVirtualPrivateCloudParams() *CreateVirtualPrivateCloudParams {
	return &CreateVirtualPrivateCloudParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualPrivateCloudParamsWithTimeout creates a new CreateVirtualPrivateCloudParams object
// with the ability to set a timeout on a request.
func NewCreateVirtualPrivateCloudParamsWithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudParams {
	return &CreateVirtualPrivateCloudParams{
		timeout: timeout,
	}
}

// NewCreateVirtualPrivateCloudParamsWithContext creates a new CreateVirtualPrivateCloudParams object
// with the ability to set a context for a request.
func NewCreateVirtualPrivateCloudParamsWithContext(ctx context.Context) *CreateVirtualPrivateCloudParams {
	return &CreateVirtualPrivateCloudParams{
		Context: ctx,
	}
}

// NewCreateVirtualPrivateCloudParamsWithHTTPClient creates a new CreateVirtualPrivateCloudParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVirtualPrivateCloudParamsWithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudParams {
	return &CreateVirtualPrivateCloudParams{
		HTTPClient: client,
	}
}

/* CreateVirtualPrivateCloudParams contains all the parameters to send to the API endpoint
   for the create virtual private cloud operation.

   Typically these are written to a http.Request.
*/
type CreateVirtualPrivateCloudParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VirtualPrivateCloudCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create virtual private cloud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudParams) WithDefaults() *CreateVirtualPrivateCloudParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create virtual private cloud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVirtualPrivateCloudParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) WithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) WithContext(ctx context.Context) *CreateVirtualPrivateCloudParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) WithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) WithContentLanguage(contentLanguage *string) *CreateVirtualPrivateCloudParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) WithRequestBody(requestBody []*models.VirtualPrivateCloudCreationParams) *CreateVirtualPrivateCloudParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create virtual private cloud params
func (o *CreateVirtualPrivateCloudParams) SetRequestBody(requestBody []*models.VirtualPrivateCloudCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualPrivateCloudParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
