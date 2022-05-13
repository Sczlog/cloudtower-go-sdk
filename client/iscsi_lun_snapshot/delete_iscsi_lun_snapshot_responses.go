// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteIscsiLunSnapshotReader is a Reader for the DeleteIscsiLunSnapshot structure.
type DeleteIscsiLunSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIscsiLunSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIscsiLunSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIscsiLunSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIscsiLunSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIscsiLunSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIscsiLunSnapshotOK creates a DeleteIscsiLunSnapshotOK with default headers values
func NewDeleteIscsiLunSnapshotOK() *DeleteIscsiLunSnapshotOK {
	return &DeleteIscsiLunSnapshotOK{}
}

/* DeleteIscsiLunSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteIscsiLunSnapshotOK struct {
	Payload []*models.WithTaskDeleteIscsiLunSnapshot
}

func (o *DeleteIscsiLunSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun-snapshot][%d] deleteIscsiLunSnapshotOK  %+v", 200, o.Payload)
}
func (o *DeleteIscsiLunSnapshotOK) GetPayload() []*models.WithTaskDeleteIscsiLunSnapshot {
	return o.Payload
}

func (o *DeleteIscsiLunSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunSnapshotBadRequest creates a DeleteIscsiLunSnapshotBadRequest with default headers values
func NewDeleteIscsiLunSnapshotBadRequest() *DeleteIscsiLunSnapshotBadRequest {
	return &DeleteIscsiLunSnapshotBadRequest{}
}

/* DeleteIscsiLunSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteIscsiLunSnapshotBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun-snapshot][%d] deleteIscsiLunSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteIscsiLunSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunSnapshotNotFound creates a DeleteIscsiLunSnapshotNotFound with default headers values
func NewDeleteIscsiLunSnapshotNotFound() *DeleteIscsiLunSnapshotNotFound {
	return &DeleteIscsiLunSnapshotNotFound{}
}

/* DeleteIscsiLunSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteIscsiLunSnapshotNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun-snapshot][%d] deleteIscsiLunSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *DeleteIscsiLunSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunSnapshotInternalServerError creates a DeleteIscsiLunSnapshotInternalServerError with default headers values
func NewDeleteIscsiLunSnapshotInternalServerError() *DeleteIscsiLunSnapshotInternalServerError {
	return &DeleteIscsiLunSnapshotInternalServerError{}
}

/* DeleteIscsiLunSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteIscsiLunSnapshotInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun-snapshot][%d] deleteIscsiLunSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteIscsiLunSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
