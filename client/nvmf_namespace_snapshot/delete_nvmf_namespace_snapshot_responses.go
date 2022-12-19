// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteNvmfNamespaceSnapshotReader is a Reader for the DeleteNvmfNamespaceSnapshot structure.
type DeleteNvmfNamespaceSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNvmfNamespaceSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNvmfNamespaceSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNvmfNamespaceSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNvmfNamespaceSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNvmfNamespaceSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNvmfNamespaceSnapshotOK creates a DeleteNvmfNamespaceSnapshotOK with default headers values
func NewDeleteNvmfNamespaceSnapshotOK() *DeleteNvmfNamespaceSnapshotOK {
	return &DeleteNvmfNamespaceSnapshotOK{}
}

/* DeleteNvmfNamespaceSnapshotOK describes a response with status code 200, with default header values.

DeleteNvmfNamespaceSnapshotOK delete nvmf namespace snapshot o k
*/
type DeleteNvmfNamespaceSnapshotOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteNvmfNamespaceSnapshot
}

func (o *DeleteNvmfNamespaceSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace-snapshot][%d] deleteNvmfNamespaceSnapshotOK  %+v", 200, o.Payload)
}
func (o *DeleteNvmfNamespaceSnapshotOK) GetPayload() []*models.WithTaskDeleteNvmfNamespaceSnapshot {
	return o.Payload
}

func (o *DeleteNvmfNamespaceSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNvmfNamespaceSnapshotBadRequest creates a DeleteNvmfNamespaceSnapshotBadRequest with default headers values
func NewDeleteNvmfNamespaceSnapshotBadRequest() *DeleteNvmfNamespaceSnapshotBadRequest {
	return &DeleteNvmfNamespaceSnapshotBadRequest{}
}

/* DeleteNvmfNamespaceSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteNvmfNamespaceSnapshotBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace-snapshot][%d] deleteNvmfNamespaceSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteNvmfNamespaceSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNvmfNamespaceSnapshotNotFound creates a DeleteNvmfNamespaceSnapshotNotFound with default headers values
func NewDeleteNvmfNamespaceSnapshotNotFound() *DeleteNvmfNamespaceSnapshotNotFound {
	return &DeleteNvmfNamespaceSnapshotNotFound{}
}

/* DeleteNvmfNamespaceSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteNvmfNamespaceSnapshotNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace-snapshot][%d] deleteNvmfNamespaceSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *DeleteNvmfNamespaceSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNvmfNamespaceSnapshotInternalServerError creates a DeleteNvmfNamespaceSnapshotInternalServerError with default headers values
func NewDeleteNvmfNamespaceSnapshotInternalServerError() *DeleteNvmfNamespaceSnapshotInternalServerError {
	return &DeleteNvmfNamespaceSnapshotInternalServerError{}
}

/* DeleteNvmfNamespaceSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteNvmfNamespaceSnapshotInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace-snapshot][%d] deleteNvmfNamespaceSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteNvmfNamespaceSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
