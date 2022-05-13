// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateRootUserReader is a Reader for the CreateRootUser structure.
type CreateRootUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRootUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRootUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRootUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRootUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRootUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRootUserOK creates a CreateRootUserOK with default headers values
func NewCreateRootUserOK() *CreateRootUserOK {
	return &CreateRootUserOK{}
}

/* CreateRootUserOK describes a response with status code 200, with default header values.

Ok
*/
type CreateRootUserOK struct {
	Payload *models.WithTaskUser
}

func (o *CreateRootUserOK) Error() string {
	return fmt.Sprintf("[POST /create-root-user][%d] createRootUserOK  %+v", 200, o.Payload)
}
func (o *CreateRootUserOK) GetPayload() *models.WithTaskUser {
	return o.Payload
}

func (o *CreateRootUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WithTaskUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootUserBadRequest creates a CreateRootUserBadRequest with default headers values
func NewCreateRootUserBadRequest() *CreateRootUserBadRequest {
	return &CreateRootUserBadRequest{}
}

/* CreateRootUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRootUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *CreateRootUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-root-user][%d] createRootUserBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRootUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRootUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootUserNotFound creates a CreateRootUserNotFound with default headers values
func NewCreateRootUserNotFound() *CreateRootUserNotFound {
	return &CreateRootUserNotFound{}
}

/* CreateRootUserNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateRootUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *CreateRootUserNotFound) Error() string {
	return fmt.Sprintf("[POST /create-root-user][%d] createRootUserNotFound  %+v", 404, o.Payload)
}
func (o *CreateRootUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRootUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRootUserInternalServerError creates a CreateRootUserInternalServerError with default headers values
func NewCreateRootUserInternalServerError() *CreateRootUserInternalServerError {
	return &CreateRootUserInternalServerError{}
}

/* CreateRootUserInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateRootUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *CreateRootUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-root-user][%d] createRootUserInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRootUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRootUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
