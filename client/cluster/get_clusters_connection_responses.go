// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetClustersConnectionReader is a Reader for the GetClustersConnection structure.
type GetClustersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClustersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClustersConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClustersConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClustersConnectionOK creates a GetClustersConnectionOK with default headers values
func NewGetClustersConnectionOK() *GetClustersConnectionOK {
	return &GetClustersConnectionOK{}
}

/* GetClustersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetClustersConnectionOK struct {
	Payload *models.ClusterConnection
}

func (o *GetClustersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClustersConnectionOK) GetPayload() *models.ClusterConnection {
	return o.Payload
}

func (o *GetClustersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersConnectionBadRequest creates a GetClustersConnectionBadRequest with default headers values
func NewGetClustersConnectionBadRequest() *GetClustersConnectionBadRequest {
	return &GetClustersConnectionBadRequest{}
}

/* GetClustersConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetClustersConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetClustersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClustersConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClustersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersConnectionNotFound creates a GetClustersConnectionNotFound with default headers values
func NewGetClustersConnectionNotFound() *GetClustersConnectionNotFound {
	return &GetClustersConnectionNotFound{}
}

/* GetClustersConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetClustersConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetClustersConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetClustersConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClustersConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersConnectionInternalServerError creates a GetClustersConnectionInternalServerError with default headers values
func NewGetClustersConnectionInternalServerError() *GetClustersConnectionInternalServerError {
	return &GetClustersConnectionInternalServerError{}
}

/* GetClustersConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetClustersConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetClustersConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetClustersConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClustersConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
