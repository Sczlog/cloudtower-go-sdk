// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_execution

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

// NewGetBackupRestoreExecutionsParams creates a new GetBackupRestoreExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupRestoreExecutionsParams() *GetBackupRestoreExecutionsParams {
	return &GetBackupRestoreExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupRestoreExecutionsParamsWithTimeout creates a new GetBackupRestoreExecutionsParams object
// with the ability to set a timeout on a request.
func NewGetBackupRestoreExecutionsParamsWithTimeout(timeout time.Duration) *GetBackupRestoreExecutionsParams {
	return &GetBackupRestoreExecutionsParams{
		timeout: timeout,
	}
}

// NewGetBackupRestoreExecutionsParamsWithContext creates a new GetBackupRestoreExecutionsParams object
// with the ability to set a context for a request.
func NewGetBackupRestoreExecutionsParamsWithContext(ctx context.Context) *GetBackupRestoreExecutionsParams {
	return &GetBackupRestoreExecutionsParams{
		Context: ctx,
	}
}

// NewGetBackupRestoreExecutionsParamsWithHTTPClient creates a new GetBackupRestoreExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupRestoreExecutionsParamsWithHTTPClient(client *http.Client) *GetBackupRestoreExecutionsParams {
	return &GetBackupRestoreExecutionsParams{
		HTTPClient: client,
	}
}

/* GetBackupRestoreExecutionsParams contains all the parameters to send to the API endpoint
   for the get backup restore executions operation.

   Typically these are written to a http.Request.
*/
type GetBackupRestoreExecutionsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBackupRestoreExecutionsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup restore executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestoreExecutionsParams) WithDefaults() *GetBackupRestoreExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup restore executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestoreExecutionsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBackupRestoreExecutionsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) WithTimeout(timeout time.Duration) *GetBackupRestoreExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) WithContext(ctx context.Context) *GetBackupRestoreExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) WithHTTPClient(client *http.Client) *GetBackupRestoreExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) WithContentLanguage(contentLanguage *string) *GetBackupRestoreExecutionsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) WithRequestBody(requestBody *models.GetBackupRestoreExecutionsRequestBody) *GetBackupRestoreExecutionsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup restore executions params
func (o *GetBackupRestoreExecutionsParams) SetRequestBody(requestBody *models.GetBackupRestoreExecutionsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupRestoreExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
