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

// CloneVMReader is a Reader for the CloneVM structure.
type CloneVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloneVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneVMOK creates a CloneVMOK with default headers values
func NewCloneVMOK() *CloneVMOK {
	return &CloneVMOK{}
}

/* CloneVMOK describes a response with status code 200, with default header values.

Ok
*/
type CloneVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *CloneVMOK) Error() string {
	return fmt.Sprintf("[POST /clone-vm][%d] cloneVmOK  %+v", 200, o.Payload)
}
func (o *CloneVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *CloneVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneVMBadRequest creates a CloneVMBadRequest with default headers values
func NewCloneVMBadRequest() *CloneVMBadRequest {
	return &CloneVMBadRequest{}
}

/* CloneVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CloneVMBadRequest struct {
	Payload *models.ErrorBody
}

func (o *CloneVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /clone-vm][%d] cloneVmBadRequest  %+v", 400, o.Payload)
}
func (o *CloneVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneVMNotFound creates a CloneVMNotFound with default headers values
func NewCloneVMNotFound() *CloneVMNotFound {
	return &CloneVMNotFound{}
}

/* CloneVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type CloneVMNotFound struct {
	Payload *models.ErrorBody
}

func (o *CloneVMNotFound) Error() string {
	return fmt.Sprintf("[POST /clone-vm][%d] cloneVmNotFound  %+v", 404, o.Payload)
}
func (o *CloneVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneVMInternalServerError creates a CloneVMInternalServerError with default headers values
func NewCloneVMInternalServerError() *CloneVMInternalServerError {
	return &CloneVMInternalServerError{}
}

/* CloneVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CloneVMInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *CloneVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clone-vm][%d] cloneVmInternalServerError  %+v", 500, o.Payload)
}
func (o *CloneVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
