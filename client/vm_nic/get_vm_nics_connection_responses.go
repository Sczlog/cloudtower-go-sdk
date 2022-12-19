// Code generated by go-swagger; DO NOT EDIT.

package vm_nic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVMNicsConnectionReader is a Reader for the GetVMNicsConnection structure.
type GetVMNicsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMNicsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMNicsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMNicsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMNicsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVMNicsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMNicsConnectionOK creates a GetVMNicsConnectionOK with default headers values
func NewGetVMNicsConnectionOK() *GetVMNicsConnectionOK {
	return &GetVMNicsConnectionOK{}
}

/* GetVMNicsConnectionOK describes a response with status code 200, with default header values.

GetVMNicsConnectionOK get Vm nics connection o k
*/
type GetVMNicsConnectionOK struct {
	XTowerRequestID string

	Payload *models.VMNicConnection
}

func (o *GetVMNicsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics-connection][%d] getVmNicsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMNicsConnectionOK) GetPayload() *models.VMNicConnection {
	return o.Payload
}

func (o *GetVMNicsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.VMNicConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMNicsConnectionBadRequest creates a GetVMNicsConnectionBadRequest with default headers values
func NewGetVMNicsConnectionBadRequest() *GetVMNicsConnectionBadRequest {
	return &GetVMNicsConnectionBadRequest{}
}

/* GetVMNicsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVMNicsConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMNicsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics-connection][%d] getVmNicsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMNicsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMNicsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMNicsConnectionNotFound creates a GetVMNicsConnectionNotFound with default headers values
func NewGetVMNicsConnectionNotFound() *GetVMNicsConnectionNotFound {
	return &GetVMNicsConnectionNotFound{}
}

/* GetVMNicsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVMNicsConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMNicsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics-connection][%d] getVmNicsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVMNicsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMNicsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMNicsConnectionInternalServerError creates a GetVMNicsConnectionInternalServerError with default headers values
func NewGetVMNicsConnectionInternalServerError() *GetVMNicsConnectionInternalServerError {
	return &GetVMNicsConnectionInternalServerError{}
}

/* GetVMNicsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVMNicsConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMNicsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics-connection][%d] getVmNicsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVMNicsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMNicsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
