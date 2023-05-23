// Code generated by go-swagger; DO NOT EDIT.

package security_group

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

// NewCreateSecurityGroupParams creates a new CreateSecurityGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSecurityGroupParams() *CreateSecurityGroupParams {
	return &CreateSecurityGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecurityGroupParamsWithTimeout creates a new CreateSecurityGroupParams object
// with the ability to set a timeout on a request.
func NewCreateSecurityGroupParamsWithTimeout(timeout time.Duration) *CreateSecurityGroupParams {
	return &CreateSecurityGroupParams{
		timeout: timeout,
	}
}

// NewCreateSecurityGroupParamsWithContext creates a new CreateSecurityGroupParams object
// with the ability to set a context for a request.
func NewCreateSecurityGroupParamsWithContext(ctx context.Context) *CreateSecurityGroupParams {
	return &CreateSecurityGroupParams{
		Context: ctx,
	}
}

// NewCreateSecurityGroupParamsWithHTTPClient creates a new CreateSecurityGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSecurityGroupParamsWithHTTPClient(client *http.Client) *CreateSecurityGroupParams {
	return &CreateSecurityGroupParams{
		HTTPClient: client,
	}
}

/* CreateSecurityGroupParams contains all the parameters to send to the API endpoint
   for the create security group operation.

   Typically these are written to a http.Request.
*/
type CreateSecurityGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SecurityGroupCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecurityGroupParams) WithDefaults() *CreateSecurityGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecurityGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateSecurityGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create security group params
func (o *CreateSecurityGroupParams) WithTimeout(timeout time.Duration) *CreateSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create security group params
func (o *CreateSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create security group params
func (o *CreateSecurityGroupParams) WithContext(ctx context.Context) *CreateSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create security group params
func (o *CreateSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create security group params
func (o *CreateSecurityGroupParams) WithHTTPClient(client *http.Client) *CreateSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create security group params
func (o *CreateSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create security group params
func (o *CreateSecurityGroupParams) WithContentLanguage(contentLanguage *string) *CreateSecurityGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create security group params
func (o *CreateSecurityGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create security group params
func (o *CreateSecurityGroupParams) WithRequestBody(requestBody *models.SecurityGroupCreateParams) *CreateSecurityGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create security group params
func (o *CreateSecurityGroupParams) SetRequestBody(requestBody *models.SecurityGroupCreateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
