// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteVMFolderReader is a Reader for the DeleteVMFolder structure.
type DeleteVMFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMFolderOK creates a DeleteVMFolderOK with default headers values
func NewDeleteVMFolderOK() *DeleteVMFolderOK {
	return &DeleteVMFolderOK{}
}

/* DeleteVMFolderOK describes a response with status code 200, with default header values.

DeleteVMFolderOK delete Vm folder o k
*/
type DeleteVMFolderOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteVMFolder
}

func (o *DeleteVMFolderOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-folder][%d] deleteVmFolderOK  %+v", 200, o.Payload)
}
func (o *DeleteVMFolderOK) GetPayload() []*models.WithTaskDeleteVMFolder {
	return o.Payload
}

func (o *DeleteVMFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMFolderBadRequest creates a DeleteVMFolderBadRequest with default headers values
func NewDeleteVMFolderBadRequest() *DeleteVMFolderBadRequest {
	return &DeleteVMFolderBadRequest{}
}

/* DeleteVMFolderBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMFolderBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-folder][%d] deleteVmFolderBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMFolderBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMFolderNotFound creates a DeleteVMFolderNotFound with default headers values
func NewDeleteVMFolderNotFound() *DeleteVMFolderNotFound {
	return &DeleteVMFolderNotFound{}
}

/* DeleteVMFolderNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMFolderNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMFolderNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm-folder][%d] deleteVmFolderNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMFolderNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMFolderInternalServerError creates a DeleteVMFolderInternalServerError with default headers values
func NewDeleteVMFolderInternalServerError() *DeleteVMFolderInternalServerError {
	return &DeleteVMFolderInternalServerError{}
}

/* DeleteVMFolderInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMFolderInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMFolderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm-folder][%d] deleteVmFolderInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMFolderInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
