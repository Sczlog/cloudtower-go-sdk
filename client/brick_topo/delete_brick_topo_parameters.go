// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

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

// NewDeleteBrickTopoParams creates a new DeleteBrickTopoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBrickTopoParams() *DeleteBrickTopoParams {
	return &DeleteBrickTopoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBrickTopoParamsWithTimeout creates a new DeleteBrickTopoParams object
// with the ability to set a timeout on a request.
func NewDeleteBrickTopoParamsWithTimeout(timeout time.Duration) *DeleteBrickTopoParams {
	return &DeleteBrickTopoParams{
		timeout: timeout,
	}
}

// NewDeleteBrickTopoParamsWithContext creates a new DeleteBrickTopoParams object
// with the ability to set a context for a request.
func NewDeleteBrickTopoParamsWithContext(ctx context.Context) *DeleteBrickTopoParams {
	return &DeleteBrickTopoParams{
		Context: ctx,
	}
}

// NewDeleteBrickTopoParamsWithHTTPClient creates a new DeleteBrickTopoParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBrickTopoParamsWithHTTPClient(client *http.Client) *DeleteBrickTopoParams {
	return &DeleteBrickTopoParams{
		HTTPClient: client,
	}
}

/* DeleteBrickTopoParams contains all the parameters to send to the API endpoint
   for the delete brick topo operation.

   Typically these are written to a http.Request.
*/
type DeleteBrickTopoParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.BrickTopoDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBrickTopoParams) WithDefaults() *DeleteBrickTopoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBrickTopoParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteBrickTopoParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete brick topo params
func (o *DeleteBrickTopoParams) WithTimeout(timeout time.Duration) *DeleteBrickTopoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete brick topo params
func (o *DeleteBrickTopoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete brick topo params
func (o *DeleteBrickTopoParams) WithContext(ctx context.Context) *DeleteBrickTopoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete brick topo params
func (o *DeleteBrickTopoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete brick topo params
func (o *DeleteBrickTopoParams) WithHTTPClient(client *http.Client) *DeleteBrickTopoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete brick topo params
func (o *DeleteBrickTopoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete brick topo params
func (o *DeleteBrickTopoParams) WithContentLanguage(contentLanguage *string) *DeleteBrickTopoParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete brick topo params
func (o *DeleteBrickTopoParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete brick topo params
func (o *DeleteBrickTopoParams) WithRequestBody(requestBody *models.BrickTopoDeletionParams) *DeleteBrickTopoParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete brick topo params
func (o *DeleteBrickTopoParams) SetRequestBody(requestBody *models.BrickTopoDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBrickTopoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
