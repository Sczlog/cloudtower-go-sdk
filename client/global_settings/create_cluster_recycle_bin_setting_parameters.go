// Code generated by go-swagger; DO NOT EDIT.

package global_settings

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

// NewCreateClusterRecycleBinSettingParams creates a new CreateClusterRecycleBinSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterRecycleBinSettingParams() *CreateClusterRecycleBinSettingParams {
	return &CreateClusterRecycleBinSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterRecycleBinSettingParamsWithTimeout creates a new CreateClusterRecycleBinSettingParams object
// with the ability to set a timeout on a request.
func NewCreateClusterRecycleBinSettingParamsWithTimeout(timeout time.Duration) *CreateClusterRecycleBinSettingParams {
	return &CreateClusterRecycleBinSettingParams{
		timeout: timeout,
	}
}

// NewCreateClusterRecycleBinSettingParamsWithContext creates a new CreateClusterRecycleBinSettingParams object
// with the ability to set a context for a request.
func NewCreateClusterRecycleBinSettingParamsWithContext(ctx context.Context) *CreateClusterRecycleBinSettingParams {
	return &CreateClusterRecycleBinSettingParams{
		Context: ctx,
	}
}

// NewCreateClusterRecycleBinSettingParamsWithHTTPClient creates a new CreateClusterRecycleBinSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterRecycleBinSettingParamsWithHTTPClient(client *http.Client) *CreateClusterRecycleBinSettingParams {
	return &CreateClusterRecycleBinSettingParams{
		HTTPClient: client,
	}
}

/* CreateClusterRecycleBinSettingParams contains all the parameters to send to the API endpoint
   for the create cluster recycle bin setting operation.

   Typically these are written to a http.Request.
*/
type CreateClusterRecycleBinSettingParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ClusterRecycleBinCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster recycle bin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterRecycleBinSettingParams) WithDefaults() *CreateClusterRecycleBinSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster recycle bin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterRecycleBinSettingParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateClusterRecycleBinSettingParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) WithTimeout(timeout time.Duration) *CreateClusterRecycleBinSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) WithContext(ctx context.Context) *CreateClusterRecycleBinSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) WithHTTPClient(client *http.Client) *CreateClusterRecycleBinSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) WithContentLanguage(contentLanguage *string) *CreateClusterRecycleBinSettingParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) WithRequestBody(requestBody *models.ClusterRecycleBinCreationParams) *CreateClusterRecycleBinSettingParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create cluster recycle bin setting params
func (o *CreateClusterRecycleBinSettingParams) SetRequestBody(requestBody *models.ClusterRecycleBinCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterRecycleBinSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
