// Code generated by go-swagger; DO NOT EDIT.

package upload_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetUploadTasksConnectionReader is a Reader for the GetUploadTasksConnection structure.
type GetUploadTasksConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUploadTasksConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUploadTasksConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUploadTasksConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUploadTasksConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUploadTasksConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUploadTasksConnectionOK creates a GetUploadTasksConnectionOK with default headers values
func NewGetUploadTasksConnectionOK() *GetUploadTasksConnectionOK {
	return &GetUploadTasksConnectionOK{}
}

/* GetUploadTasksConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetUploadTasksConnectionOK struct {
	Payload *models.UploadTaskConnection
}

func (o *GetUploadTasksConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-upload-tasks-connection][%d] getUploadTasksConnectionOK  %+v", 200, o.Payload)
}
func (o *GetUploadTasksConnectionOK) GetPayload() *models.UploadTaskConnection {
	return o.Payload
}

func (o *GetUploadTasksConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadTaskConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadTasksConnectionBadRequest creates a GetUploadTasksConnectionBadRequest with default headers values
func NewGetUploadTasksConnectionBadRequest() *GetUploadTasksConnectionBadRequest {
	return &GetUploadTasksConnectionBadRequest{}
}

/* GetUploadTasksConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUploadTasksConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUploadTasksConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-upload-tasks-connection][%d] getUploadTasksConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetUploadTasksConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUploadTasksConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadTasksConnectionNotFound creates a GetUploadTasksConnectionNotFound with default headers values
func NewGetUploadTasksConnectionNotFound() *GetUploadTasksConnectionNotFound {
	return &GetUploadTasksConnectionNotFound{}
}

/* GetUploadTasksConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetUploadTasksConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUploadTasksConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-upload-tasks-connection][%d] getUploadTasksConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetUploadTasksConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUploadTasksConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadTasksConnectionInternalServerError creates a GetUploadTasksConnectionInternalServerError with default headers values
func NewGetUploadTasksConnectionInternalServerError() *GetUploadTasksConnectionInternalServerError {
	return &GetUploadTasksConnectionInternalServerError{}
}

/* GetUploadTasksConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetUploadTasksConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUploadTasksConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-upload-tasks-connection][%d] getUploadTasksConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUploadTasksConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUploadTasksConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
