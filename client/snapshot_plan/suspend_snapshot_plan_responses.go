// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// SuspendSnapshotPlanReader is a Reader for the SuspendSnapshotPlan structure.
type SuspendSnapshotPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendSnapshotPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendSnapshotPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuspendSnapshotPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuspendSnapshotPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSuspendSnapshotPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuspendSnapshotPlanOK creates a SuspendSnapshotPlanOK with default headers values
func NewSuspendSnapshotPlanOK() *SuspendSnapshotPlanOK {
	return &SuspendSnapshotPlanOK{}
}

/* SuspendSnapshotPlanOK describes a response with status code 200, with default header values.

SuspendSnapshotPlanOK suspend snapshot plan o k
*/
type SuspendSnapshotPlanOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskSnapshotPlan
}

func (o *SuspendSnapshotPlanOK) Error() string {
	return fmt.Sprintf("[POST /suspend-snapshot-plan][%d] suspendSnapshotPlanOK  %+v", 200, o.Payload)
}
func (o *SuspendSnapshotPlanOK) GetPayload() []*models.WithTaskSnapshotPlan {
	return o.Payload
}

func (o *SuspendSnapshotPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendSnapshotPlanBadRequest creates a SuspendSnapshotPlanBadRequest with default headers values
func NewSuspendSnapshotPlanBadRequest() *SuspendSnapshotPlanBadRequest {
	return &SuspendSnapshotPlanBadRequest{}
}

/* SuspendSnapshotPlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SuspendSnapshotPlanBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendSnapshotPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /suspend-snapshot-plan][%d] suspendSnapshotPlanBadRequest  %+v", 400, o.Payload)
}
func (o *SuspendSnapshotPlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendSnapshotPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendSnapshotPlanNotFound creates a SuspendSnapshotPlanNotFound with default headers values
func NewSuspendSnapshotPlanNotFound() *SuspendSnapshotPlanNotFound {
	return &SuspendSnapshotPlanNotFound{}
}

/* SuspendSnapshotPlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type SuspendSnapshotPlanNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendSnapshotPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /suspend-snapshot-plan][%d] suspendSnapshotPlanNotFound  %+v", 404, o.Payload)
}
func (o *SuspendSnapshotPlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendSnapshotPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendSnapshotPlanInternalServerError creates a SuspendSnapshotPlanInternalServerError with default headers values
func NewSuspendSnapshotPlanInternalServerError() *SuspendSnapshotPlanInternalServerError {
	return &SuspendSnapshotPlanInternalServerError{}
}

/* SuspendSnapshotPlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SuspendSnapshotPlanInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendSnapshotPlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /suspend-snapshot-plan][%d] suspendSnapshotPlanInternalServerError  %+v", 500, o.Payload)
}
func (o *SuspendSnapshotPlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendSnapshotPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
