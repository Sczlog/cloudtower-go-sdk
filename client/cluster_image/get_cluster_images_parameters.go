// Code generated by go-swagger; DO NOT EDIT.

package cluster_image

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

// NewGetClusterImagesParams creates a new GetClusterImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterImagesParams() *GetClusterImagesParams {
	return &GetClusterImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterImagesParamsWithTimeout creates a new GetClusterImagesParams object
// with the ability to set a timeout on a request.
func NewGetClusterImagesParamsWithTimeout(timeout time.Duration) *GetClusterImagesParams {
	return &GetClusterImagesParams{
		timeout: timeout,
	}
}

// NewGetClusterImagesParamsWithContext creates a new GetClusterImagesParams object
// with the ability to set a context for a request.
func NewGetClusterImagesParamsWithContext(ctx context.Context) *GetClusterImagesParams {
	return &GetClusterImagesParams{
		Context: ctx,
	}
}

// NewGetClusterImagesParamsWithHTTPClient creates a new GetClusterImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterImagesParamsWithHTTPClient(client *http.Client) *GetClusterImagesParams {
	return &GetClusterImagesParams{
		HTTPClient: client,
	}
}

/* GetClusterImagesParams contains all the parameters to send to the API endpoint
   for the get cluster images operation.

   Typically these are written to a http.Request.
*/
type GetClusterImagesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetClusterImagesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterImagesParams) WithDefaults() *GetClusterImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterImagesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetClusterImagesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cluster images params
func (o *GetClusterImagesParams) WithTimeout(timeout time.Duration) *GetClusterImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster images params
func (o *GetClusterImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster images params
func (o *GetClusterImagesParams) WithContext(ctx context.Context) *GetClusterImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster images params
func (o *GetClusterImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster images params
func (o *GetClusterImagesParams) WithHTTPClient(client *http.Client) *GetClusterImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster images params
func (o *GetClusterImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get cluster images params
func (o *GetClusterImagesParams) WithContentLanguage(contentLanguage *string) *GetClusterImagesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get cluster images params
func (o *GetClusterImagesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get cluster images params
func (o *GetClusterImagesParams) WithRequestBody(requestBody *models.GetClusterImagesRequestBody) *GetClusterImagesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster images params
func (o *GetClusterImagesParams) SetRequestBody(requestBody *models.GetClusterImagesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
