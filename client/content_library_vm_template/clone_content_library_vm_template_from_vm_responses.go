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

// CloneContentLibraryVMTemplateFromVMReader is a Reader for the CloneContentLibraryVMTemplateFromVM structure.
type CloneContentLibraryVMTemplateFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneContentLibraryVMTemplateFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneContentLibraryVMTemplateFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloneContentLibraryVMTemplateFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneContentLibraryVMTemplateFromVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneContentLibraryVMTemplateFromVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneContentLibraryVMTemplateFromVMOK creates a CloneContentLibraryVMTemplateFromVMOK with default headers values
func NewCloneContentLibraryVMTemplateFromVMOK() *CloneContentLibraryVMTemplateFromVMOK {
	return &CloneContentLibraryVMTemplateFromVMOK{}
}

/* CloneContentLibraryVMTemplateFromVMOK describes a response with status code 200, with default header values.

Ok
*/
type CloneContentLibraryVMTemplateFromVMOK struct {
	Payload []*models.WithTaskContentLibraryVMTemplate
}

func (o *CloneContentLibraryVMTemplateFromVMOK) Error() string {
	return fmt.Sprintf("[POST /clone-content-library-vm-template-from-vm][%d] cloneContentLibraryVmTemplateFromVmOK  %+v", 200, o.Payload)
}
func (o *CloneContentLibraryVMTemplateFromVMOK) GetPayload() []*models.WithTaskContentLibraryVMTemplate {
	return o.Payload
}

func (o *CloneContentLibraryVMTemplateFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneContentLibraryVMTemplateFromVMBadRequest creates a CloneContentLibraryVMTemplateFromVMBadRequest with default headers values
func NewCloneContentLibraryVMTemplateFromVMBadRequest() *CloneContentLibraryVMTemplateFromVMBadRequest {
	return &CloneContentLibraryVMTemplateFromVMBadRequest{}
}

/* CloneContentLibraryVMTemplateFromVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CloneContentLibraryVMTemplateFromVMBadRequest struct {
	Payload *models.ErrorBody
}

func (o *CloneContentLibraryVMTemplateFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /clone-content-library-vm-template-from-vm][%d] cloneContentLibraryVmTemplateFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *CloneContentLibraryVMTemplateFromVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneContentLibraryVMTemplateFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneContentLibraryVMTemplateFromVMNotFound creates a CloneContentLibraryVMTemplateFromVMNotFound with default headers values
func NewCloneContentLibraryVMTemplateFromVMNotFound() *CloneContentLibraryVMTemplateFromVMNotFound {
	return &CloneContentLibraryVMTemplateFromVMNotFound{}
}

/* CloneContentLibraryVMTemplateFromVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type CloneContentLibraryVMTemplateFromVMNotFound struct {
	Payload *models.ErrorBody
}

func (o *CloneContentLibraryVMTemplateFromVMNotFound) Error() string {
	return fmt.Sprintf("[POST /clone-content-library-vm-template-from-vm][%d] cloneContentLibraryVmTemplateFromVmNotFound  %+v", 404, o.Payload)
}
func (o *CloneContentLibraryVMTemplateFromVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneContentLibraryVMTemplateFromVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneContentLibraryVMTemplateFromVMInternalServerError creates a CloneContentLibraryVMTemplateFromVMInternalServerError with default headers values
func NewCloneContentLibraryVMTemplateFromVMInternalServerError() *CloneContentLibraryVMTemplateFromVMInternalServerError {
	return &CloneContentLibraryVMTemplateFromVMInternalServerError{}
}

/* CloneContentLibraryVMTemplateFromVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CloneContentLibraryVMTemplateFromVMInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *CloneContentLibraryVMTemplateFromVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clone-content-library-vm-template-from-vm][%d] cloneContentLibraryVmTemplateFromVmInternalServerError  %+v", 500, o.Payload)
}
func (o *CloneContentLibraryVMTemplateFromVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CloneContentLibraryVMTemplateFromVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
