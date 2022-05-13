// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateIscsiLunSnapshotReader is a Reader for the CreateIscsiLunSnapshot structure.
type CreateIscsiLunSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIscsiLunSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIscsiLunSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIscsiLunSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateIscsiLunSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateIscsiLunSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateIscsiLunSnapshotOK creates a CreateIscsiLunSnapshotOK with default headers values
func NewCreateIscsiLunSnapshotOK() *CreateIscsiLunSnapshotOK {
	return &CreateIscsiLunSnapshotOK{}
}

/* CreateIscsiLunSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CreateIscsiLunSnapshotOK struct {
	Payload []*models.WithTaskIscsiLunSnapshot
}

func (o *CreateIscsiLunSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateIscsiLunSnapshotOK) GetPayload() []*models.WithTaskIscsiLunSnapshot {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIscsiLunSnapshotBadRequest creates a CreateIscsiLunSnapshotBadRequest with default headers values
func NewCreateIscsiLunSnapshotBadRequest() *CreateIscsiLunSnapshotBadRequest {
	return &CreateIscsiLunSnapshotBadRequest{}
}

/* CreateIscsiLunSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateIscsiLunSnapshotBadRequest struct {
	Payload *models.ErrorBody
}

func (o *CreateIscsiLunSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateIscsiLunSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIscsiLunSnapshotNotFound creates a CreateIscsiLunSnapshotNotFound with default headers values
func NewCreateIscsiLunSnapshotNotFound() *CreateIscsiLunSnapshotNotFound {
	return &CreateIscsiLunSnapshotNotFound{}
}

/* CreateIscsiLunSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateIscsiLunSnapshotNotFound struct {
	Payload *models.ErrorBody
}

func (o *CreateIscsiLunSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *CreateIscsiLunSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIscsiLunSnapshotInternalServerError creates a CreateIscsiLunSnapshotInternalServerError with default headers values
func NewCreateIscsiLunSnapshotInternalServerError() *CreateIscsiLunSnapshotInternalServerError {
	return &CreateIscsiLunSnapshotInternalServerError{}
}

/* CreateIscsiLunSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateIscsiLunSnapshotInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *CreateIscsiLunSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateIscsiLunSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
