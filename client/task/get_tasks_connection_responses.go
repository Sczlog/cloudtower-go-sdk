// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetTasksConnectionReader is a Reader for the GetTasksConnection structure.
type GetTasksConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTasksConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTasksConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTasksConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksConnectionOK creates a GetTasksConnectionOK with default headers values
func NewGetTasksConnectionOK() *GetTasksConnectionOK {
	return &GetTasksConnectionOK{}
}

/* GetTasksConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetTasksConnectionOK struct {
	Payload *models.TaskConnection
}

func (o *GetTasksConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-tasks-connection][%d] getTasksConnectionOK  %+v", 200, o.Payload)
}
func (o *GetTasksConnectionOK) GetPayload() *models.TaskConnection {
	return o.Payload
}

func (o *GetTasksConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksConnectionBadRequest creates a GetTasksConnectionBadRequest with default headers values
func NewGetTasksConnectionBadRequest() *GetTasksConnectionBadRequest {
	return &GetTasksConnectionBadRequest{}
}

/* GetTasksConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetTasksConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTasksConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-tasks-connection][%d] getTasksConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetTasksConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTasksConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksConnectionNotFound creates a GetTasksConnectionNotFound with default headers values
func NewGetTasksConnectionNotFound() *GetTasksConnectionNotFound {
	return &GetTasksConnectionNotFound{}
}

/* GetTasksConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTasksConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTasksConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-tasks-connection][%d] getTasksConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetTasksConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTasksConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksConnectionInternalServerError creates a GetTasksConnectionInternalServerError with default headers values
func NewGetTasksConnectionInternalServerError() *GetTasksConnectionInternalServerError {
	return &GetTasksConnectionInternalServerError{}
}

/* GetTasksConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetTasksConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTasksConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-tasks-connection][%d] getTasksConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTasksConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTasksConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
