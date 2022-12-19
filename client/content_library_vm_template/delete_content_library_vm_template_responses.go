// Code generated by go-swagger; DO NOT EDIT.

package content_library_vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteContentLibraryVMTemplateReader is a Reader for the DeleteContentLibraryVMTemplate structure.
type DeleteContentLibraryVMTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentLibraryVMTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteContentLibraryVMTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteContentLibraryVMTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteContentLibraryVMTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteContentLibraryVMTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteContentLibraryVMTemplateOK creates a DeleteContentLibraryVMTemplateOK with default headers values
func NewDeleteContentLibraryVMTemplateOK() *DeleteContentLibraryVMTemplateOK {
	return &DeleteContentLibraryVMTemplateOK{}
}

/* DeleteContentLibraryVMTemplateOK describes a response with status code 200, with default header values.

DeleteContentLibraryVMTemplateOK delete content library Vm template o k
*/
type DeleteContentLibraryVMTemplateOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteContentLibraryVMTemplate
}

func (o *DeleteContentLibraryVMTemplateOK) Error() string {
	return fmt.Sprintf("[POST /delete-content-library-vm-template][%d] deleteContentLibraryVmTemplateOK  %+v", 200, o.Payload)
}
func (o *DeleteContentLibraryVMTemplateOK) GetPayload() []*models.WithTaskDeleteContentLibraryVMTemplate {
	return o.Payload
}

func (o *DeleteContentLibraryVMTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteContentLibraryVMTemplateBadRequest creates a DeleteContentLibraryVMTemplateBadRequest with default headers values
func NewDeleteContentLibraryVMTemplateBadRequest() *DeleteContentLibraryVMTemplateBadRequest {
	return &DeleteContentLibraryVMTemplateBadRequest{}
}

/* DeleteContentLibraryVMTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteContentLibraryVMTemplateBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteContentLibraryVMTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-content-library-vm-template][%d] deleteContentLibraryVmTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteContentLibraryVMTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentLibraryVMTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteContentLibraryVMTemplateNotFound creates a DeleteContentLibraryVMTemplateNotFound with default headers values
func NewDeleteContentLibraryVMTemplateNotFound() *DeleteContentLibraryVMTemplateNotFound {
	return &DeleteContentLibraryVMTemplateNotFound{}
}

/* DeleteContentLibraryVMTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteContentLibraryVMTemplateNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteContentLibraryVMTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-content-library-vm-template][%d] deleteContentLibraryVmTemplateNotFound  %+v", 404, o.Payload)
}
func (o *DeleteContentLibraryVMTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentLibraryVMTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteContentLibraryVMTemplateInternalServerError creates a DeleteContentLibraryVMTemplateInternalServerError with default headers values
func NewDeleteContentLibraryVMTemplateInternalServerError() *DeleteContentLibraryVMTemplateInternalServerError {
	return &DeleteContentLibraryVMTemplateInternalServerError{}
}

/* DeleteContentLibraryVMTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteContentLibraryVMTemplateInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteContentLibraryVMTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-content-library-vm-template][%d] deleteContentLibraryVmTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteContentLibraryVMTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentLibraryVMTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
