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

// ExportContentLibraryVMTemplateReader is a Reader for the ExportContentLibraryVMTemplate structure.
type ExportContentLibraryVMTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportContentLibraryVMTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportContentLibraryVMTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExportContentLibraryVMTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportContentLibraryVMTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExportContentLibraryVMTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportContentLibraryVMTemplateOK creates a ExportContentLibraryVMTemplateOK with default headers values
func NewExportContentLibraryVMTemplateOK() *ExportContentLibraryVMTemplateOK {
	return &ExportContentLibraryVMTemplateOK{}
}

/* ExportContentLibraryVMTemplateOK describes a response with status code 200, with default header values.

ExportContentLibraryVMTemplateOK export content library Vm template o k
*/
type ExportContentLibraryVMTemplateOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVMExportFile
}

func (o *ExportContentLibraryVMTemplateOK) Error() string {
	return fmt.Sprintf("[POST /export-content-library-vm-template][%d] exportContentLibraryVmTemplateOK  %+v", 200, o.Payload)
}
func (o *ExportContentLibraryVMTemplateOK) GetPayload() []*models.WithTaskVMExportFile {
	return o.Payload
}

func (o *ExportContentLibraryVMTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExportContentLibraryVMTemplateBadRequest creates a ExportContentLibraryVMTemplateBadRequest with default headers values
func NewExportContentLibraryVMTemplateBadRequest() *ExportContentLibraryVMTemplateBadRequest {
	return &ExportContentLibraryVMTemplateBadRequest{}
}

/* ExportContentLibraryVMTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ExportContentLibraryVMTemplateBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExportContentLibraryVMTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /export-content-library-vm-template][%d] exportContentLibraryVmTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *ExportContentLibraryVMTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExportContentLibraryVMTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExportContentLibraryVMTemplateNotFound creates a ExportContentLibraryVMTemplateNotFound with default headers values
func NewExportContentLibraryVMTemplateNotFound() *ExportContentLibraryVMTemplateNotFound {
	return &ExportContentLibraryVMTemplateNotFound{}
}

/* ExportContentLibraryVMTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type ExportContentLibraryVMTemplateNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExportContentLibraryVMTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /export-content-library-vm-template][%d] exportContentLibraryVmTemplateNotFound  %+v", 404, o.Payload)
}
func (o *ExportContentLibraryVMTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExportContentLibraryVMTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExportContentLibraryVMTemplateInternalServerError creates a ExportContentLibraryVMTemplateInternalServerError with default headers values
func NewExportContentLibraryVMTemplateInternalServerError() *ExportContentLibraryVMTemplateInternalServerError {
	return &ExportContentLibraryVMTemplateInternalServerError{}
}

/* ExportContentLibraryVMTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ExportContentLibraryVMTemplateInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExportContentLibraryVMTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /export-content-library-vm-template][%d] exportContentLibraryVmTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *ExportContentLibraryVMTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExportContentLibraryVMTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
