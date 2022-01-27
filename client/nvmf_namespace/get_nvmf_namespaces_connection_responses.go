// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetNvmfNamespacesConnectionReader is a Reader for the GetNvmfNamespacesConnection structure.
type GetNvmfNamespacesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNvmfNamespacesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNvmfNamespacesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNvmfNamespacesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNvmfNamespacesConnectionOK creates a GetNvmfNamespacesConnectionOK with default headers values
func NewGetNvmfNamespacesConnectionOK() *GetNvmfNamespacesConnectionOK {
	return &GetNvmfNamespacesConnectionOK{}
}

/* GetNvmfNamespacesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNvmfNamespacesConnectionOK struct {
	Payload *models.NvmfNamespaceConnection
}

func (o *GetNvmfNamespacesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespaces-connection][%d] getNvmfNamespacesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNvmfNamespacesConnectionOK) GetPayload() *models.NvmfNamespaceConnection {
	return o.Payload
}

func (o *GetNvmfNamespacesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmfNamespaceConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNvmfNamespacesConnectionBadRequest creates a GetNvmfNamespacesConnectionBadRequest with default headers values
func NewGetNvmfNamespacesConnectionBadRequest() *GetNvmfNamespacesConnectionBadRequest {
	return &GetNvmfNamespacesConnectionBadRequest{}
}

/* GetNvmfNamespacesConnectionBadRequest describes a response with status code 400, with default header values.

GetNvmfNamespacesConnectionBadRequest get nvmf namespaces connection bad request
*/
type GetNvmfNamespacesConnectionBadRequest struct {
	Payload string
}

func (o *GetNvmfNamespacesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespaces-connection][%d] getNvmfNamespacesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNvmfNamespacesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNvmfNamespacesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
