// Code generated by go-swagger; DO NOT EDIT.

package consistency_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new consistency group snapshot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for consistency group snapshot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateConsistencyGroupSnapshot(params *CreateConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConsistencyGroupSnapshotOK, error)

	DeleteConsistencyGroupSnapshot(params *DeleteConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConsistencyGroupSnapshotOK, error)

	GetConsistencyGroupSnapshots(params *GetConsistencyGroupSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsistencyGroupSnapshotsOK, error)

	GetConsistencyGroupSnapshotsConnection(params *GetConsistencyGroupSnapshotsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsistencyGroupSnapshotsConnectionOK, error)

	UpdateConsistencyGroupSnapshot(params *UpdateConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConsistencyGroupSnapshotOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateConsistencyGroupSnapshot create consistency group snapshot API
*/
func (a *Client) CreateConsistencyGroupSnapshot(params *CreateConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateConsistencyGroupSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConsistencyGroupSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateConsistencyGroupSnapshot",
		Method:             "POST",
		PathPattern:        "/create-consistency-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateConsistencyGroupSnapshotReader{formats: a.formats},
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
	success, ok := result.(*CreateConsistencyGroupSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateConsistencyGroupSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteConsistencyGroupSnapshot delete consistency group snapshot API
*/
func (a *Client) DeleteConsistencyGroupSnapshot(params *DeleteConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteConsistencyGroupSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConsistencyGroupSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteConsistencyGroupSnapshot",
		Method:             "POST",
		PathPattern:        "/delete-consistency-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteConsistencyGroupSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteConsistencyGroupSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteConsistencyGroupSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConsistencyGroupSnapshots get consistency group snapshots API
*/
func (a *Client) GetConsistencyGroupSnapshots(params *GetConsistencyGroupSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsistencyGroupSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsistencyGroupSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConsistencyGroupSnapshots",
		Method:             "POST",
		PathPattern:        "/get-consistency-group-snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConsistencyGroupSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*GetConsistencyGroupSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConsistencyGroupSnapshots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConsistencyGroupSnapshotsConnection get consistency group snapshots connection API
*/
func (a *Client) GetConsistencyGroupSnapshotsConnection(params *GetConsistencyGroupSnapshotsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsistencyGroupSnapshotsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsistencyGroupSnapshotsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConsistencyGroupSnapshotsConnection",
		Method:             "POST",
		PathPattern:        "/get-consistency-group-snapshots-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConsistencyGroupSnapshotsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetConsistencyGroupSnapshotsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConsistencyGroupSnapshotsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateConsistencyGroupSnapshot update consistency group snapshot API
*/
func (a *Client) UpdateConsistencyGroupSnapshot(params *UpdateConsistencyGroupSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateConsistencyGroupSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConsistencyGroupSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateConsistencyGroupSnapshot",
		Method:             "POST",
		PathPattern:        "/rollback-consistency-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateConsistencyGroupSnapshotReader{formats: a.formats},
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
	success, ok := result.(*UpdateConsistencyGroupSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateConsistencyGroupSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
