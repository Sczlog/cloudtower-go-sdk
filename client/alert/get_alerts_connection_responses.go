// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetAlertsConnectionReader is a Reader for the GetAlertsConnection structure.
type GetAlertsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertsConnectionOK creates a GetAlertsConnectionOK with default headers values
func NewGetAlertsConnectionOK() *GetAlertsConnectionOK {
	return &GetAlertsConnectionOK{}
}

/* GetAlertsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetAlertsConnectionOK struct {
	Payload *models.AlertConnection
}

func (o *GetAlertsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-alerts-connection][%d] getAlertsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetAlertsConnectionOK) GetPayload() *models.AlertConnection {
	return o.Payload
}

func (o *GetAlertsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsConnectionBadRequest creates a GetAlertsConnectionBadRequest with default headers values
func NewGetAlertsConnectionBadRequest() *GetAlertsConnectionBadRequest {
	return &GetAlertsConnectionBadRequest{}
}

/* GetAlertsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAlertsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-alerts-connection][%d] getAlertsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsConnectionNotFound creates a GetAlertsConnectionNotFound with default headers values
func NewGetAlertsConnectionNotFound() *GetAlertsConnectionNotFound {
	return &GetAlertsConnectionNotFound{}
}

/* GetAlertsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAlertsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAlertsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-alerts-connection][%d] getAlertsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetAlertsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertsConnectionInternalServerError creates a GetAlertsConnectionInternalServerError with default headers values
func NewGetAlertsConnectionInternalServerError() *GetAlertsConnectionInternalServerError {
	return &GetAlertsConnectionInternalServerError{}
}

/* GetAlertsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetAlertsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAlertsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-alerts-connection][%d] getAlertsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
