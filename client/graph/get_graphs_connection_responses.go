// Code generated by go-swagger; DO NOT EDIT.

package graph

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetGraphsConnectionReader is a Reader for the GetGraphsConnection structure.
type GetGraphsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGraphsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGraphsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGraphsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGraphsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGraphsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGraphsConnectionOK creates a GetGraphsConnectionOK with default headers values
func NewGetGraphsConnectionOK() *GetGraphsConnectionOK {
	return &GetGraphsConnectionOK{}
}

/* GetGraphsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetGraphsConnectionOK struct {
	Payload *models.GraphConnection
}

func (o *GetGraphsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-graphs-connection][%d] getGraphsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetGraphsConnectionOK) GetPayload() *models.GraphConnection {
	return o.Payload
}

func (o *GetGraphsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGraphsConnectionBadRequest creates a GetGraphsConnectionBadRequest with default headers values
func NewGetGraphsConnectionBadRequest() *GetGraphsConnectionBadRequest {
	return &GetGraphsConnectionBadRequest{}
}

/* GetGraphsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetGraphsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGraphsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-graphs-connection][%d] getGraphsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetGraphsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGraphsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGraphsConnectionNotFound creates a GetGraphsConnectionNotFound with default headers values
func NewGetGraphsConnectionNotFound() *GetGraphsConnectionNotFound {
	return &GetGraphsConnectionNotFound{}
}

/* GetGraphsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetGraphsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGraphsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-graphs-connection][%d] getGraphsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetGraphsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGraphsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGraphsConnectionInternalServerError creates a GetGraphsConnectionInternalServerError with default headers values
func NewGetGraphsConnectionInternalServerError() *GetGraphsConnectionInternalServerError {
	return &GetGraphsConnectionInternalServerError{}
}

/* GetGraphsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetGraphsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGraphsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-graphs-connection][%d] getGraphsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGraphsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGraphsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
