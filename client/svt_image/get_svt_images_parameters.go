// Code generated by go-swagger; DO NOT EDIT.

package svt_image

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

// NewGetSvtImagesParams creates a new GetSvtImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSvtImagesParams() *GetSvtImagesParams {
	return &GetSvtImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSvtImagesParamsWithTimeout creates a new GetSvtImagesParams object
// with the ability to set a timeout on a request.
func NewGetSvtImagesParamsWithTimeout(timeout time.Duration) *GetSvtImagesParams {
	return &GetSvtImagesParams{
		timeout: timeout,
	}
}

// NewGetSvtImagesParamsWithContext creates a new GetSvtImagesParams object
// with the ability to set a context for a request.
func NewGetSvtImagesParamsWithContext(ctx context.Context) *GetSvtImagesParams {
	return &GetSvtImagesParams{
		Context: ctx,
	}
}

// NewGetSvtImagesParamsWithHTTPClient creates a new GetSvtImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSvtImagesParamsWithHTTPClient(client *http.Client) *GetSvtImagesParams {
	return &GetSvtImagesParams{
		HTTPClient: client,
	}
}

/* GetSvtImagesParams contains all the parameters to send to the API endpoint
   for the get svt images operation.

   Typically these are written to a http.Request.
*/
type GetSvtImagesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetSvtImagesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get svt images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSvtImagesParams) WithDefaults() *GetSvtImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get svt images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSvtImagesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetSvtImagesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get svt images params
func (o *GetSvtImagesParams) WithTimeout(timeout time.Duration) *GetSvtImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get svt images params
func (o *GetSvtImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get svt images params
func (o *GetSvtImagesParams) WithContext(ctx context.Context) *GetSvtImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get svt images params
func (o *GetSvtImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get svt images params
func (o *GetSvtImagesParams) WithHTTPClient(client *http.Client) *GetSvtImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get svt images params
func (o *GetSvtImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get svt images params
func (o *GetSvtImagesParams) WithContentLanguage(contentLanguage *string) *GetSvtImagesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get svt images params
func (o *GetSvtImagesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get svt images params
func (o *GetSvtImagesParams) WithRequestBody(requestBody *models.GetSvtImagesRequestBody) *GetSvtImagesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get svt images params
func (o *GetSvtImagesParams) SetRequestBody(requestBody *models.GetSvtImagesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSvtImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
