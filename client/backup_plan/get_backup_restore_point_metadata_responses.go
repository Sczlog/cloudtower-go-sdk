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

// GetBackupRestorePointMetadataReader is a Reader for the GetBackupRestorePointMetadata structure.
type GetBackupRestorePointMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupRestorePointMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupRestorePointMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupRestorePointMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBackupRestorePointMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBackupRestorePointMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupRestorePointMetadataOK creates a GetBackupRestorePointMetadataOK with default headers values
func NewGetBackupRestorePointMetadataOK() *GetBackupRestorePointMetadataOK {
	return &GetBackupRestorePointMetadataOK{}
}

/* GetBackupRestorePointMetadataOK describes a response with status code 200, with default header values.

GetBackupRestorePointMetadataOK get backup restore point metadata o k
*/
type GetBackupRestorePointMetadataOK struct {
	XTowerRequestID string

	Payload *models.VMMetaData
}

func (o *GetBackupRestorePointMetadataOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-point-metadata][%d] getBackupRestorePointMetadataOK  %+v", 200, o.Payload)
}
func (o *GetBackupRestorePointMetadataOK) GetPayload() *models.VMMetaData {
	return o.Payload
}

func (o *GetBackupRestorePointMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.VMMetaData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupRestorePointMetadataBadRequest creates a GetBackupRestorePointMetadataBadRequest with default headers values
func NewGetBackupRestorePointMetadataBadRequest() *GetBackupRestorePointMetadataBadRequest {
	return &GetBackupRestorePointMetadataBadRequest{}
}

/* GetBackupRestorePointMetadataBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBackupRestorePointMetadataBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointMetadataBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-point-metadata][%d] getBackupRestorePointMetadataBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupRestorePointMetadataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestorePointMetadataNotFound creates a GetBackupRestorePointMetadataNotFound with default headers values
func NewGetBackupRestorePointMetadataNotFound() *GetBackupRestorePointMetadataNotFound {
	return &GetBackupRestorePointMetadataNotFound{}
}

/* GetBackupRestorePointMetadataNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBackupRestorePointMetadataNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointMetadataNotFound) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-point-metadata][%d] getBackupRestorePointMetadataNotFound  %+v", 404, o.Payload)
}
func (o *GetBackupRestorePointMetadataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupRestorePointMetadataInternalServerError creates a GetBackupRestorePointMetadataInternalServerError with default headers values
func NewGetBackupRestorePointMetadataInternalServerError() *GetBackupRestorePointMetadataInternalServerError {
	return &GetBackupRestorePointMetadataInternalServerError{}
}

/* GetBackupRestorePointMetadataInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetBackupRestorePointMetadataInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupRestorePointMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-point-metadata][%d] getBackupRestorePointMetadataInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBackupRestorePointMetadataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupRestorePointMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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