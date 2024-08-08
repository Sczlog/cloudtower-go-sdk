// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

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

// NewDeleteVirtualPrivateCloudParams creates a new DeleteVirtualPrivateCloudParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualPrivateCloudParams() *DeleteVirtualPrivateCloudParams {
	return &DeleteVirtualPrivateCloudParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualPrivateCloudParamsWithTimeout creates a new DeleteVirtualPrivateCloudParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualPrivateCloudParamsWithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudParams {
	return &DeleteVirtualPrivateCloudParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualPrivateCloudParamsWithContext creates a new DeleteVirtualPrivateCloudParams object
// with the ability to set a context for a request.
func NewDeleteVirtualPrivateCloudParamsWithContext(ctx context.Context) *DeleteVirtualPrivateCloudParams {
	return &DeleteVirtualPrivateCloudParams{
		Context: ctx,
	}
}

// NewDeleteVirtualPrivateCloudParamsWithHTTPClient creates a new DeleteVirtualPrivateCloudParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualPrivateCloudParamsWithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudParams {
	return &DeleteVirtualPrivateCloudParams{
		HTTPClient: client,
	}
}

/* DeleteVirtualPrivateCloudParams contains all the parameters to send to the API endpoint
   for the delete virtual private cloud operation.

   Typically these are written to a http.Request.
*/
type DeleteVirtualPrivateCloudParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VirtualPrivateCloudDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual private cloud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudParams) WithDefaults() *DeleteVirtualPrivateCloudParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual private cloud params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVirtualPrivateCloudParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) WithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) WithContext(ctx context.Context) *DeleteVirtualPrivateCloudParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) WithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) WithContentLanguage(contentLanguage *string) *DeleteVirtualPrivateCloudParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) WithRequestBody(requestBody *models.VirtualPrivateCloudDeletionParams) *DeleteVirtualPrivateCloudParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete virtual private cloud params
func (o *DeleteVirtualPrivateCloudParams) SetRequestBody(requestBody *models.VirtualPrivateCloudDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualPrivateCloudParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
