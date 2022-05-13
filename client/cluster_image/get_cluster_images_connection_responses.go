// Code generated by go-swagger; DO NOT EDIT.

package cluster_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetClusterImagesConnectionReader is a Reader for the GetClusterImagesConnection structure.
type GetClusterImagesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterImagesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterImagesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterImagesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterImagesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterImagesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterImagesConnectionOK creates a GetClusterImagesConnectionOK with default headers values
func NewGetClusterImagesConnectionOK() *GetClusterImagesConnectionOK {
	return &GetClusterImagesConnectionOK{}
}

/* GetClusterImagesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterImagesConnectionOK struct {
	Payload *models.ClusterImageConnection
}

func (o *GetClusterImagesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images-connection][%d] getClusterImagesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClusterImagesConnectionOK) GetPayload() *models.ClusterImageConnection {
	return o.Payload
}

func (o *GetClusterImagesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterImageConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterImagesConnectionBadRequest creates a GetClusterImagesConnectionBadRequest with default headers values
func NewGetClusterImagesConnectionBadRequest() *GetClusterImagesConnectionBadRequest {
	return &GetClusterImagesConnectionBadRequest{}
}

/* GetClusterImagesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetClusterImagesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetClusterImagesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images-connection][%d] getClusterImagesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterImagesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterImagesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterImagesConnectionNotFound creates a GetClusterImagesConnectionNotFound with default headers values
func NewGetClusterImagesConnectionNotFound() *GetClusterImagesConnectionNotFound {
	return &GetClusterImagesConnectionNotFound{}
}

/* GetClusterImagesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetClusterImagesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetClusterImagesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images-connection][%d] getClusterImagesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetClusterImagesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterImagesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterImagesConnectionInternalServerError creates a GetClusterImagesConnectionInternalServerError with default headers values
func NewGetClusterImagesConnectionInternalServerError() *GetClusterImagesConnectionInternalServerError {
	return &GetClusterImagesConnectionInternalServerError{}
}

/* GetClusterImagesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetClusterImagesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetClusterImagesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images-connection][%d] getClusterImagesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetClusterImagesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterImagesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
