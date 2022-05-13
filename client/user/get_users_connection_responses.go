// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetUsersConnectionReader is a Reader for the GetUsersConnection structure.
type GetUsersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersConnectionOK creates a GetUsersConnectionOK with default headers values
func NewGetUsersConnectionOK() *GetUsersConnectionOK {
	return &GetUsersConnectionOK{}
}

/* GetUsersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetUsersConnectionOK struct {
	Payload *models.UserConnection
}

func (o *GetUsersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-users-connection][%d] getUsersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetUsersConnectionOK) GetPayload() *models.UserConnection {
	return o.Payload
}

func (o *GetUsersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersConnectionBadRequest creates a GetUsersConnectionBadRequest with default headers values
func NewGetUsersConnectionBadRequest() *GetUsersConnectionBadRequest {
	return &GetUsersConnectionBadRequest{}
}

/* GetUsersConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUsersConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUsersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-users-connection][%d] getUsersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetUsersConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersConnectionNotFound creates a GetUsersConnectionNotFound with default headers values
func NewGetUsersConnectionNotFound() *GetUsersConnectionNotFound {
	return &GetUsersConnectionNotFound{}
}

/* GetUsersConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetUsersConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUsersConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-users-connection][%d] getUsersConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetUsersConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersConnectionInternalServerError creates a GetUsersConnectionInternalServerError with default headers values
func NewGetUsersConnectionInternalServerError() *GetUsersConnectionInternalServerError {
	return &GetUsersConnectionInternalServerError{}
}

/* GetUsersConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetUsersConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUsersConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-users-connection][%d] getUsersConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUsersConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
