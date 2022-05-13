// Code generated by go-swagger; DO NOT EDIT.

package snmp_trap_receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteSnmpTrapReceiverReader is a Reader for the DeleteSnmpTrapReceiver structure.
type DeleteSnmpTrapReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnmpTrapReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSnmpTrapReceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSnmpTrapReceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSnmpTrapReceiverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSnmpTrapReceiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSnmpTrapReceiverOK creates a DeleteSnmpTrapReceiverOK with default headers values
func NewDeleteSnmpTrapReceiverOK() *DeleteSnmpTrapReceiverOK {
	return &DeleteSnmpTrapReceiverOK{}
}

/* DeleteSnmpTrapReceiverOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteSnmpTrapReceiverOK struct {
	Payload []*models.WithTaskDeleteSnmpTrapReceiver
}

func (o *DeleteSnmpTrapReceiverOK) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-trap-receiver][%d] deleteSnmpTrapReceiverOK  %+v", 200, o.Payload)
}
func (o *DeleteSnmpTrapReceiverOK) GetPayload() []*models.WithTaskDeleteSnmpTrapReceiver {
	return o.Payload
}

func (o *DeleteSnmpTrapReceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnmpTrapReceiverBadRequest creates a DeleteSnmpTrapReceiverBadRequest with default headers values
func NewDeleteSnmpTrapReceiverBadRequest() *DeleteSnmpTrapReceiverBadRequest {
	return &DeleteSnmpTrapReceiverBadRequest{}
}

/* DeleteSnmpTrapReceiverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteSnmpTrapReceiverBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteSnmpTrapReceiverBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-trap-receiver][%d] deleteSnmpTrapReceiverBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteSnmpTrapReceiverBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSnmpTrapReceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnmpTrapReceiverNotFound creates a DeleteSnmpTrapReceiverNotFound with default headers values
func NewDeleteSnmpTrapReceiverNotFound() *DeleteSnmpTrapReceiverNotFound {
	return &DeleteSnmpTrapReceiverNotFound{}
}

/* DeleteSnmpTrapReceiverNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteSnmpTrapReceiverNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteSnmpTrapReceiverNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-trap-receiver][%d] deleteSnmpTrapReceiverNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSnmpTrapReceiverNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSnmpTrapReceiverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnmpTrapReceiverInternalServerError creates a DeleteSnmpTrapReceiverInternalServerError with default headers values
func NewDeleteSnmpTrapReceiverInternalServerError() *DeleteSnmpTrapReceiverInternalServerError {
	return &DeleteSnmpTrapReceiverInternalServerError{}
}

/* DeleteSnmpTrapReceiverInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteSnmpTrapReceiverInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteSnmpTrapReceiverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-trap-receiver][%d] deleteSnmpTrapReceiverInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteSnmpTrapReceiverInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSnmpTrapReceiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
