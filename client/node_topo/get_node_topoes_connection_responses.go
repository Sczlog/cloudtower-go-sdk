// Code generated by go-swagger; DO NOT EDIT.

package node_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetNodeTopoesConnectionReader is a Reader for the GetNodeTopoesConnection structure.
type GetNodeTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNodeTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNodeTopoesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNodeTopoesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNodeTopoesConnectionOK creates a GetNodeTopoesConnectionOK with default headers values
func NewGetNodeTopoesConnectionOK() *GetNodeTopoesConnectionOK {
	return &GetNodeTopoesConnectionOK{}
}

/* GetNodeTopoesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNodeTopoesConnectionOK struct {
	Payload *models.NodeTopoConnection
}

func (o *GetNodeTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-node-topoes-connection][%d] getNodeTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNodeTopoesConnectionOK) GetPayload() *models.NodeTopoConnection {
	return o.Payload
}

func (o *GetNodeTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeTopoesConnectionBadRequest creates a GetNodeTopoesConnectionBadRequest with default headers values
func NewGetNodeTopoesConnectionBadRequest() *GetNodeTopoesConnectionBadRequest {
	return &GetNodeTopoesConnectionBadRequest{}
}

/* GetNodeTopoesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetNodeTopoesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetNodeTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-node-topoes-connection][%d] getNodeTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNodeTopoesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNodeTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeTopoesConnectionNotFound creates a GetNodeTopoesConnectionNotFound with default headers values
func NewGetNodeTopoesConnectionNotFound() *GetNodeTopoesConnectionNotFound {
	return &GetNodeTopoesConnectionNotFound{}
}

/* GetNodeTopoesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetNodeTopoesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetNodeTopoesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-node-topoes-connection][%d] getNodeTopoesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetNodeTopoesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNodeTopoesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeTopoesConnectionInternalServerError creates a GetNodeTopoesConnectionInternalServerError with default headers values
func NewGetNodeTopoesConnectionInternalServerError() *GetNodeTopoesConnectionInternalServerError {
	return &GetNodeTopoesConnectionInternalServerError{}
}

/* GetNodeTopoesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetNodeTopoesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetNodeTopoesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-node-topoes-connection][%d] getNodeTopoesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNodeTopoesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNodeTopoesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
