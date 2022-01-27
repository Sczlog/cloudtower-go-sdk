// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan_task

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

// NewGetSnapshotPlanTasksParams creates a new GetSnapshotPlanTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotPlanTasksParams() *GetSnapshotPlanTasksParams {
	return &GetSnapshotPlanTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotPlanTasksParamsWithTimeout creates a new GetSnapshotPlanTasksParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotPlanTasksParamsWithTimeout(timeout time.Duration) *GetSnapshotPlanTasksParams {
	return &GetSnapshotPlanTasksParams{
		timeout: timeout,
	}
}

// NewGetSnapshotPlanTasksParamsWithContext creates a new GetSnapshotPlanTasksParams object
// with the ability to set a context for a request.
func NewGetSnapshotPlanTasksParamsWithContext(ctx context.Context) *GetSnapshotPlanTasksParams {
	return &GetSnapshotPlanTasksParams{
		Context: ctx,
	}
}

// NewGetSnapshotPlanTasksParamsWithHTTPClient creates a new GetSnapshotPlanTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotPlanTasksParamsWithHTTPClient(client *http.Client) *GetSnapshotPlanTasksParams {
	return &GetSnapshotPlanTasksParams{
		HTTPClient: client,
	}
}

/* GetSnapshotPlanTasksParams contains all the parameters to send to the API endpoint
   for the get snapshot plan tasks operation.

   Typically these are written to a http.Request.
*/
type GetSnapshotPlanTasksParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetSnapshotPlanTasksRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot plan tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotPlanTasksParams) WithDefaults() *GetSnapshotPlanTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot plan tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotPlanTasksParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetSnapshotPlanTasksParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) WithTimeout(timeout time.Duration) *GetSnapshotPlanTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) WithContext(ctx context.Context) *GetSnapshotPlanTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) WithHTTPClient(client *http.Client) *GetSnapshotPlanTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) WithContentLanguage(contentLanguage *string) *GetSnapshotPlanTasksParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) WithRequestBody(requestBody *models.GetSnapshotPlanTasksRequestBody) *GetSnapshotPlanTasksParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get snapshot plan tasks params
func (o *GetSnapshotPlanTasksParams) SetRequestBody(requestBody *models.GetSnapshotPlanTasksRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotPlanTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
