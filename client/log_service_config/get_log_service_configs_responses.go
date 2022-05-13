// Code generated by go-swagger; DO NOT EDIT.

package log_service_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetLogServiceConfigsReader is a Reader for the GetLogServiceConfigs structure.
type GetLogServiceConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogServiceConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogServiceConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogServiceConfigsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLogServiceConfigsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogServiceConfigsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogServiceConfigsOK creates a GetLogServiceConfigsOK with default headers values
func NewGetLogServiceConfigsOK() *GetLogServiceConfigsOK {
	return &GetLogServiceConfigsOK{}
}

/* GetLogServiceConfigsOK describes a response with status code 200, with default header values.

Ok
*/
type GetLogServiceConfigsOK struct {
	Payload []*models.LogServiceConfig
}

func (o *GetLogServiceConfigsOK) Error() string {
	return fmt.Sprintf("[POST /get-log-service-configs][%d] getLogServiceConfigsOK  %+v", 200, o.Payload)
}
func (o *GetLogServiceConfigsOK) GetPayload() []*models.LogServiceConfig {
	return o.Payload
}

func (o *GetLogServiceConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogServiceConfigsBadRequest creates a GetLogServiceConfigsBadRequest with default headers values
func NewGetLogServiceConfigsBadRequest() *GetLogServiceConfigsBadRequest {
	return &GetLogServiceConfigsBadRequest{}
}

/* GetLogServiceConfigsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetLogServiceConfigsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLogServiceConfigsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-log-service-configs][%d] getLogServiceConfigsBadRequest  %+v", 400, o.Payload)
}
func (o *GetLogServiceConfigsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLogServiceConfigsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogServiceConfigsNotFound creates a GetLogServiceConfigsNotFound with default headers values
func NewGetLogServiceConfigsNotFound() *GetLogServiceConfigsNotFound {
	return &GetLogServiceConfigsNotFound{}
}

/* GetLogServiceConfigsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetLogServiceConfigsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLogServiceConfigsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-log-service-configs][%d] getLogServiceConfigsNotFound  %+v", 404, o.Payload)
}
func (o *GetLogServiceConfigsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLogServiceConfigsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogServiceConfigsInternalServerError creates a GetLogServiceConfigsInternalServerError with default headers values
func NewGetLogServiceConfigsInternalServerError() *GetLogServiceConfigsInternalServerError {
	return &GetLogServiceConfigsInternalServerError{}
}

/* GetLogServiceConfigsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetLogServiceConfigsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLogServiceConfigsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-log-service-configs][%d] getLogServiceConfigsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLogServiceConfigsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLogServiceConfigsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
