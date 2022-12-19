// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateHostReader is a Reader for the CreateHost structure.
type CreateHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateHostOK creates a CreateHostOK with default headers values
func NewCreateHostOK() *CreateHostOK {
	return &CreateHostOK{}
}

/* CreateHostOK describes a response with status code 200, with default header values.

CreateHostOK create host o k
*/
type CreateHostOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskBatchHosts
}

func (o *CreateHostOK) Error() string {
	return fmt.Sprintf("[POST /create-host][%d] createHostOK  %+v", 200, o.Payload)
}
func (o *CreateHostOK) GetPayload() []*models.WithTaskBatchHosts {
	return o.Payload
}

func (o *CreateHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHostBadRequest creates a CreateHostBadRequest with default headers values
func NewCreateHostBadRequest() *CreateHostBadRequest {
	return &CreateHostBadRequest{}
}

/* CreateHostBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateHostBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateHostBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-host][%d] createHostBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHostBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHostNotFound creates a CreateHostNotFound with default headers values
func NewCreateHostNotFound() *CreateHostNotFound {
	return &CreateHostNotFound{}
}

/* CreateHostNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateHostNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateHostNotFound) Error() string {
	return fmt.Sprintf("[POST /create-host][%d] createHostNotFound  %+v", 404, o.Payload)
}
func (o *CreateHostNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateHostInternalServerError creates a CreateHostInternalServerError with default headers values
func NewCreateHostInternalServerError() *CreateHostInternalServerError {
	return &CreateHostInternalServerError{}
}

/* CreateHostInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateHostInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateHostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-host][%d] createHostInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateHostInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
