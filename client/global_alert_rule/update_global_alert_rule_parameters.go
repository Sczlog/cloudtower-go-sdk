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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewUpdateGlobalAlertRuleParams creates a new UpdateGlobalAlertRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGlobalAlertRuleParams() *UpdateGlobalAlertRuleParams {
	return &UpdateGlobalAlertRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGlobalAlertRuleParamsWithTimeout creates a new UpdateGlobalAlertRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateGlobalAlertRuleParamsWithTimeout(timeout time.Duration) *UpdateGlobalAlertRuleParams {
	return &UpdateGlobalAlertRuleParams{
		timeout: timeout,
	}
}

// NewUpdateGlobalAlertRuleParamsWithContext creates a new UpdateGlobalAlertRuleParams object
// with the ability to set a context for a request.
func NewUpdateGlobalAlertRuleParamsWithContext(ctx context.Context) *UpdateGlobalAlertRuleParams {
	return &UpdateGlobalAlertRuleParams{
		Context: ctx,
	}
}

// NewUpdateGlobalAlertRuleParamsWithHTTPClient creates a new UpdateGlobalAlertRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGlobalAlertRuleParamsWithHTTPClient(client *http.Client) *UpdateGlobalAlertRuleParams {
	return &UpdateGlobalAlertRuleParams{
		HTTPClient: client,
	}
}

/* UpdateGlobalAlertRuleParams contains all the parameters to send to the API endpoint
   for the update global alert rule operation.

   Typically these are written to a http.Request.
*/
type UpdateGlobalAlertRuleParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GlobalAlertRuleUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update global alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalAlertRuleParams) WithDefaults() *UpdateGlobalAlertRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update global alert rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalAlertRuleParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateGlobalAlertRuleParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) WithTimeout(timeout time.Duration) *UpdateGlobalAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) WithContext(ctx context.Context) *UpdateGlobalAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) WithHTTPClient(client *http.Client) *UpdateGlobalAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) WithContentLanguage(contentLanguage *string) *UpdateGlobalAlertRuleParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) WithRequestBody(requestBody *models.GlobalAlertRuleUpdationParams) *UpdateGlobalAlertRuleParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update global alert rule params
func (o *UpdateGlobalAlertRuleParams) SetRequestBody(requestBody *models.GlobalAlertRuleUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGlobalAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
