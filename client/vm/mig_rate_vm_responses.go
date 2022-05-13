// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// MigRateVMReader is a Reader for the MigRateVM structure.
type MigRateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigRateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigRateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMigRateVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMigRateVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMigRateVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigRateVMOK creates a MigRateVMOK with default headers values
func NewMigRateVMOK() *MigRateVMOK {
	return &MigRateVMOK{}
}

/* MigRateVMOK describes a response with status code 200, with default header values.

Ok
*/
type MigRateVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *MigRateVMOK) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migRateVmOK  %+v", 200, o.Payload)
}
func (o *MigRateVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *MigRateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigRateVMBadRequest creates a MigRateVMBadRequest with default headers values
func NewMigRateVMBadRequest() *MigRateVMBadRequest {
	return &MigRateVMBadRequest{}
}

/* MigRateVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type MigRateVMBadRequest struct {
	Payload *models.ErrorBody
}

func (o *MigRateVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migRateVmBadRequest  %+v", 400, o.Payload)
}
func (o *MigRateVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigRateVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigRateVMNotFound creates a MigRateVMNotFound with default headers values
func NewMigRateVMNotFound() *MigRateVMNotFound {
	return &MigRateVMNotFound{}
}

/* MigRateVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type MigRateVMNotFound struct {
	Payload *models.ErrorBody
}

func (o *MigRateVMNotFound) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migRateVmNotFound  %+v", 404, o.Payload)
}
func (o *MigRateVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigRateVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigRateVMInternalServerError creates a MigRateVMInternalServerError with default headers values
func NewMigRateVMInternalServerError() *MigRateVMInternalServerError {
	return &MigRateVMInternalServerError{}
}

/* MigRateVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type MigRateVMInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *MigRateVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migRateVmInternalServerError  %+v", 500, o.Payload)
}
func (o *MigRateVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigRateVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
