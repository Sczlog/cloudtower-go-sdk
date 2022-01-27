// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDeleteLogCollectionParams creates a new DeleteLogCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLogCollectionParams() *DeleteLogCollectionParams {
	return &DeleteLogCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLogCollectionParamsWithTimeout creates a new DeleteLogCollectionParams object
// with the ability to set a timeout on a request.
func NewDeleteLogCollectionParamsWithTimeout(timeout time.Duration) *DeleteLogCollectionParams {
	return &DeleteLogCollectionParams{
		timeout: timeout,
	}
}

// NewDeleteLogCollectionParamsWithContext creates a new DeleteLogCollectionParams object
// with the ability to set a context for a request.
func NewDeleteLogCollectionParamsWithContext(ctx context.Context) *DeleteLogCollectionParams {
	return &DeleteLogCollectionParams{
		Context: ctx,
	}
}

// NewDeleteLogCollectionParamsWithHTTPClient creates a new DeleteLogCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLogCollectionParamsWithHTTPClient(client *http.Client) *DeleteLogCollectionParams {
	return &DeleteLogCollectionParams{
		HTTPClient: client,
	}
}

/* DeleteLogCollectionParams contains all the parameters to send to the API endpoint
   for the delete log collection operation.

   Typically these are written to a http.Request.
*/
type DeleteLogCollectionParams struct {

	// RequestBody.
	RequestBody *models.LogCollectionDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete log collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogCollectionParams) WithDefaults() *DeleteLogCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete log collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogCollectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete log collection params
func (o *DeleteLogCollectionParams) WithTimeout(timeout time.Duration) *DeleteLogCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete log collection params
func (o *DeleteLogCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete log collection params
func (o *DeleteLogCollectionParams) WithContext(ctx context.Context) *DeleteLogCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete log collection params
func (o *DeleteLogCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete log collection params
func (o *DeleteLogCollectionParams) WithHTTPClient(client *http.Client) *DeleteLogCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete log collection params
func (o *DeleteLogCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete log collection params
func (o *DeleteLogCollectionParams) WithRequestBody(requestBody *models.LogCollectionDeletionParams) *DeleteLogCollectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete log collection params
func (o *DeleteLogCollectionParams) SetRequestBody(requestBody *models.LogCollectionDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLogCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
