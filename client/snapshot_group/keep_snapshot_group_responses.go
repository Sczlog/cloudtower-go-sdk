// Code generated by go-swagger; DO NOT EDIT.

package snapshot_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// KeepSnapshotGroupReader is a Reader for the KeepSnapshotGroup structure.
type KeepSnapshotGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeepSnapshotGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeepSnapshotGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeepSnapshotGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeepSnapshotGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKeepSnapshotGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeepSnapshotGroupOK creates a KeepSnapshotGroupOK with default headers values
func NewKeepSnapshotGroupOK() *KeepSnapshotGroupOK {
	return &KeepSnapshotGroupOK{}
}

/* KeepSnapshotGroupOK describes a response with status code 200, with default header values.

Ok
*/
type KeepSnapshotGroupOK struct {
	Payload []*models.WithTaskSnapshotGroup
}

func (o *KeepSnapshotGroupOK) Error() string {
	return fmt.Sprintf("[POST /keep-snapshot-group][%d] keepSnapshotGroupOK  %+v", 200, o.Payload)
}
func (o *KeepSnapshotGroupOK) GetPayload() []*models.WithTaskSnapshotGroup {
	return o.Payload
}

func (o *KeepSnapshotGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeepSnapshotGroupBadRequest creates a KeepSnapshotGroupBadRequest with default headers values
func NewKeepSnapshotGroupBadRequest() *KeepSnapshotGroupBadRequest {
	return &KeepSnapshotGroupBadRequest{}
}

/* KeepSnapshotGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type KeepSnapshotGroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *KeepSnapshotGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /keep-snapshot-group][%d] keepSnapshotGroupBadRequest  %+v", 400, o.Payload)
}
func (o *KeepSnapshotGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *KeepSnapshotGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeepSnapshotGroupNotFound creates a KeepSnapshotGroupNotFound with default headers values
func NewKeepSnapshotGroupNotFound() *KeepSnapshotGroupNotFound {
	return &KeepSnapshotGroupNotFound{}
}

/* KeepSnapshotGroupNotFound describes a response with status code 404, with default header values.

Not found
*/
type KeepSnapshotGroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *KeepSnapshotGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /keep-snapshot-group][%d] keepSnapshotGroupNotFound  %+v", 404, o.Payload)
}
func (o *KeepSnapshotGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *KeepSnapshotGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeepSnapshotGroupInternalServerError creates a KeepSnapshotGroupInternalServerError with default headers values
func NewKeepSnapshotGroupInternalServerError() *KeepSnapshotGroupInternalServerError {
	return &KeepSnapshotGroupInternalServerError{}
}

/* KeepSnapshotGroupInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type KeepSnapshotGroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *KeepSnapshotGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /keep-snapshot-group][%d] keepSnapshotGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *KeepSnapshotGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *KeepSnapshotGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
