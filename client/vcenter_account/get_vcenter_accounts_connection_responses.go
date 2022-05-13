// Code generated by go-swagger; DO NOT EDIT.

package vcenter_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVcenterAccountsConnectionReader is a Reader for the GetVcenterAccountsConnection structure.
type GetVcenterAccountsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVcenterAccountsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVcenterAccountsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVcenterAccountsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVcenterAccountsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVcenterAccountsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVcenterAccountsConnectionOK creates a GetVcenterAccountsConnectionOK with default headers values
func NewGetVcenterAccountsConnectionOK() *GetVcenterAccountsConnectionOK {
	return &GetVcenterAccountsConnectionOK{}
}

/* GetVcenterAccountsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVcenterAccountsConnectionOK struct {
	Payload *models.VcenterAccountConnection
}

func (o *GetVcenterAccountsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVcenterAccountsConnectionOK) GetPayload() *models.VcenterAccountConnection {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcenterAccountConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterAccountsConnectionBadRequest creates a GetVcenterAccountsConnectionBadRequest with default headers values
func NewGetVcenterAccountsConnectionBadRequest() *GetVcenterAccountsConnectionBadRequest {
	return &GetVcenterAccountsConnectionBadRequest{}
}

/* GetVcenterAccountsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVcenterAccountsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVcenterAccountsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVcenterAccountsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterAccountsConnectionNotFound creates a GetVcenterAccountsConnectionNotFound with default headers values
func NewGetVcenterAccountsConnectionNotFound() *GetVcenterAccountsConnectionNotFound {
	return &GetVcenterAccountsConnectionNotFound{}
}

/* GetVcenterAccountsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVcenterAccountsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVcenterAccountsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVcenterAccountsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterAccountsConnectionInternalServerError creates a GetVcenterAccountsConnectionInternalServerError with default headers values
func NewGetVcenterAccountsConnectionInternalServerError() *GetVcenterAccountsConnectionInternalServerError {
	return &GetVcenterAccountsConnectionInternalServerError{}
}

/* GetVcenterAccountsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVcenterAccountsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVcenterAccountsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVcenterAccountsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
