// Code generated by go-swagger; DO NOT EDIT.

package backup_plan

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

// NewGetBackupRestorePointMetadataParams creates a new GetBackupRestorePointMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupRestorePointMetadataParams() *GetBackupRestorePointMetadataParams {
	return &GetBackupRestorePointMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupRestorePointMetadataParamsWithTimeout creates a new GetBackupRestorePointMetadataParams object
// with the ability to set a timeout on a request.
func NewGetBackupRestorePointMetadataParamsWithTimeout(timeout time.Duration) *GetBackupRestorePointMetadataParams {
	return &GetBackupRestorePointMetadataParams{
		timeout: timeout,
	}
}

// NewGetBackupRestorePointMetadataParamsWithContext creates a new GetBackupRestorePointMetadataParams object
// with the ability to set a context for a request.
func NewGetBackupRestorePointMetadataParamsWithContext(ctx context.Context) *GetBackupRestorePointMetadataParams {
	return &GetBackupRestorePointMetadataParams{
		Context: ctx,
	}
}

// NewGetBackupRestorePointMetadataParamsWithHTTPClient creates a new GetBackupRestorePointMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupRestorePointMetadataParamsWithHTTPClient(client *http.Client) *GetBackupRestorePointMetadataParams {
	return &GetBackupRestorePointMetadataParams{
		HTTPClient: client,
	}
}

/* GetBackupRestorePointMetadataParams contains all the parameters to send to the API endpoint
   for the get backup restore point metadata operation.

   Typically these are written to a http.Request.
*/
type GetBackupRestorePointMetadataParams struct {

	// RequestBody.
	RequestBody *models.GetBackupRestorePointMetadataRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup restore point metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestorePointMetadataParams) WithDefaults() *GetBackupRestorePointMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup restore point metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestorePointMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) WithTimeout(timeout time.Duration) *GetBackupRestorePointMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) WithContext(ctx context.Context) *GetBackupRestorePointMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) WithHTTPClient(client *http.Client) *GetBackupRestorePointMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) WithRequestBody(requestBody *models.GetBackupRestorePointMetadataRequestBody) *GetBackupRestorePointMetadataParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup restore point metadata params
func (o *GetBackupRestorePointMetadataParams) SetRequestBody(requestBody *models.GetBackupRestorePointMetadataRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupRestorePointMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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