// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_security_policy

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

// NewDeleteVirtualPrivateCloudSecurityPolicyParams creates a new DeleteVirtualPrivateCloudSecurityPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualPrivateCloudSecurityPolicyParams() *DeleteVirtualPrivateCloudSecurityPolicyParams {
	return &DeleteVirtualPrivateCloudSecurityPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithTimeout creates a new DeleteVirtualPrivateCloudSecurityPolicyParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	return &DeleteVirtualPrivateCloudSecurityPolicyParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithContext creates a new DeleteVirtualPrivateCloudSecurityPolicyParams object
// with the ability to set a context for a request.
func NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithContext(ctx context.Context) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	return &DeleteVirtualPrivateCloudSecurityPolicyParams{
		Context: ctx,
	}
}

// NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithHTTPClient creates a new DeleteVirtualPrivateCloudSecurityPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualPrivateCloudSecurityPolicyParamsWithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	return &DeleteVirtualPrivateCloudSecurityPolicyParams{
		HTTPClient: client,
	}
}

/* DeleteVirtualPrivateCloudSecurityPolicyParams contains all the parameters to send to the API endpoint
   for the delete virtual private cloud security policy operation.

   Typically these are written to a http.Request.
*/
type DeleteVirtualPrivateCloudSecurityPolicyParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VirtualPrivateCloudSecurityPolicyDeleteParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual private cloud security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithDefaults() *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual private cloud security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVirtualPrivateCloudSecurityPolicyParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithContext(ctx context.Context) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithContentLanguage(contentLanguage *string) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WithRequestBody(requestBody *models.VirtualPrivateCloudSecurityPolicyDeleteParams) *DeleteVirtualPrivateCloudSecurityPolicyParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete virtual private cloud security policy params
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) SetRequestBody(requestBody *models.VirtualPrivateCloudSecurityPolicyDeleteParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualPrivateCloudSecurityPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
