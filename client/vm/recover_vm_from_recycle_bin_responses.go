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

// RecoverVMFromRecycleBinReader is a Reader for the RecoverVMFromRecycleBin structure.
type RecoverVMFromRecycleBinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverVMFromRecycleBinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRecoverVMFromRecycleBinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRecoverVMFromRecycleBinBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRecoverVMFromRecycleBinNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRecoverVMFromRecycleBinInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRecoverVMFromRecycleBinOK creates a RecoverVMFromRecycleBinOK with default headers values
func NewRecoverVMFromRecycleBinOK() *RecoverVMFromRecycleBinOK {
	return &RecoverVMFromRecycleBinOK{}
}

/* RecoverVMFromRecycleBinOK describes a response with status code 200, with default header values.

Ok
*/
type RecoverVMFromRecycleBinOK struct {
	Payload []*models.WithTaskDeleteVM
}

func (o *RecoverVMFromRecycleBinOK) Error() string {
	return fmt.Sprintf("[POST /recover-vm-from-recycle][%d] recoverVmFromRecycleBinOK  %+v", 200, o.Payload)
}
func (o *RecoverVMFromRecycleBinOK) GetPayload() []*models.WithTaskDeleteVM {
	return o.Payload
}

func (o *RecoverVMFromRecycleBinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecoverVMFromRecycleBinBadRequest creates a RecoverVMFromRecycleBinBadRequest with default headers values
func NewRecoverVMFromRecycleBinBadRequest() *RecoverVMFromRecycleBinBadRequest {
	return &RecoverVMFromRecycleBinBadRequest{}
}

/* RecoverVMFromRecycleBinBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RecoverVMFromRecycleBinBadRequest struct {
	Payload *models.ErrorBody
}

func (o *RecoverVMFromRecycleBinBadRequest) Error() string {
	return fmt.Sprintf("[POST /recover-vm-from-recycle][%d] recoverVmFromRecycleBinBadRequest  %+v", 400, o.Payload)
}
func (o *RecoverVMFromRecycleBinBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RecoverVMFromRecycleBinBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecoverVMFromRecycleBinNotFound creates a RecoverVMFromRecycleBinNotFound with default headers values
func NewRecoverVMFromRecycleBinNotFound() *RecoverVMFromRecycleBinNotFound {
	return &RecoverVMFromRecycleBinNotFound{}
}

/* RecoverVMFromRecycleBinNotFound describes a response with status code 404, with default header values.

Not found
*/
type RecoverVMFromRecycleBinNotFound struct {
	Payload *models.ErrorBody
}

func (o *RecoverVMFromRecycleBinNotFound) Error() string {
	return fmt.Sprintf("[POST /recover-vm-from-recycle][%d] recoverVmFromRecycleBinNotFound  %+v", 404, o.Payload)
}
func (o *RecoverVMFromRecycleBinNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RecoverVMFromRecycleBinNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecoverVMFromRecycleBinInternalServerError creates a RecoverVMFromRecycleBinInternalServerError with default headers values
func NewRecoverVMFromRecycleBinInternalServerError() *RecoverVMFromRecycleBinInternalServerError {
	return &RecoverVMFromRecycleBinInternalServerError{}
}

/* RecoverVMFromRecycleBinInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RecoverVMFromRecycleBinInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *RecoverVMFromRecycleBinInternalServerError) Error() string {
	return fmt.Sprintf("[POST /recover-vm-from-recycle][%d] recoverVmFromRecycleBinInternalServerError  %+v", 500, o.Payload)
}
func (o *RecoverVMFromRecycleBinInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RecoverVMFromRecycleBinInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
