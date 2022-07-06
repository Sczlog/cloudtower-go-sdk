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

// ResetVMGuestOsPasswordReader is a Reader for the ResetVMGuestOsPassword structure.
type ResetVMGuestOsPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetVMGuestOsPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetVMGuestOsPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewResetVMGuestOsPasswordNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewResetVMGuestOsPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetVMGuestOsPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResetVMGuestOsPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetVMGuestOsPasswordOK creates a ResetVMGuestOsPasswordOK with default headers values
func NewResetVMGuestOsPasswordOK() *ResetVMGuestOsPasswordOK {
	return &ResetVMGuestOsPasswordOK{}
}

/* ResetVMGuestOsPasswordOK describes a response with status code 200, with default header values.

Ok
*/
type ResetVMGuestOsPasswordOK struct {
	Payload []interface{}
}

func (o *ResetVMGuestOsPasswordOK) Error() string {
	return fmt.Sprintf("[POST /reset-vm-guest-os-password][%d] resetVmGuestOsPasswordOK  %+v", 200, o.Payload)
}
func (o *ResetVMGuestOsPasswordOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *ResetVMGuestOsPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetVMGuestOsPasswordNotModified creates a ResetVMGuestOsPasswordNotModified with default headers values
func NewResetVMGuestOsPasswordNotModified() *ResetVMGuestOsPasswordNotModified {
	return &ResetVMGuestOsPasswordNotModified{}
}

/* ResetVMGuestOsPasswordNotModified describes a response with status code 304, with default header values.

Not modified
*/
type ResetVMGuestOsPasswordNotModified struct {
}

func (o *ResetVMGuestOsPasswordNotModified) Error() string {
	return fmt.Sprintf("[POST /reset-vm-guest-os-password][%d] resetVmGuestOsPasswordNotModified ", 304)
}

func (o *ResetVMGuestOsPasswordNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetVMGuestOsPasswordBadRequest creates a ResetVMGuestOsPasswordBadRequest with default headers values
func NewResetVMGuestOsPasswordBadRequest() *ResetVMGuestOsPasswordBadRequest {
	return &ResetVMGuestOsPasswordBadRequest{}
}

/* ResetVMGuestOsPasswordBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ResetVMGuestOsPasswordBadRequest struct {
	Payload *models.ErrorBody
}

func (o *ResetVMGuestOsPasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /reset-vm-guest-os-password][%d] resetVmGuestOsPasswordBadRequest  %+v", 400, o.Payload)
}
func (o *ResetVMGuestOsPasswordBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ResetVMGuestOsPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetVMGuestOsPasswordNotFound creates a ResetVMGuestOsPasswordNotFound with default headers values
func NewResetVMGuestOsPasswordNotFound() *ResetVMGuestOsPasswordNotFound {
	return &ResetVMGuestOsPasswordNotFound{}
}

/* ResetVMGuestOsPasswordNotFound describes a response with status code 404, with default header values.

Not found
*/
type ResetVMGuestOsPasswordNotFound struct {
	Payload *models.ErrorBody
}

func (o *ResetVMGuestOsPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /reset-vm-guest-os-password][%d] resetVmGuestOsPasswordNotFound  %+v", 404, o.Payload)
}
func (o *ResetVMGuestOsPasswordNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ResetVMGuestOsPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetVMGuestOsPasswordInternalServerError creates a ResetVMGuestOsPasswordInternalServerError with default headers values
func NewResetVMGuestOsPasswordInternalServerError() *ResetVMGuestOsPasswordInternalServerError {
	return &ResetVMGuestOsPasswordInternalServerError{}
}

/* ResetVMGuestOsPasswordInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ResetVMGuestOsPasswordInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *ResetVMGuestOsPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /reset-vm-guest-os-password][%d] resetVmGuestOsPasswordInternalServerError  %+v", 500, o.Payload)
}
func (o *ResetVMGuestOsPasswordInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ResetVMGuestOsPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
