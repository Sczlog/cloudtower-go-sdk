// Code generated by go-swagger; DO NOT EDIT.

package consistency_group

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

// NewGetConsistencyGroupsConnectionParams creates a new GetConsistencyGroupsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConsistencyGroupsConnectionParams() *GetConsistencyGroupsConnectionParams {
	return &GetConsistencyGroupsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsistencyGroupsConnectionParamsWithTimeout creates a new GetConsistencyGroupsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetConsistencyGroupsConnectionParamsWithTimeout(timeout time.Duration) *GetConsistencyGroupsConnectionParams {
	return &GetConsistencyGroupsConnectionParams{
		timeout: timeout,
	}
}

// NewGetConsistencyGroupsConnectionParamsWithContext creates a new GetConsistencyGroupsConnectionParams object
// with the ability to set a context for a request.
func NewGetConsistencyGroupsConnectionParamsWithContext(ctx context.Context) *GetConsistencyGroupsConnectionParams {
	return &GetConsistencyGroupsConnectionParams{
		Context: ctx,
	}
}

// NewGetConsistencyGroupsConnectionParamsWithHTTPClient creates a new GetConsistencyGroupsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConsistencyGroupsConnectionParamsWithHTTPClient(client *http.Client) *GetConsistencyGroupsConnectionParams {
	return &GetConsistencyGroupsConnectionParams{
		HTTPClient: client,
	}
}

/* GetConsistencyGroupsConnectionParams contains all the parameters to send to the API endpoint
   for the get consistency groups connection operation.

   Typically these are written to a http.Request.
*/
type GetConsistencyGroupsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetConsistencyGroupsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get consistency groups connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsistencyGroupsConnectionParams) WithDefaults() *GetConsistencyGroupsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get consistency groups connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsistencyGroupsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetConsistencyGroupsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) WithTimeout(timeout time.Duration) *GetConsistencyGroupsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) WithContext(ctx context.Context) *GetConsistencyGroupsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) WithHTTPClient(client *http.Client) *GetConsistencyGroupsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) WithContentLanguage(contentLanguage *string) *GetConsistencyGroupsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) WithRequestBody(requestBody *models.GetConsistencyGroupsConnectionRequestBody) *GetConsistencyGroupsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get consistency groups connection params
func (o *GetConsistencyGroupsConnectionParams) SetRequestBody(requestBody *models.GetConsistencyGroupsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsistencyGroupsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
