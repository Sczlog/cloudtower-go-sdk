// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// AddVMToFolderReader is a Reader for the AddVMToFolder structure.
type AddVMToFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVMToFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddVMToFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddVMToFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddVMToFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddVMToFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddVMToFolderOK creates a AddVMToFolderOK with default headers values
func NewAddVMToFolderOK() *AddVMToFolderOK {
	return &AddVMToFolderOK{}
}

/* AddVMToFolderOK describes a response with status code 200, with default header values.

AddVMToFolderOK add Vm to folder o k
*/
type AddVMToFolderOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *AddVMToFolderOK) Error() string {
	return fmt.Sprintf("[POST /add-vm-to-folder][%d] addVmToFolderOK  %+v", 200, o.Payload)
}
func (o *AddVMToFolderOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *AddVMToFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddVMToFolderBadRequest creates a AddVMToFolderBadRequest with default headers values
func NewAddVMToFolderBadRequest() *AddVMToFolderBadRequest {
	return &AddVMToFolderBadRequest{}
}

/* AddVMToFolderBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddVMToFolderBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *AddVMToFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /add-vm-to-folder][%d] addVmToFolderBadRequest  %+v", 400, o.Payload)
}
func (o *AddVMToFolderBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *AddVMToFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddVMToFolderNotFound creates a AddVMToFolderNotFound with default headers values
func NewAddVMToFolderNotFound() *AddVMToFolderNotFound {
	return &AddVMToFolderNotFound{}
}

/* AddVMToFolderNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddVMToFolderNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *AddVMToFolderNotFound) Error() string {
	return fmt.Sprintf("[POST /add-vm-to-folder][%d] addVmToFolderNotFound  %+v", 404, o.Payload)
}
func (o *AddVMToFolderNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *AddVMToFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddVMToFolderInternalServerError creates a AddVMToFolderInternalServerError with default headers values
func NewAddVMToFolderInternalServerError() *AddVMToFolderInternalServerError {
	return &AddVMToFolderInternalServerError{}
}

/* AddVMToFolderInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AddVMToFolderInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *AddVMToFolderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /add-vm-to-folder][%d] addVmToFolderInternalServerError  %+v", 500, o.Payload)
}
func (o *AddVMToFolderInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *AddVMToFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
