// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateViewReader is a Reader for the CreateView structure.
type CreateViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateViewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateViewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateViewOK creates a CreateViewOK with default headers values
func NewCreateViewOK() *CreateViewOK {
	return &CreateViewOK{}
}

/* CreateViewOK describes a response with status code 200, with default header values.

CreateViewOK create view o k
*/
type CreateViewOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskView
}

func (o *CreateViewOK) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewOK  %+v", 200, o.Payload)
}
func (o *CreateViewOK) GetPayload() []*models.WithTaskView {
	return o.Payload
}

func (o *CreateViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateViewBadRequest creates a CreateViewBadRequest with default headers values
func NewCreateViewBadRequest() *CreateViewBadRequest {
	return &CreateViewBadRequest{}
}

/* CreateViewBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateViewBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateViewBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewBadRequest  %+v", 400, o.Payload)
}
func (o *CreateViewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateViewNotFound creates a CreateViewNotFound with default headers values
func NewCreateViewNotFound() *CreateViewNotFound {
	return &CreateViewNotFound{}
}

/* CreateViewNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateViewNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateViewNotFound) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewNotFound  %+v", 404, o.Payload)
}
func (o *CreateViewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateViewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateViewInternalServerError creates a CreateViewInternalServerError with default headers values
func NewCreateViewInternalServerError() *CreateViewInternalServerError {
	return &CreateViewInternalServerError{}
}

/* CreateViewInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateViewInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateViewInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateViewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateViewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
