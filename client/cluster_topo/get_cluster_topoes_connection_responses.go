// Code generated by go-swagger; DO NOT EDIT.

package cluster_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetClusterTopoesConnectionReader is a Reader for the GetClusterTopoesConnection structure.
type GetClusterTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterTopoesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterTopoesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterTopoesConnectionOK creates a GetClusterTopoesConnectionOK with default headers values
func NewGetClusterTopoesConnectionOK() *GetClusterTopoesConnectionOK {
	return &GetClusterTopoesConnectionOK{}
}

/* GetClusterTopoesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterTopoesConnectionOK struct {
	Payload *models.ClusterTopoConnection
}

func (o *GetClusterTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-topoes-connection][%d] getClusterTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClusterTopoesConnectionOK) GetPayload() *models.ClusterTopoConnection {
	return o.Payload
}

func (o *GetClusterTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTopoesConnectionBadRequest creates a GetClusterTopoesConnectionBadRequest with default headers values
func NewGetClusterTopoesConnectionBadRequest() *GetClusterTopoesConnectionBadRequest {
	return &GetClusterTopoesConnectionBadRequest{}
}

/* GetClusterTopoesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetClusterTopoesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetClusterTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-topoes-connection][%d] getClusterTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterTopoesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTopoesConnectionNotFound creates a GetClusterTopoesConnectionNotFound with default headers values
func NewGetClusterTopoesConnectionNotFound() *GetClusterTopoesConnectionNotFound {
	return &GetClusterTopoesConnectionNotFound{}
}

/* GetClusterTopoesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetClusterTopoesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetClusterTopoesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cluster-topoes-connection][%d] getClusterTopoesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetClusterTopoesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterTopoesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTopoesConnectionInternalServerError creates a GetClusterTopoesConnectionInternalServerError with default headers values
func NewGetClusterTopoesConnectionInternalServerError() *GetClusterTopoesConnectionInternalServerError {
	return &GetClusterTopoesConnectionInternalServerError{}
}

/* GetClusterTopoesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetClusterTopoesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetClusterTopoesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cluster-topoes-connection][%d] getClusterTopoesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetClusterTopoesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterTopoesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
