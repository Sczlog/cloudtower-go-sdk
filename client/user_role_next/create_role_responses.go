// Code generated by go-swagger; DO NOT EDIT.

package user_role_next

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateRoleReader is a Reader for the CreateRole structure.
type CreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRoleOK creates a CreateRoleOK with default headers values
func NewCreateRoleOK() *CreateRoleOK {
	return &CreateRoleOK{}
}

/* CreateRoleOK describes a response with status code 200, with default header values.

CreateRoleOK create role o k
*/
type CreateRoleOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskUserRoleNext
}

func (o *CreateRoleOK) Error() string {
	return fmt.Sprintf("[POST /create-role][%d] createRoleOK  %+v", 200, o.Payload)
}
func (o *CreateRoleOK) GetPayload() []*models.WithTaskUserRoleNext {
	return o.Payload
}

func (o *CreateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRoleBadRequest creates a CreateRoleBadRequest with default headers values
func NewCreateRoleBadRequest() *CreateRoleBadRequest {
	return &CreateRoleBadRequest{}
}

/* CreateRoleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRoleBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-role][%d] createRoleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRoleNotFound creates a CreateRoleNotFound with default headers values
func NewCreateRoleNotFound() *CreateRoleNotFound {
	return &CreateRoleNotFound{}
}

/* CreateRoleNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateRoleNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /create-role][%d] createRoleNotFound  %+v", 404, o.Payload)
}
func (o *CreateRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRoleInternalServerError creates a CreateRoleInternalServerError with default headers values
func NewCreateRoleInternalServerError() *CreateRoleInternalServerError {
	return &CreateRoleInternalServerError{}
}

/* CreateRoleInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateRoleInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-role][%d] createRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
