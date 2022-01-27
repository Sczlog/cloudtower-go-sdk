// Code generated by go-swagger; DO NOT EDIT.

package global_alert_rule

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

// NewUpdateCustomizeAlertRuleParams creates a new UpdateCustomizeAlertRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCustomizeAlertRuleParams() *UpdateCustomizeAlertRuleParams {
	return &UpdateCustomizeAlertRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomizeAlertRuleParamsWithTimeout creates a new UpdateCustomizeAlertRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateCustomizeAlertRuleParamsWithTimeout(timeout time.Duration) *UpdateCustomizeAlertRuleParams {
	return &UpdateCustomizeAlertRuleParams{
		timeout: timeout,
	}
}

// NewUpdateCustomizeAlertRuleParamsWithContext creates a new UpdateCustomizeAlertRuleParams object
// with the ability to set a context for a request.
func NewUpdateCustomizeAlertRuleParamsWithContext(ctx context.Context) *UpdateCustomizeAlertRuleParams {
	return &UpdateCustomizeAlertRuleParams{
		Context: ctx,
	}
}

// NewUpdateCustomizeAlertRuleParamsWithHTTPClient creates a new UpdateCustomizeAlertRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCustomizeAlertRuleParamsWithHTTPClient(client *http.Client) *UpdateCustomizeAlertRuleParams {
	return &UpdateCustomizeAlertRuleParams{
		HTTPClient: client,
	}
}

/* UpdateCustomizeAlertRuleParams contains all the parameters to send to the API endpoint
   for the update customize alert rule operation.

   Typically these are written to a http.Request.
*/
type UpdateCustomizeAlertRuleParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.CustomizeAlertRuleUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update customize alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomizeAlertRuleParams) WithDefaults() *UpdateCustomizeAlertRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update customize alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomizeAlertRuleParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateCustomizeAlertRuleParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) WithTimeout(timeout time.Duration) *UpdateCustomizeAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) WithContext(ctx context.Context) *UpdateCustomizeAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) WithHTTPClient(client *http.Client) *UpdateCustomizeAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) WithContentLanguage(contentLanguage *string) *UpdateCustomizeAlertRuleParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) WithRequestBody(requestBody *models.CustomizeAlertRuleUpdationParams) *UpdateCustomizeAlertRuleParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update customize alert rule params
func (o *UpdateCustomizeAlertRuleParams) SetRequestBody(requestBody *models.CustomizeAlertRuleUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomizeAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
