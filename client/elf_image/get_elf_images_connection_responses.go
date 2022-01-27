// Code generated by go-swagger; DO NOT EDIT.

package elf_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetElfImagesConnectionReader is a Reader for the GetElfImagesConnection structure.
type GetElfImagesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetElfImagesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetElfImagesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetElfImagesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetElfImagesConnectionOK creates a GetElfImagesConnectionOK with default headers values
func NewGetElfImagesConnectionOK() *GetElfImagesConnectionOK {
	return &GetElfImagesConnectionOK{}
}

/* GetElfImagesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetElfImagesConnectionOK struct {
	Payload *models.ElfImageConnection
}

func (o *GetElfImagesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-elf-images-connection][%d] getElfImagesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetElfImagesConnectionOK) GetPayload() *models.ElfImageConnection {
	return o.Payload
}

func (o *GetElfImagesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElfImageConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetElfImagesConnectionBadRequest creates a GetElfImagesConnectionBadRequest with default headers values
func NewGetElfImagesConnectionBadRequest() *GetElfImagesConnectionBadRequest {
	return &GetElfImagesConnectionBadRequest{}
}

/* GetElfImagesConnectionBadRequest describes a response with status code 400, with default header values.

GetElfImagesConnectionBadRequest get elf images connection bad request
*/
type GetElfImagesConnectionBadRequest struct {
	Payload string
}

func (o *GetElfImagesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-elf-images-connection][%d] getElfImagesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetElfImagesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetElfImagesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
