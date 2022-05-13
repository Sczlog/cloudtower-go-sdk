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

// GetSnmpTransportsReader is a Reader for the GetSnmpTransports structure.
type GetSnmpTransportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnmpTransportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnmpTransportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnmpTransportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnmpTransportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSnmpTransportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnmpTransportsOK creates a GetSnmpTransportsOK with default headers values
func NewGetSnmpTransportsOK() *GetSnmpTransportsOK {
	return &GetSnmpTransportsOK{}
}

/* GetSnmpTransportsOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnmpTransportsOK struct {
	Payload []*models.SnmpTransport
}

func (o *GetSnmpTransportsOK) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports][%d] getSnmpTransportsOK  %+v", 200, o.Payload)
}
func (o *GetSnmpTransportsOK) GetPayload() []*models.SnmpTransport {
	return o.Payload
}

func (o *GetSnmpTransportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsBadRequest creates a GetSnmpTransportsBadRequest with default headers values
func NewGetSnmpTransportsBadRequest() *GetSnmpTransportsBadRequest {
	return &GetSnmpTransportsBadRequest{}
}

/* GetSnmpTransportsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSnmpTransportsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports][%d] getSnmpTransportsBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnmpTransportsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsNotFound creates a GetSnmpTransportsNotFound with default headers values
func NewGetSnmpTransportsNotFound() *GetSnmpTransportsNotFound {
	return &GetSnmpTransportsNotFound{}
}

/* GetSnmpTransportsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSnmpTransportsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports][%d] getSnmpTransportsNotFound  %+v", 404, o.Payload)
}
func (o *GetSnmpTransportsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTransportsInternalServerError creates a GetSnmpTransportsInternalServerError with default headers values
func NewGetSnmpTransportsInternalServerError() *GetSnmpTransportsInternalServerError {
	return &GetSnmpTransportsInternalServerError{}
}

/* GetSnmpTransportsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetSnmpTransportsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSnmpTransportsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-snmp-transports][%d] getSnmpTransportsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSnmpTransportsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnmpTransportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
