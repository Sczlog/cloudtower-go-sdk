// Code generated by go-swagger; DO NOT EDIT.

package backup_target_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetBackupTargetExecutionsConnectionReader is a Reader for the GetBackupTargetExecutionsConnection structure.
type GetBackupTargetExecutionsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupTargetExecutionsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupTargetExecutionsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupTargetExecutionsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupTargetExecutionsConnectionOK creates a GetBackupTargetExecutionsConnectionOK with default headers values
func NewGetBackupTargetExecutionsConnectionOK() *GetBackupTargetExecutionsConnectionOK {
	return &GetBackupTargetExecutionsConnectionOK{}
}

/* GetBackupTargetExecutionsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupTargetExecutionsConnectionOK struct {
	Payload *models.BackupTargetExecutionConnection
}

func (o *GetBackupTargetExecutionsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-target-executions-connection][%d] getBackupTargetExecutionsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBackupTargetExecutionsConnectionOK) GetPayload() *models.BackupTargetExecutionConnection {
	return o.Payload
}

func (o *GetBackupTargetExecutionsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupTargetExecutionConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupTargetExecutionsConnectionBadRequest creates a GetBackupTargetExecutionsConnectionBadRequest with default headers values
func NewGetBackupTargetExecutionsConnectionBadRequest() *GetBackupTargetExecutionsConnectionBadRequest {
	return &GetBackupTargetExecutionsConnectionBadRequest{}
}

/* GetBackupTargetExecutionsConnectionBadRequest describes a response with status code 400, with default header values.

GetBackupTargetExecutionsConnectionBadRequest get backup target executions connection bad request
*/
type GetBackupTargetExecutionsConnectionBadRequest struct {
	Payload string
}

func (o *GetBackupTargetExecutionsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-target-executions-connection][%d] getBackupTargetExecutionsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupTargetExecutionsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBackupTargetExecutionsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
