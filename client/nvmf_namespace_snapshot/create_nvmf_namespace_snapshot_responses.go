// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// CreateNvmfNamespaceSnapshotReader is a Reader for the CreateNvmfNamespaceSnapshot structure.
type CreateNvmfNamespaceSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNvmfNamespaceSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNvmfNamespaceSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNvmfNamespaceSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNvmfNamespaceSnapshotOK creates a CreateNvmfNamespaceSnapshotOK with default headers values
func NewCreateNvmfNamespaceSnapshotOK() *CreateNvmfNamespaceSnapshotOK {
	return &CreateNvmfNamespaceSnapshotOK{}
}

/* CreateNvmfNamespaceSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CreateNvmfNamespaceSnapshotOK struct {
	Payload []*models.WithTaskNvmfNamespaceSnapshot
}

func (o *CreateNvmfNamespaceSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-namespace-snapshot][%d] createNvmfNamespaceSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateNvmfNamespaceSnapshotOK) GetPayload() []*models.WithTaskNvmfNamespaceSnapshot {
	return o.Payload
}

func (o *CreateNvmfNamespaceSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNvmfNamespaceSnapshotBadRequest creates a CreateNvmfNamespaceSnapshotBadRequest with default headers values
func NewCreateNvmfNamespaceSnapshotBadRequest() *CreateNvmfNamespaceSnapshotBadRequest {
	return &CreateNvmfNamespaceSnapshotBadRequest{}
}

/* CreateNvmfNamespaceSnapshotBadRequest describes a response with status code 400, with default header values.

CreateNvmfNamespaceSnapshotBadRequest create nvmf namespace snapshot bad request
*/
type CreateNvmfNamespaceSnapshotBadRequest struct {
	Payload string
}

func (o *CreateNvmfNamespaceSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-namespace-snapshot][%d] createNvmfNamespaceSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateNvmfNamespaceSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateNvmfNamespaceSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
