// Code generated by go-swagger; DO NOT EDIT.

package snmp_transport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// DeleteSnmpTransportReader is a Reader for the DeleteSnmpTransport structure.
type DeleteSnmpTransportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnmpTransportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSnmpTransportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSnmpTransportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSnmpTransportOK creates a DeleteSnmpTransportOK with default headers values
func NewDeleteSnmpTransportOK() *DeleteSnmpTransportOK {
	return &DeleteSnmpTransportOK{}
}

/* DeleteSnmpTransportOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteSnmpTransportOK struct {
	Payload []*models.WithTaskDeleteSnmpTransport
}

func (o *DeleteSnmpTransportOK) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-transport][%d] deleteSnmpTransportOK  %+v", 200, o.Payload)
}
func (o *DeleteSnmpTransportOK) GetPayload() []*models.WithTaskDeleteSnmpTransport {
	return o.Payload
}

func (o *DeleteSnmpTransportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnmpTransportBadRequest creates a DeleteSnmpTransportBadRequest with default headers values
func NewDeleteSnmpTransportBadRequest() *DeleteSnmpTransportBadRequest {
	return &DeleteSnmpTransportBadRequest{}
}

/* DeleteSnmpTransportBadRequest describes a response with status code 400, with default header values.

DeleteSnmpTransportBadRequest delete snmp transport bad request
*/
type DeleteSnmpTransportBadRequest struct {
	Payload string
}

func (o *DeleteSnmpTransportBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-snmp-transport][%d] deleteSnmpTransportBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteSnmpTransportBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteSnmpTransportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
