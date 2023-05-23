// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateSecurityGroupReader is a Reader for the CreateSecurityGroup structure.
type CreateSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSecurityGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSecurityGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSecurityGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSecurityGroupOK creates a CreateSecurityGroupOK with default headers values
func NewCreateSecurityGroupOK() *CreateSecurityGroupOK {
	return &CreateSecurityGroupOK{}
}

/* CreateSecurityGroupOK describes a response with status code 200, with default header values.

CreateSecurityGroupOK create security group o k
*/
type CreateSecurityGroupOK struct {
	XTowerRequestID string

	Payload *models.WithTaskSecurityGroup
}

func (o *CreateSecurityGroupOK) Error() string {
	return fmt.Sprintf("[POST /create-security-group][%d] createSecurityGroupOK  %+v", 200, o.Payload)
}
func (o *CreateSecurityGroupOK) GetPayload() *models.WithTaskSecurityGroup {
	return o.Payload
}

func (o *CreateSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.WithTaskSecurityGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityGroupBadRequest creates a CreateSecurityGroupBadRequest with default headers values
func NewCreateSecurityGroupBadRequest() *CreateSecurityGroupBadRequest {
	return &CreateSecurityGroupBadRequest{}
}

/* CreateSecurityGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSecurityGroupBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-security-group][%d] createSecurityGroupBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSecurityGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSecurityGroupNotFound creates a CreateSecurityGroupNotFound with default headers values
func NewCreateSecurityGroupNotFound() *CreateSecurityGroupNotFound {
	return &CreateSecurityGroupNotFound{}
}

/* CreateSecurityGroupNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateSecurityGroupNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /create-security-group][%d] createSecurityGroupNotFound  %+v", 404, o.Payload)
}
func (o *CreateSecurityGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSecurityGroupInternalServerError creates a CreateSecurityGroupInternalServerError with default headers values
func NewCreateSecurityGroupInternalServerError() *CreateSecurityGroupInternalServerError {
	return &CreateSecurityGroupInternalServerError{}
}

/* CreateSecurityGroupInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateSecurityGroupInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-security-group][%d] createSecurityGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateSecurityGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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