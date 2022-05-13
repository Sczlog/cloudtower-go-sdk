// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateIscsiLunReader is a Reader for the UpdateIscsiLun structure.
type UpdateIscsiLunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIscsiLunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIscsiLunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIscsiLunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateIscsiLunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIscsiLunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateIscsiLunOK creates a UpdateIscsiLunOK with default headers values
func NewUpdateIscsiLunOK() *UpdateIscsiLunOK {
	return &UpdateIscsiLunOK{}
}

/* UpdateIscsiLunOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateIscsiLunOK struct {
	Payload []*models.WithTaskIscsiLun
}

func (o *UpdateIscsiLunOK) Error() string {
	return fmt.Sprintf("[POST /update-iscsi-lun][%d] updateIscsiLunOK  %+v", 200, o.Payload)
}
func (o *UpdateIscsiLunOK) GetPayload() []*models.WithTaskIscsiLun {
	return o.Payload
}

func (o *UpdateIscsiLunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIscsiLunBadRequest creates a UpdateIscsiLunBadRequest with default headers values
func NewUpdateIscsiLunBadRequest() *UpdateIscsiLunBadRequest {
	return &UpdateIscsiLunBadRequest{}
}

/* UpdateIscsiLunBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateIscsiLunBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateIscsiLunBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-iscsi-lun][%d] updateIscsiLunBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateIscsiLunBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateIscsiLunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIscsiLunNotFound creates a UpdateIscsiLunNotFound with default headers values
func NewUpdateIscsiLunNotFound() *UpdateIscsiLunNotFound {
	return &UpdateIscsiLunNotFound{}
}

/* UpdateIscsiLunNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateIscsiLunNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateIscsiLunNotFound) Error() string {
	return fmt.Sprintf("[POST /update-iscsi-lun][%d] updateIscsiLunNotFound  %+v", 404, o.Payload)
}
func (o *UpdateIscsiLunNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateIscsiLunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIscsiLunInternalServerError creates a UpdateIscsiLunInternalServerError with default headers values
func NewUpdateIscsiLunInternalServerError() *UpdateIscsiLunInternalServerError {
	return &UpdateIscsiLunInternalServerError{}
}

/* UpdateIscsiLunInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateIscsiLunInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateIscsiLunInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-iscsi-lun][%d] updateIscsiLunInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateIscsiLunInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateIscsiLunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
