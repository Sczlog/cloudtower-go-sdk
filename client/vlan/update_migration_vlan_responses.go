// Code generated by go-swagger; DO NOT EDIT.

package vlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateMigrationVlanReader is a Reader for the UpdateMigrationVlan structure.
type UpdateMigrationVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMigrationVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMigrationVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMigrationVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMigrationVlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMigrationVlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateMigrationVlanOK creates a UpdateMigrationVlanOK with default headers values
func NewUpdateMigrationVlanOK() *UpdateMigrationVlanOK {
	return &UpdateMigrationVlanOK{}
}

/* UpdateMigrationVlanOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateMigrationVlanOK struct {
	Payload []*models.WithTaskVlan
}

func (o *UpdateMigrationVlanOK) Error() string {
	return fmt.Sprintf("[POST /update-migration-vlan][%d] updateMigrationVlanOK  %+v", 200, o.Payload)
}
func (o *UpdateMigrationVlanOK) GetPayload() []*models.WithTaskVlan {
	return o.Payload
}

func (o *UpdateMigrationVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMigrationVlanBadRequest creates a UpdateMigrationVlanBadRequest with default headers values
func NewUpdateMigrationVlanBadRequest() *UpdateMigrationVlanBadRequest {
	return &UpdateMigrationVlanBadRequest{}
}

/* UpdateMigrationVlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateMigrationVlanBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateMigrationVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-migration-vlan][%d] updateMigrationVlanBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateMigrationVlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateMigrationVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMigrationVlanNotFound creates a UpdateMigrationVlanNotFound with default headers values
func NewUpdateMigrationVlanNotFound() *UpdateMigrationVlanNotFound {
	return &UpdateMigrationVlanNotFound{}
}

/* UpdateMigrationVlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateMigrationVlanNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateMigrationVlanNotFound) Error() string {
	return fmt.Sprintf("[POST /update-migration-vlan][%d] updateMigrationVlanNotFound  %+v", 404, o.Payload)
}
func (o *UpdateMigrationVlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateMigrationVlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMigrationVlanInternalServerError creates a UpdateMigrationVlanInternalServerError with default headers values
func NewUpdateMigrationVlanInternalServerError() *UpdateMigrationVlanInternalServerError {
	return &UpdateMigrationVlanInternalServerError{}
}

/* UpdateMigrationVlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateMigrationVlanInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateMigrationVlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-migration-vlan][%d] updateMigrationVlanInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateMigrationVlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateMigrationVlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
