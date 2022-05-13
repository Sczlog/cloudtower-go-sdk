// Code generated by go-swagger; DO NOT EDIT.

package upload_task

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

// NewGetUploadTasksConnectionParams creates a new GetUploadTasksConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUploadTasksConnectionParams() *GetUploadTasksConnectionParams {
	return &GetUploadTasksConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUploadTasksConnectionParamsWithTimeout creates a new GetUploadTasksConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUploadTasksConnectionParamsWithTimeout(timeout time.Duration) *GetUploadTasksConnectionParams {
	return &GetUploadTasksConnectionParams{
		timeout: timeout,
	}
}

// NewGetUploadTasksConnectionParamsWithContext creates a new GetUploadTasksConnectionParams object
// with the ability to set a context for a request.
func NewGetUploadTasksConnectionParamsWithContext(ctx context.Context) *GetUploadTasksConnectionParams {
	return &GetUploadTasksConnectionParams{
		Context: ctx,
	}
}

// NewGetUploadTasksConnectionParamsWithHTTPClient creates a new GetUploadTasksConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUploadTasksConnectionParamsWithHTTPClient(client *http.Client) *GetUploadTasksConnectionParams {
	return &GetUploadTasksConnectionParams{
		HTTPClient: client,
	}
}

/* GetUploadTasksConnectionParams contains all the parameters to send to the API endpoint
   for the get upload tasks connection operation.

   Typically these are written to a http.Request.
*/
type GetUploadTasksConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUploadTasksConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upload tasks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUploadTasksConnectionParams) WithDefaults() *GetUploadTasksConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upload tasks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUploadTasksConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUploadTasksConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) WithTimeout(timeout time.Duration) *GetUploadTasksConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) WithContext(ctx context.Context) *GetUploadTasksConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) WithHTTPClient(client *http.Client) *GetUploadTasksConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) WithContentLanguage(contentLanguage *string) *GetUploadTasksConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) WithRequestBody(requestBody *models.GetUploadTasksConnectionRequestBody) *GetUploadTasksConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get upload tasks connection params
func (o *GetUploadTasksConnectionParams) SetRequestBody(requestBody *models.GetUploadTasksConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUploadTasksConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
