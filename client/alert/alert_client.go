// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alert API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alert API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAlerts(params *GetAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsOK, error)

	GetAlertsConnection(params *GetAlertsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsConnectionOK, error)

	ResolveAlert(params *ResolveAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResolveAlertOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAlerts get alerts API
*/
func (a *Client) GetAlerts(params *GetAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAlerts",
		Method:             "POST",
		PathPattern:        "/get-alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAlertsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlertsConnection get alerts connection API
*/
func (a *Client) GetAlertsConnection(params *GetAlertsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAlertsConnection",
		Method:             "POST",
		PathPattern:        "/get-alerts-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAlertsConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAlertsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAlertsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResolveAlert resolve alert API
*/
func (a *Client) ResolveAlert(params *ResolveAlertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResolveAlertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResolveAlertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResolveAlert",
		Method:             "POST",
		PathPattern:        "/resolve-alert",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResolveAlertReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResolveAlertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResolveAlert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
