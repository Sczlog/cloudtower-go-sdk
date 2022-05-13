// Code generated by go-swagger; DO NOT EDIT.

package elf_storage_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetElfStoragePoliciesConnectionReader is a Reader for the GetElfStoragePoliciesConnection structure.
type GetElfStoragePoliciesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetElfStoragePoliciesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetElfStoragePoliciesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetElfStoragePoliciesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetElfStoragePoliciesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetElfStoragePoliciesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetElfStoragePoliciesConnectionOK creates a GetElfStoragePoliciesConnectionOK with default headers values
func NewGetElfStoragePoliciesConnectionOK() *GetElfStoragePoliciesConnectionOK {
	return &GetElfStoragePoliciesConnectionOK{}
}

/* GetElfStoragePoliciesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetElfStoragePoliciesConnectionOK struct {
	Payload *models.ElfStoragePolicyConnection
}

func (o *GetElfStoragePoliciesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-elf-storage-policies-connection][%d] getElfStoragePoliciesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetElfStoragePoliciesConnectionOK) GetPayload() *models.ElfStoragePolicyConnection {
	return o.Payload
}

func (o *GetElfStoragePoliciesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElfStoragePolicyConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetElfStoragePoliciesConnectionBadRequest creates a GetElfStoragePoliciesConnectionBadRequest with default headers values
func NewGetElfStoragePoliciesConnectionBadRequest() *GetElfStoragePoliciesConnectionBadRequest {
	return &GetElfStoragePoliciesConnectionBadRequest{}
}

/* GetElfStoragePoliciesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetElfStoragePoliciesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetElfStoragePoliciesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-elf-storage-policies-connection][%d] getElfStoragePoliciesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetElfStoragePoliciesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetElfStoragePoliciesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetElfStoragePoliciesConnectionNotFound creates a GetElfStoragePoliciesConnectionNotFound with default headers values
func NewGetElfStoragePoliciesConnectionNotFound() *GetElfStoragePoliciesConnectionNotFound {
	return &GetElfStoragePoliciesConnectionNotFound{}
}

/* GetElfStoragePoliciesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetElfStoragePoliciesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetElfStoragePoliciesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-elf-storage-policies-connection][%d] getElfStoragePoliciesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetElfStoragePoliciesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetElfStoragePoliciesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetElfStoragePoliciesConnectionInternalServerError creates a GetElfStoragePoliciesConnectionInternalServerError with default headers values
func NewGetElfStoragePoliciesConnectionInternalServerError() *GetElfStoragePoliciesConnectionInternalServerError {
	return &GetElfStoragePoliciesConnectionInternalServerError{}
}

/* GetElfStoragePoliciesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetElfStoragePoliciesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetElfStoragePoliciesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-elf-storage-policies-connection][%d] getElfStoragePoliciesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetElfStoragePoliciesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetElfStoragePoliciesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
