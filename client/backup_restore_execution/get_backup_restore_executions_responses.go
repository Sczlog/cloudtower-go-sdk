// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetBackupRestoreExecutionsReader is a Reader for the GetBackupRestoreExecutions structure.
type GetBackupRestoreExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupRestoreExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupRestoreExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupRestoreExecutionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBackupRestoreExecutionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupRestoreExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupRestoreExecutionsOK creates a GetBackupRestoreExecutionsOK with default headers values
func NewGetBackupRestoreExecutionsOK() *GetBackupRestoreExecutionsOK {
	return &GetBackupRestoreExecutionsOK{}
}

/* GetBackupRestoreExecutionsOK describes a response with status code 200, with default header values.

GetBackupRestoreExecutionsOK get backup restore executions o k
*/
type GetBackupRestoreExecutionsOK struct {
	XTowerRequestID string

	Payload []*models.BackupRestoreExecution
}

func (o *GetBackupRestoreExecutionsOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions][%d] getBackupRestoreExecutionsOK  %+v", 200, o.Payload)
}
func (o *GetBackupRestoreExecutionsOK) GetPayload() []*models.BackupRestoreExecution {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestoreExecutionsBadRequest creates a GetBackupRestoreExecutionsBadRequest with default headers values
func NewGetBackupRestoreExecutionsBadRequest() *GetBackupRestoreExecutionsBadRequest {
	return &GetBackupRestoreExecutionsBadRequest{}
}

/* GetBackupRestoreExecutionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBackupRestoreExecutionsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestoreExecutionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions][%d] getBackupRestoreExecutionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupRestoreExecutionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestoreExecutionsNotFound creates a GetBackupRestoreExecutionsNotFound with default headers values
func NewGetBackupRestoreExecutionsNotFound() *GetBackupRestoreExecutionsNotFound {
	return &GetBackupRestoreExecutionsNotFound{}
}

/* GetBackupRestoreExecutionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBackupRestoreExecutionsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestoreExecutionsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions][%d] getBackupRestoreExecutionsNotFound  %+v", 404, o.Payload)
}
func (o *GetBackupRestoreExecutionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestoreExecutionsInternalServerError creates a GetBackupRestoreExecutionsInternalServerError with default headers values
func NewGetBackupRestoreExecutionsInternalServerError() *GetBackupRestoreExecutionsInternalServerError {
	return &GetBackupRestoreExecutionsInternalServerError{}
}

/* GetBackupRestoreExecutionsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetBackupRestoreExecutionsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestoreExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions][%d] getBackupRestoreExecutionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBackupRestoreExecutionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
