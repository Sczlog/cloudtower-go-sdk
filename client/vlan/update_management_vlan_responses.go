// Code generated by go-swagger; DO NOT EDIT.

package vlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateManagementVlanReader is a Reader for the UpdateManagementVlan structure.
type UpdateManagementVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateManagementVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateManagementVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateManagementVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateManagementVlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateManagementVlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateManagementVlanOK creates a UpdateManagementVlanOK with default headers values
func NewUpdateManagementVlanOK() *UpdateManagementVlanOK {
	return &UpdateManagementVlanOK{}
}

/* UpdateManagementVlanOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateManagementVlanOK struct {
	Payload []*models.WithTaskVlan
}

func (o *UpdateManagementVlanOK) Error() string {
	return fmt.Sprintf("[POST /update-management-vlan][%d] updateManagementVlanOK  %+v", 200, o.Payload)
}
func (o *UpdateManagementVlanOK) GetPayload() []*models.WithTaskVlan {
	return o.Payload
}

func (o *UpdateManagementVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateManagementVlanBadRequest creates a UpdateManagementVlanBadRequest with default headers values
func NewUpdateManagementVlanBadRequest() *UpdateManagementVlanBadRequest {
	return &UpdateManagementVlanBadRequest{}
}

/* UpdateManagementVlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateManagementVlanBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateManagementVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-management-vlan][%d] updateManagementVlanBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateManagementVlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManagementVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateManagementVlanNotFound creates a UpdateManagementVlanNotFound with default headers values
func NewUpdateManagementVlanNotFound() *UpdateManagementVlanNotFound {
	return &UpdateManagementVlanNotFound{}
}

/* UpdateManagementVlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateManagementVlanNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateManagementVlanNotFound) Error() string {
	return fmt.Sprintf("[POST /update-management-vlan][%d] updateManagementVlanNotFound  %+v", 404, o.Payload)
}
func (o *UpdateManagementVlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManagementVlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateManagementVlanInternalServerError creates a UpdateManagementVlanInternalServerError with default headers values
func NewUpdateManagementVlanInternalServerError() *UpdateManagementVlanInternalServerError {
	return &UpdateManagementVlanInternalServerError{}
}

/* UpdateManagementVlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateManagementVlanInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateManagementVlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-management-vlan][%d] updateManagementVlanInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateManagementVlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManagementVlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
