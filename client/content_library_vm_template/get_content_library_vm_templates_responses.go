// Code generated by go-swagger; DO NOT EDIT.

package content_library_vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetContentLibraryVMTemplatesReader is a Reader for the GetContentLibraryVMTemplates structure.
type GetContentLibraryVMTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentLibraryVMTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentLibraryVMTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentLibraryVMTemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentLibraryVMTemplatesOK creates a GetContentLibraryVMTemplatesOK with default headers values
func NewGetContentLibraryVMTemplatesOK() *GetContentLibraryVMTemplatesOK {
	return &GetContentLibraryVMTemplatesOK{}
}

/* GetContentLibraryVMTemplatesOK describes a response with status code 200, with default header values.

Ok
*/
type GetContentLibraryVMTemplatesOK struct {
	Payload []*models.ContentLibraryVMTemplate
}

func (o *GetContentLibraryVMTemplatesOK) Error() string {
	return fmt.Sprintf("[POST /get-content-library-vm-templates][%d] getContentLibraryVmTemplatesOK  %+v", 200, o.Payload)
}
func (o *GetContentLibraryVMTemplatesOK) GetPayload() []*models.ContentLibraryVMTemplate {
	return o.Payload
}

func (o *GetContentLibraryVMTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentLibraryVMTemplatesBadRequest creates a GetContentLibraryVMTemplatesBadRequest with default headers values
func NewGetContentLibraryVMTemplatesBadRequest() *GetContentLibraryVMTemplatesBadRequest {
	return &GetContentLibraryVMTemplatesBadRequest{}
}

/* GetContentLibraryVMTemplatesBadRequest describes a response with status code 400, with default header values.

GetContentLibraryVMTemplatesBadRequest get content library Vm templates bad request
*/
type GetContentLibraryVMTemplatesBadRequest struct {
	Payload string
}

func (o *GetContentLibraryVMTemplatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-content-library-vm-templates][%d] getContentLibraryVmTemplatesBadRequest  %+v", 400, o.Payload)
}
func (o *GetContentLibraryVMTemplatesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetContentLibraryVMTemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
