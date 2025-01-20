// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetBackupRestorePointsReader is a Reader for the GetBackupRestorePoints structure.
type GetBackupRestorePointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupRestorePointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupRestorePointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupRestorePointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBackupRestorePointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupRestorePointsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupRestorePointsOK creates a GetBackupRestorePointsOK with default headers values
func NewGetBackupRestorePointsOK() *GetBackupRestorePointsOK {
	return &GetBackupRestorePointsOK{}
}

/* GetBackupRestorePointsOK describes a response with status code 200, with default header values.

GetBackupRestorePointsOK get backup restore points o k
*/
type GetBackupRestorePointsOK struct {
	XTowerRequestID string

	Payload []*models.BackupRestorePoint
}

func (o *GetBackupRestorePointsOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-points][%d] getBackupRestorePointsOK  %+v", 200, o.Payload)
}
func (o *GetBackupRestorePointsOK) GetPayload() []*models.BackupRestorePoint {
	return o.Payload
}

func (o *GetBackupRestorePointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestorePointsBadRequest creates a GetBackupRestorePointsBadRequest with default headers values
func NewGetBackupRestorePointsBadRequest() *GetBackupRestorePointsBadRequest {
	return &GetBackupRestorePointsBadRequest{}
}

/* GetBackupRestorePointsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBackupRestorePointsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-points][%d] getBackupRestorePointsBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupRestorePointsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestorePointsNotFound creates a GetBackupRestorePointsNotFound with default headers values
func NewGetBackupRestorePointsNotFound() *GetBackupRestorePointsNotFound {
	return &GetBackupRestorePointsNotFound{}
}

/* GetBackupRestorePointsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBackupRestorePointsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-points][%d] getBackupRestorePointsNotFound  %+v", 404, o.Payload)
}
func (o *GetBackupRestorePointsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestorePointsInternalServerError creates a GetBackupRestorePointsInternalServerError with default headers values
func NewGetBackupRestorePointsInternalServerError() *GetBackupRestorePointsInternalServerError {
	return &GetBackupRestorePointsInternalServerError{}
}

/* GetBackupRestorePointsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetBackupRestorePointsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-points][%d] getBackupRestorePointsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBackupRestorePointsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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