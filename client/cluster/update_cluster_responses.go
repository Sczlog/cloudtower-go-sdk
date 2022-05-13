// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/* UpdateClusterOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterOK struct {
	Payload []*models.WithTaskCluster
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterBadRequest creates a UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {
	return &UpdateClusterBadRequest{}
}

/* UpdateClusterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateClusterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterNotFound creates a UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

/* UpdateClusterNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateClusterNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterNotFound  %+v", 404, o.Payload)
}
func (o *UpdateClusterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInternalServerError creates a UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {
	return &UpdateClusterInternalServerError{}
}

/* UpdateClusterInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateClusterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateClusterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
