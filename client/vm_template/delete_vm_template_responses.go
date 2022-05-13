// Code generated by go-swagger; DO NOT EDIT.

package vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteVMTemplateReader is a Reader for the DeleteVMTemplate structure.
type DeleteVMTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMTemplateOK creates a DeleteVMTemplateOK with default headers values
func NewDeleteVMTemplateOK() *DeleteVMTemplateOK {
	return &DeleteVMTemplateOK{}
}

/* DeleteVMTemplateOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMTemplateOK struct {
	Payload []*models.WithTaskDeleteVMTemplate
}

func (o *DeleteVMTemplateOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-template][%d] deleteVmTemplateOK  %+v", 200, o.Payload)
}
func (o *DeleteVMTemplateOK) GetPayload() []*models.WithTaskDeleteVMTemplate {
	return o.Payload
}

func (o *DeleteVMTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMTemplateBadRequest creates a DeleteVMTemplateBadRequest with default headers values
func NewDeleteVMTemplateBadRequest() *DeleteVMTemplateBadRequest {
	return &DeleteVMTemplateBadRequest{}
}

/* DeleteVMTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMTemplateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-template][%d] deleteVmTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMTemplateNotFound creates a DeleteVMTemplateNotFound with default headers values
func NewDeleteVMTemplateNotFound() *DeleteVMTemplateNotFound {
	return &DeleteVMTemplateNotFound{}
}

/* DeleteVMTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMTemplateNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm-template][%d] deleteVmTemplateNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMTemplateInternalServerError creates a DeleteVMTemplateInternalServerError with default headers values
func NewDeleteVMTemplateInternalServerError() *DeleteVMTemplateInternalServerError {
	return &DeleteVMTemplateInternalServerError{}
}

/* DeleteVMTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm-template][%d] deleteVmTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
