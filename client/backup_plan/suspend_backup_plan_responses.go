// Code generated by go-swagger; DO NOT EDIT.

package backup_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// SuspendBackupPlanReader is a Reader for the SuspendBackupPlan structure.
type SuspendBackupPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendBackupPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendBackupPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuspendBackupPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuspendBackupPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSuspendBackupPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuspendBackupPlanOK creates a SuspendBackupPlanOK with default headers values
func NewSuspendBackupPlanOK() *SuspendBackupPlanOK {
	return &SuspendBackupPlanOK{}
}

/* SuspendBackupPlanOK describes a response with status code 200, with default header values.

SuspendBackupPlanOK suspend backup plan o k
*/
type SuspendBackupPlanOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskBackupPlan
}

func (o *SuspendBackupPlanOK) Error() string {
	return fmt.Sprintf("[POST /suspend-backup-plan][%d] suspendBackupPlanOK  %+v", 200, o.Payload)
}
func (o *SuspendBackupPlanOK) GetPayload() []*models.WithTaskBackupPlan {
	return o.Payload
}

func (o *SuspendBackupPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendBackupPlanBadRequest creates a SuspendBackupPlanBadRequest with default headers values
func NewSuspendBackupPlanBadRequest() *SuspendBackupPlanBadRequest {
	return &SuspendBackupPlanBadRequest{}
}

/* SuspendBackupPlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SuspendBackupPlanBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendBackupPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /suspend-backup-plan][%d] suspendBackupPlanBadRequest  %+v", 400, o.Payload)
}
func (o *SuspendBackupPlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendBackupPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendBackupPlanNotFound creates a SuspendBackupPlanNotFound with default headers values
func NewSuspendBackupPlanNotFound() *SuspendBackupPlanNotFound {
	return &SuspendBackupPlanNotFound{}
}

/* SuspendBackupPlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type SuspendBackupPlanNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendBackupPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /suspend-backup-plan][%d] suspendBackupPlanNotFound  %+v", 404, o.Payload)
}
func (o *SuspendBackupPlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendBackupPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendBackupPlanInternalServerError creates a SuspendBackupPlanInternalServerError with default headers values
func NewSuspendBackupPlanInternalServerError() *SuspendBackupPlanInternalServerError {
	return &SuspendBackupPlanInternalServerError{}
}

/* SuspendBackupPlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SuspendBackupPlanInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendBackupPlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /suspend-backup-plan][%d] suspendBackupPlanInternalServerError  %+v", 500, o.Payload)
}
func (o *SuspendBackupPlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendBackupPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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