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

// NewGetVirtualPrivateCloudSecurityPoliciesParams creates a new GetVirtualPrivateCloudSecurityPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudSecurityPoliciesParams() *GetVirtualPrivateCloudSecurityPoliciesParams {
	return &GetVirtualPrivateCloudSecurityPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudSecurityPoliciesParamsWithTimeout creates a new GetVirtualPrivateCloudSecurityPoliciesParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudSecurityPoliciesParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudSecurityPoliciesParams {
	return &GetVirtualPrivateCloudSecurityPoliciesParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudSecurityPoliciesParamsWithContext creates a new GetVirtualPrivateCloudSecurityPoliciesParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudSecurityPoliciesParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudSecurityPoliciesParams {
	return &GetVirtualPrivateCloudSecurityPoliciesParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudSecurityPoliciesParamsWithHTTPClient creates a new GetVirtualPrivateCloudSecurityPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudSecurityPoliciesParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudSecurityPoliciesParams {
	return &GetVirtualPrivateCloudSecurityPoliciesParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudSecurityPoliciesParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud security policies operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudSecurityPoliciesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudSecurityPoliciesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud security policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithDefaults() *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud security policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudSecurityPoliciesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudSecurityPoliciesRequestBody) *GetVirtualPrivateCloudSecurityPoliciesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud security policies params
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudSecurityPoliciesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudSecurityPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
