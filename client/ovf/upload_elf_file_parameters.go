// Code generated by go-swagger; DO NOT EDIT.

package ovf

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
)

// NewUploadElfFileParams creates a new UploadElfFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadElfFileParams() *UploadElfFileParams {
	return &UploadElfFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadElfFileParamsWithTimeout creates a new UploadElfFileParams object
// with the ability to set a timeout on a request.
func NewUploadElfFileParamsWithTimeout(timeout time.Duration) *UploadElfFileParams {
	return &UploadElfFileParams{
		timeout: timeout,
	}
}

// NewUploadElfFileParamsWithContext creates a new UploadElfFileParams object
// with the ability to set a context for a request.
func NewUploadElfFileParamsWithContext(ctx context.Context) *UploadElfFileParams {
	return &UploadElfFileParams{
		Context: ctx,
	}
}

// NewUploadElfFileParamsWithHTTPClient creates a new UploadElfFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadElfFileParamsWithHTTPClient(client *http.Client) *UploadElfFileParams {
	return &UploadElfFileParams{
		HTTPClient: client,
	}
}

/* UploadElfFileParams contains all the parameters to send to the API endpoint
   for the upload elf file operation.

   Typically these are written to a http.Request.
*/
type UploadElfFileParams struct {

	// ClusterID.
	ClusterID *string

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// File.
	File runtime.NamedReadCloser

	// Name.
	Name *string

	// Size.
	Size *string

	// SizeUnit.
	SizeUnit *string

	// UploadTaskID.
	UploadTaskID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload elf file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadElfFileParams) WithDefaults() *UploadElfFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload elf file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadElfFileParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UploadElfFileParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upload elf file params
func (o *UploadElfFileParams) WithTimeout(timeout time.Duration) *UploadElfFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload elf file params
func (o *UploadElfFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload elf file params
func (o *UploadElfFileParams) WithContext(ctx context.Context) *UploadElfFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload elf file params
func (o *UploadElfFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload elf file params
func (o *UploadElfFileParams) WithHTTPClient(client *http.Client) *UploadElfFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload elf file params
func (o *UploadElfFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the upload elf file params
func (o *UploadElfFileParams) WithClusterID(clusterID *string) *UploadElfFileParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the upload elf file params
func (o *UploadElfFileParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithContentLanguage adds the contentLanguage to the upload elf file params
func (o *UploadElfFileParams) WithContentLanguage(contentLanguage *string) *UploadElfFileParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the upload elf file params
func (o *UploadElfFileParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithFile adds the file to the upload elf file params
func (o *UploadElfFileParams) WithFile(file runtime.NamedReadCloser) *UploadElfFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload elf file params
func (o *UploadElfFileParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithName adds the name to the upload elf file params
func (o *UploadElfFileParams) WithName(name *string) *UploadElfFileParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upload elf file params
func (o *UploadElfFileParams) SetName(name *string) {
	o.Name = name
}

// WithSize adds the size to the upload elf file params
func (o *UploadElfFileParams) WithSize(size *string) *UploadElfFileParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the upload elf file params
func (o *UploadElfFileParams) SetSize(size *string) {
	o.Size = size
}

// WithSizeUnit adds the sizeUnit to the upload elf file params
func (o *UploadElfFileParams) WithSizeUnit(sizeUnit *string) *UploadElfFileParams {
	o.SetSizeUnit(sizeUnit)
	return o
}

// SetSizeUnit adds the sizeUnit to the upload elf file params
func (o *UploadElfFileParams) SetSizeUnit(sizeUnit *string) {
	o.SizeUnit = sizeUnit
}

// WithUploadTaskID adds the uploadTaskID to the upload elf file params
func (o *UploadElfFileParams) WithUploadTaskID(uploadTaskID *string) *UploadElfFileParams {
	o.SetUploadTaskID(uploadTaskID)
	return o
}

// SetUploadTaskID adds the uploadTaskId to the upload elf file params
func (o *UploadElfFileParams) SetUploadTaskID(uploadTaskID *string) {
	o.UploadTaskID = uploadTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadElfFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterID != nil {

		// form param cluster_id
		var frClusterID string
		if o.ClusterID != nil {
			frClusterID = *o.ClusterID
		}
		fClusterID := frClusterID
		if fClusterID != "" {
			if err := r.SetFormParam("cluster_id", fClusterID); err != nil {
				return err
			}
		}
	}

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// form param size
		var frSize string
		if o.Size != nil {
			frSize = *o.Size
		}
		fSize := frSize
		if fSize != "" {
			if err := r.SetFormParam("size", fSize); err != nil {
				return err
			}
		}
	}

	if o.SizeUnit != nil {

		// form param size_unit
		var frSizeUnit string
		if o.SizeUnit != nil {
			frSizeUnit = *o.SizeUnit
		}
		fSizeUnit := frSizeUnit
		if fSizeUnit != "" {
			if err := r.SetFormParam("size_unit", fSizeUnit); err != nil {
				return err
			}
		}
	}

	if o.UploadTaskID != nil {

		// form param upload_task_id
		var frUploadTaskID string
		if o.UploadTaskID != nil {
			frUploadTaskID = *o.UploadTaskID
		}
		fUploadTaskID := frUploadTaskID
		if fUploadTaskID != "" {
			if err := r.SetFormParam("upload_task_id", fUploadTaskID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}