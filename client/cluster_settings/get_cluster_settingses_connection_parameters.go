// Code generated by go-swagger; DO NOT EDIT.

package cluster_settings

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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewGetClusterSettingsesConnectionParams creates a new GetClusterSettingsesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterSettingsesConnectionParams() *GetClusterSettingsesConnectionParams {
	return &GetClusterSettingsesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSettingsesConnectionParamsWithTimeout creates a new GetClusterSettingsesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetClusterSettingsesConnectionParamsWithTimeout(timeout time.Duration) *GetClusterSettingsesConnectionParams {
	return &GetClusterSettingsesConnectionParams{
		timeout: timeout,
	}
}

// NewGetClusterSettingsesConnectionParamsWithContext creates a new GetClusterSettingsesConnectionParams object
// with the ability to set a context for a request.
func NewGetClusterSettingsesConnectionParamsWithContext(ctx context.Context) *GetClusterSettingsesConnectionParams {
	return &GetClusterSettingsesConnectionParams{
		Context: ctx,
	}
}

// NewGetClusterSettingsesConnectionParamsWithHTTPClient creates a new GetClusterSettingsesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterSettingsesConnectionParamsWithHTTPClient(client *http.Client) *GetClusterSettingsesConnectionParams {
	return &GetClusterSettingsesConnectionParams{
		HTTPClient: client,
	}
}

/* GetClusterSettingsesConnectionParams contains all the parameters to send to the API endpoint
   for the get cluster settingses connection operation.

   Typically these are written to a http.Request.
*/
type GetClusterSettingsesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetClusterSettingsesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster settingses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSettingsesConnectionParams) WithDefaults() *GetClusterSettingsesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster settingses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSettingsesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetClusterSettingsesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) WithTimeout(timeout time.Duration) *GetClusterSettingsesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) WithContext(ctx context.Context) *GetClusterSettingsesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) WithHTTPClient(client *http.Client) *GetClusterSettingsesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) WithContentLanguage(contentLanguage *string) *GetClusterSettingsesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) WithRequestBody(requestBody *models.GetClusterSettingsesConnectionRequestBody) *GetClusterSettingsesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster settingses connection params
func (o *GetClusterSettingsesConnectionParams) SetRequestBody(requestBody *models.GetClusterSettingsesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSettingsesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
