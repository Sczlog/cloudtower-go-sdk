// Code generated by go-swagger; DO NOT EDIT.

package vds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteVdsReader is a Reader for the DeleteVds structure.
type DeleteVdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVdsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVdsOK creates a DeleteVdsOK with default headers values
func NewDeleteVdsOK() *DeleteVdsOK {
	return &DeleteVdsOK{}
}

/* DeleteVdsOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVdsOK struct {
	Payload []*models.WithTaskDeleteVds
}

func (o *DeleteVdsOK) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsOK  %+v", 200, o.Payload)
}
func (o *DeleteVdsOK) GetPayload() []*models.WithTaskDeleteVds {
	return o.Payload
}

func (o *DeleteVdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVdsBadRequest creates a DeleteVdsBadRequest with default headers values
func NewDeleteVdsBadRequest() *DeleteVdsBadRequest {
	return &DeleteVdsBadRequest{}
}

/* DeleteVdsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVdsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteVdsBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVdsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVdsNotFound creates a DeleteVdsNotFound with default headers values
func NewDeleteVdsNotFound() *DeleteVdsNotFound {
	return &DeleteVdsNotFound{}
}

/* DeleteVdsNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVdsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteVdsNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVdsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVdsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVdsInternalServerError creates a DeleteVdsInternalServerError with default headers values
func NewDeleteVdsInternalServerError() *DeleteVdsInternalServerError {
	return &DeleteVdsInternalServerError{}
}

/* DeleteVdsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVdsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteVdsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVdsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
