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

// RemoveContentLibraryVMTemplateClustersReader is a Reader for the RemoveContentLibraryVMTemplateClusters structure.
type RemoveContentLibraryVMTemplateClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveContentLibraryVMTemplateClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveContentLibraryVMTemplateClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveContentLibraryVMTemplateClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveContentLibraryVMTemplateClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveContentLibraryVMTemplateClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveContentLibraryVMTemplateClustersOK creates a RemoveContentLibraryVMTemplateClustersOK with default headers values
func NewRemoveContentLibraryVMTemplateClustersOK() *RemoveContentLibraryVMTemplateClustersOK {
	return &RemoveContentLibraryVMTemplateClustersOK{}
}

/* RemoveContentLibraryVMTemplateClustersOK describes a response with status code 200, with default header values.

RemoveContentLibraryVMTemplateClustersOK remove content library Vm template clusters o k
*/
type RemoveContentLibraryVMTemplateClustersOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskContentLibraryVMTemplate
}

func (o *RemoveContentLibraryVMTemplateClustersOK) Error() string {
	return fmt.Sprintf("[POST /remove-content-library-vm-template-clusters][%d] removeContentLibraryVmTemplateClustersOK  %+v", 200, o.Payload)
}
func (o *RemoveContentLibraryVMTemplateClustersOK) GetPayload() []*models.WithTaskContentLibraryVMTemplate {
	return o.Payload
}

func (o *RemoveContentLibraryVMTemplateClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveContentLibraryVMTemplateClustersBadRequest creates a RemoveContentLibraryVMTemplateClustersBadRequest with default headers values
func NewRemoveContentLibraryVMTemplateClustersBadRequest() *RemoveContentLibraryVMTemplateClustersBadRequest {
	return &RemoveContentLibraryVMTemplateClustersBadRequest{}
}

/* RemoveContentLibraryVMTemplateClustersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RemoveContentLibraryVMTemplateClustersBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveContentLibraryVMTemplateClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /remove-content-library-vm-template-clusters][%d] removeContentLibraryVmTemplateClustersBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveContentLibraryVMTemplateClustersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveContentLibraryVMTemplateClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveContentLibraryVMTemplateClustersNotFound creates a RemoveContentLibraryVMTemplateClustersNotFound with default headers values
func NewRemoveContentLibraryVMTemplateClustersNotFound() *RemoveContentLibraryVMTemplateClustersNotFound {
	return &RemoveContentLibraryVMTemplateClustersNotFound{}
}

/* RemoveContentLibraryVMTemplateClustersNotFound describes a response with status code 404, with default header values.

Not found
*/
type RemoveContentLibraryVMTemplateClustersNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveContentLibraryVMTemplateClustersNotFound) Error() string {
	return fmt.Sprintf("[POST /remove-content-library-vm-template-clusters][%d] removeContentLibraryVmTemplateClustersNotFound  %+v", 404, o.Payload)
}
func (o *RemoveContentLibraryVMTemplateClustersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveContentLibraryVMTemplateClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveContentLibraryVMTemplateClustersInternalServerError creates a RemoveContentLibraryVMTemplateClustersInternalServerError with default headers values
func NewRemoveContentLibraryVMTemplateClustersInternalServerError() *RemoveContentLibraryVMTemplateClustersInternalServerError {
	return &RemoveContentLibraryVMTemplateClustersInternalServerError{}
}

/* RemoveContentLibraryVMTemplateClustersInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RemoveContentLibraryVMTemplateClustersInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveContentLibraryVMTemplateClustersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /remove-content-library-vm-template-clusters][%d] removeContentLibraryVmTemplateClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveContentLibraryVMTemplateClustersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveContentLibraryVMTemplateClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
