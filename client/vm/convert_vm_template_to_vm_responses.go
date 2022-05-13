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

// ConvertVMTemplateToVMReader is a Reader for the ConvertVMTemplateToVM structure.
type ConvertVMTemplateToVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConvertVMTemplateToVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConvertVMTemplateToVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConvertVMTemplateToVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConvertVMTemplateToVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConvertVMTemplateToVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConvertVMTemplateToVMOK creates a ConvertVMTemplateToVMOK with default headers values
func NewConvertVMTemplateToVMOK() *ConvertVMTemplateToVMOK {
	return &ConvertVMTemplateToVMOK{}
}

/* ConvertVMTemplateToVMOK describes a response with status code 200, with default header values.

Ok
*/
type ConvertVMTemplateToVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *ConvertVMTemplateToVMOK) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-to-vm][%d] convertVmTemplateToVmOK  %+v", 200, o.Payload)
}
func (o *ConvertVMTemplateToVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ConvertVMTemplateToVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertVMTemplateToVMBadRequest creates a ConvertVMTemplateToVMBadRequest with default headers values
func NewConvertVMTemplateToVMBadRequest() *ConvertVMTemplateToVMBadRequest {
	return &ConvertVMTemplateToVMBadRequest{}
}

/* ConvertVMTemplateToVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ConvertVMTemplateToVMBadRequest struct {
	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateToVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-to-vm][%d] convertVmTemplateToVmBadRequest  %+v", 400, o.Payload)
}
func (o *ConvertVMTemplateToVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateToVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertVMTemplateToVMNotFound creates a ConvertVMTemplateToVMNotFound with default headers values
func NewConvertVMTemplateToVMNotFound() *ConvertVMTemplateToVMNotFound {
	return &ConvertVMTemplateToVMNotFound{}
}

/* ConvertVMTemplateToVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type ConvertVMTemplateToVMNotFound struct {
	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateToVMNotFound) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-to-vm][%d] convertVmTemplateToVmNotFound  %+v", 404, o.Payload)
}
func (o *ConvertVMTemplateToVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateToVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertVMTemplateToVMInternalServerError creates a ConvertVMTemplateToVMInternalServerError with default headers values
func NewConvertVMTemplateToVMInternalServerError() *ConvertVMTemplateToVMInternalServerError {
	return &ConvertVMTemplateToVMInternalServerError{}
}

/* ConvertVMTemplateToVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ConvertVMTemplateToVMInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateToVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-to-vm][%d] convertVmTemplateToVmInternalServerError  %+v", 500, o.Payload)
}
func (o *ConvertVMTemplateToVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateToVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
