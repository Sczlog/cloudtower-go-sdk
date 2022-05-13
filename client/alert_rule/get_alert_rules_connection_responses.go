// Code generated by go-swagger; DO NOT EDIT.

package alert_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetAlertRulesConnectionReader is a Reader for the GetAlertRulesConnection structure.
type GetAlertRulesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRulesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRulesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertRulesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertRulesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertRulesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertRulesConnectionOK creates a GetAlertRulesConnectionOK with default headers values
func NewGetAlertRulesConnectionOK() *GetAlertRulesConnectionOK {
	return &GetAlertRulesConnectionOK{}
}

/* GetAlertRulesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetAlertRulesConnectionOK struct {
	Payload *models.AlertRuleConnection
}

func (o *GetAlertRulesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-alert-rules-connection][%d] getAlertRulesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetAlertRulesConnectionOK) GetPayload() *models.AlertRuleConnection {
	return o.Payload
}

func (o *GetAlertRulesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRuleConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRulesConnectionBadRequest creates a GetAlertRulesConnectionBadRequest with default headers values
func NewGetAlertRulesConnectionBadRequest() *GetAlertRulesConnectionBadRequest {
	return &GetAlertRulesConnectionBadRequest{}
}

/* GetAlertRulesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertRulesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAlertRulesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-alert-rules-connection][%d] getAlertRulesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertRulesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertRulesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRulesConnectionNotFound creates a GetAlertRulesConnectionNotFound with default headers values
func NewGetAlertRulesConnectionNotFound() *GetAlertRulesConnectionNotFound {
	return &GetAlertRulesConnectionNotFound{}
}

/* GetAlertRulesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAlertRulesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAlertRulesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-alert-rules-connection][%d] getAlertRulesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetAlertRulesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertRulesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRulesConnectionInternalServerError creates a GetAlertRulesConnectionInternalServerError with default headers values
func NewGetAlertRulesConnectionInternalServerError() *GetAlertRulesConnectionInternalServerError {
	return &GetAlertRulesConnectionInternalServerError{}
}

/* GetAlertRulesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetAlertRulesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAlertRulesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-alert-rules-connection][%d] getAlertRulesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertRulesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertRulesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
