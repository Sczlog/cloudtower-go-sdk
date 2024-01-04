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

// ImportContentLibraryVMTemplateReader is a Reader for the ImportContentLibraryVMTemplate structure.
type ImportContentLibraryVMTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportContentLibraryVMTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportContentLibraryVMTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportContentLibraryVMTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportContentLibraryVMTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportContentLibraryVMTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportContentLibraryVMTemplateOK creates a ImportContentLibraryVMTemplateOK with default headers values
func NewImportContentLibraryVMTemplateOK() *ImportContentLibraryVMTemplateOK {
	return &ImportContentLibraryVMTemplateOK{}
}

/* ImportContentLibraryVMTemplateOK describes a response with status code 200, with default header values.

ImportContentLibraryVMTemplateOK import content library Vm template o k
*/
type ImportContentLibraryVMTemplateOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskContentLibraryVMTemplate
}

func (o *ImportContentLibraryVMTemplateOK) Error() string {
	return fmt.Sprintf("[POST /import-content-library-vm-template][%d] importContentLibraryVmTemplateOK  %+v", 200, o.Payload)
}
func (o *ImportContentLibraryVMTemplateOK) GetPayload() []*models.WithTaskContentLibraryVMTemplate {
	return o.Payload
}

func (o *ImportContentLibraryVMTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportContentLibraryVMTemplateBadRequest creates a ImportContentLibraryVMTemplateBadRequest with default headers values
func NewImportContentLibraryVMTemplateBadRequest() *ImportContentLibraryVMTemplateBadRequest {
	return &ImportContentLibraryVMTemplateBadRequest{}
}

/* ImportContentLibraryVMTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ImportContentLibraryVMTemplateBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryVMTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /import-content-library-vm-template][%d] importContentLibraryVmTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *ImportContentLibraryVMTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryVMTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportContentLibraryVMTemplateNotFound creates a ImportContentLibraryVMTemplateNotFound with default headers values
func NewImportContentLibraryVMTemplateNotFound() *ImportContentLibraryVMTemplateNotFound {
	return &ImportContentLibraryVMTemplateNotFound{}
}

/* ImportContentLibraryVMTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type ImportContentLibraryVMTemplateNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryVMTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /import-content-library-vm-template][%d] importContentLibraryVmTemplateNotFound  %+v", 404, o.Payload)
}
func (o *ImportContentLibraryVMTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryVMTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportContentLibraryVMTemplateInternalServerError creates a ImportContentLibraryVMTemplateInternalServerError with default headers values
func NewImportContentLibraryVMTemplateInternalServerError() *ImportContentLibraryVMTemplateInternalServerError {
	return &ImportContentLibraryVMTemplateInternalServerError{}
}

/* ImportContentLibraryVMTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImportContentLibraryVMTemplateInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportContentLibraryVMTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /import-content-library-vm-template][%d] importContentLibraryVmTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *ImportContentLibraryVMTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportContentLibraryVMTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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