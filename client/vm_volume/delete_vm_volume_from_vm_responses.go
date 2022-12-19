// Code generated by go-swagger; DO NOT EDIT.

package vm_volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteVMVolumeFromVMReader is a Reader for the DeleteVMVolumeFromVM structure.
type DeleteVMVolumeFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMVolumeFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMVolumeFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMVolumeFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMVolumeFromVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMVolumeFromVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMVolumeFromVMOK creates a DeleteVMVolumeFromVMOK with default headers values
func NewDeleteVMVolumeFromVMOK() *DeleteVMVolumeFromVMOK {
	return &DeleteVMVolumeFromVMOK{}
}

/* DeleteVMVolumeFromVMOK describes a response with status code 200, with default header values.

DeleteVMVolumeFromVMOK delete Vm volume from Vm o k
*/
type DeleteVMVolumeFromVMOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteVMVolume
}

func (o *DeleteVMVolumeFromVMOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmOK  %+v", 200, o.Payload)
}
func (o *DeleteVMVolumeFromVMOK) GetPayload() []*models.WithTaskDeleteVMVolume {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeFromVMBadRequest creates a DeleteVMVolumeFromVMBadRequest with default headers values
func NewDeleteVMVolumeFromVMBadRequest() *DeleteVMVolumeFromVMBadRequest {
	return &DeleteVMVolumeFromVMBadRequest{}
}

/* DeleteVMVolumeFromVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMVolumeFromVMBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMVolumeFromVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeFromVMNotFound creates a DeleteVMVolumeFromVMNotFound with default headers values
func NewDeleteVMVolumeFromVMNotFound() *DeleteVMVolumeFromVMNotFound {
	return &DeleteVMVolumeFromVMNotFound{}
}

/* DeleteVMVolumeFromVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMVolumeFromVMNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeFromVMNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMVolumeFromVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeFromVMInternalServerError creates a DeleteVMVolumeFromVMInternalServerError with default headers values
func NewDeleteVMVolumeFromVMInternalServerError() *DeleteVMVolumeFromVMInternalServerError {
	return &DeleteVMVolumeFromVMInternalServerError{}
}

/* DeleteVMVolumeFromVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMVolumeFromVMInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeFromVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMVolumeFromVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
