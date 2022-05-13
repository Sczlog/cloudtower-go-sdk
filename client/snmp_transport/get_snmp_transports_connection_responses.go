// Code generated by go-swagger; DO NOT EDIT.

package snmp_transport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetSnmpTransportsConnectionReader is a Reader for the GetSnmpTransportsConnection structure.
type GetSnmpTransportsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnmpTransportsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnmpTransportsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnmpTransportsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnmpTransportsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSnmpTransportsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnmpTransportsConnectionOK creates a GetSnmpTransportsConnectionOK with default headers values
func NewGetSnmpTransportsConnectionOK() *GetSnmpTransportsConnectionOK {
	return &GetSnmpTransportsConnectionOK{}
}

/* GetSnmpTransportsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnmpTransportsConnectionOK struct {
	Payload *models.SnmpTransportConnection
}

func (o *GetSnmpTransportsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports-connection][%d] getSnmpTransportsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetSnmpTransportsConnectionOK) GetPayload() *models.SnmpTransportConnection {
	return o.Payload
}

func (o *GetSnmpTransportsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpTransportConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsConnectionBadRequest creates a GetSnmpTransportsConnectionBadRequest with default headers values
func NewGetSnmpTransportsConnectionBadRequest() *GetSnmpTransportsConnectionBadRequest {
	return &GetSnmpTransportsConnectionBadRequest{}
}

/* GetSnmpTransportsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSnmpTransportsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports-connection][%d] getSnmpTransportsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnmpTransportsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsConnectionNotFound creates a GetSnmpTransportsConnectionNotFound with default headers values
func NewGetSnmpTransportsConnectionNotFound() *GetSnmpTransportsConnectionNotFound {
	return &GetSnmpTransportsConnectionNotFound{}
}

/* GetSnmpTransportsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSnmpTransportsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports-connection][%d] getSnmpTransportsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetSnmpTransportsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsConnectionInternalServerError creates a GetSnmpTransportsConnectionInternalServerError with default headers values
func NewGetSnmpTransportsConnectionInternalServerError() *GetSnmpTransportsConnectionInternalServerError {
	return &GetSnmpTransportsConnectionInternalServerError{}
}

/* GetSnmpTransportsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetSnmpTransportsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports-connection][%d] getSnmpTransportsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSnmpTransportsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
