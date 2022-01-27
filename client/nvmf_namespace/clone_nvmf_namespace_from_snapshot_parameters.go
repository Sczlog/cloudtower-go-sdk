// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

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

// NewCloneNvmfNamespaceFromSnapshotParams creates a new CloneNvmfNamespaceFromSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneNvmfNamespaceFromSnapshotParams() *CloneNvmfNamespaceFromSnapshotParams {
	return &CloneNvmfNamespaceFromSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneNvmfNamespaceFromSnapshotParamsWithTimeout creates a new CloneNvmfNamespaceFromSnapshotParams object
// with the ability to set a timeout on a request.
func NewCloneNvmfNamespaceFromSnapshotParamsWithTimeout(timeout time.Duration) *CloneNvmfNamespaceFromSnapshotParams {
	return &CloneNvmfNamespaceFromSnapshotParams{
		timeout: timeout,
	}
}

// NewCloneNvmfNamespaceFromSnapshotParamsWithContext creates a new CloneNvmfNamespaceFromSnapshotParams object
// with the ability to set a context for a request.
func NewCloneNvmfNamespaceFromSnapshotParamsWithContext(ctx context.Context) *CloneNvmfNamespaceFromSnapshotParams {
	return &CloneNvmfNamespaceFromSnapshotParams{
		Context: ctx,
	}
}

// NewCloneNvmfNamespaceFromSnapshotParamsWithHTTPClient creates a new CloneNvmfNamespaceFromSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneNvmfNamespaceFromSnapshotParamsWithHTTPClient(client *http.Client) *CloneNvmfNamespaceFromSnapshotParams {
	return &CloneNvmfNamespaceFromSnapshotParams{
		HTTPClient: client,
	}
}

/* CloneNvmfNamespaceFromSnapshotParams contains all the parameters to send to the API endpoint
   for the clone nvmf namespace from snapshot operation.

   Typically these are written to a http.Request.
*/
type CloneNvmfNamespaceFromSnapshotParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.NvmfNamespaceCloneParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone nvmf namespace from snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneNvmfNamespaceFromSnapshotParams) WithDefaults() *CloneNvmfNamespaceFromSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone nvmf namespace from snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneNvmfNamespaceFromSnapshotParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CloneNvmfNamespaceFromSnapshotParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) WithTimeout(timeout time.Duration) *CloneNvmfNamespaceFromSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) WithContext(ctx context.Context) *CloneNvmfNamespaceFromSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) WithHTTPClient(client *http.Client) *CloneNvmfNamespaceFromSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) WithContentLanguage(contentLanguage *string) *CloneNvmfNamespaceFromSnapshotParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) WithRequestBody(requestBody []*models.NvmfNamespaceCloneParams) *CloneNvmfNamespaceFromSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the clone nvmf namespace from snapshot params
func (o *CloneNvmfNamespaceFromSnapshotParams) SetRequestBody(requestBody []*models.NvmfNamespaceCloneParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CloneNvmfNamespaceFromSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
