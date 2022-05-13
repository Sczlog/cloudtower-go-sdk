// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

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

// NewGetIscsiLunSnapshotsParams creates a new GetIscsiLunSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIscsiLunSnapshotsParams() *GetIscsiLunSnapshotsParams {
	return &GetIscsiLunSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIscsiLunSnapshotsParamsWithTimeout creates a new GetIscsiLunSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetIscsiLunSnapshotsParamsWithTimeout(timeout time.Duration) *GetIscsiLunSnapshotsParams {
	return &GetIscsiLunSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetIscsiLunSnapshotsParamsWithContext creates a new GetIscsiLunSnapshotsParams object
// with the ability to set a context for a request.
func NewGetIscsiLunSnapshotsParamsWithContext(ctx context.Context) *GetIscsiLunSnapshotsParams {
	return &GetIscsiLunSnapshotsParams{
		Context: ctx,
	}
}

// NewGetIscsiLunSnapshotsParamsWithHTTPClient creates a new GetIscsiLunSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIscsiLunSnapshotsParamsWithHTTPClient(client *http.Client) *GetIscsiLunSnapshotsParams {
	return &GetIscsiLunSnapshotsParams{
		HTTPClient: client,
	}
}

/* GetIscsiLunSnapshotsParams contains all the parameters to send to the API endpoint
   for the get iscsi lun snapshots operation.

   Typically these are written to a http.Request.
*/
type GetIscsiLunSnapshotsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetIscsiLunSnapshotsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get iscsi lun snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIscsiLunSnapshotsParams) WithDefaults() *GetIscsiLunSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get iscsi lun snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIscsiLunSnapshotsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetIscsiLunSnapshotsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) WithTimeout(timeout time.Duration) *GetIscsiLunSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) WithContext(ctx context.Context) *GetIscsiLunSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) WithHTTPClient(client *http.Client) *GetIscsiLunSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) WithContentLanguage(contentLanguage *string) *GetIscsiLunSnapshotsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) WithRequestBody(requestBody *models.GetIscsiLunSnapshotsRequestBody) *GetIscsiLunSnapshotsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get iscsi lun snapshots params
func (o *GetIscsiLunSnapshotsParams) SetRequestBody(requestBody *models.GetIscsiLunSnapshotsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetIscsiLunSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
