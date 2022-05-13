// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

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

// NewExecuteSnapshotPlanParams creates a new ExecuteSnapshotPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteSnapshotPlanParams() *ExecuteSnapshotPlanParams {
	return &ExecuteSnapshotPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteSnapshotPlanParamsWithTimeout creates a new ExecuteSnapshotPlanParams object
// with the ability to set a timeout on a request.
func NewExecuteSnapshotPlanParamsWithTimeout(timeout time.Duration) *ExecuteSnapshotPlanParams {
	return &ExecuteSnapshotPlanParams{
		timeout: timeout,
	}
}

// NewExecuteSnapshotPlanParamsWithContext creates a new ExecuteSnapshotPlanParams object
// with the ability to set a context for a request.
func NewExecuteSnapshotPlanParamsWithContext(ctx context.Context) *ExecuteSnapshotPlanParams {
	return &ExecuteSnapshotPlanParams{
		Context: ctx,
	}
}

// NewExecuteSnapshotPlanParamsWithHTTPClient creates a new ExecuteSnapshotPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteSnapshotPlanParamsWithHTTPClient(client *http.Client) *ExecuteSnapshotPlanParams {
	return &ExecuteSnapshotPlanParams{
		HTTPClient: client,
	}
}

/* ExecuteSnapshotPlanParams contains all the parameters to send to the API endpoint
   for the execute snapshot plan operation.

   Typically these are written to a http.Request.
*/
type ExecuteSnapshotPlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnapshotPlanExecutionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteSnapshotPlanParams) WithDefaults() *ExecuteSnapshotPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteSnapshotPlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := ExecuteSnapshotPlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) WithTimeout(timeout time.Duration) *ExecuteSnapshotPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) WithContext(ctx context.Context) *ExecuteSnapshotPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) WithHTTPClient(client *http.Client) *ExecuteSnapshotPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) WithContentLanguage(contentLanguage *string) *ExecuteSnapshotPlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) WithRequestBody(requestBody *models.SnapshotPlanExecutionParams) *ExecuteSnapshotPlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the execute snapshot plan params
func (o *ExecuteSnapshotPlanParams) SetRequestBody(requestBody *models.SnapshotPlanExecutionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteSnapshotPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
