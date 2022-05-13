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

// DeleteIscsiLunReader is a Reader for the DeleteIscsiLun structure.
type DeleteIscsiLunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIscsiLunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIscsiLunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIscsiLunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIscsiLunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIscsiLunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIscsiLunOK creates a DeleteIscsiLunOK with default headers values
func NewDeleteIscsiLunOK() *DeleteIscsiLunOK {
	return &DeleteIscsiLunOK{}
}

/* DeleteIscsiLunOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteIscsiLunOK struct {
	Payload []*models.WithTaskDeleteIscsiLun
}

func (o *DeleteIscsiLunOK) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun][%d] deleteIscsiLunOK  %+v", 200, o.Payload)
}
func (o *DeleteIscsiLunOK) GetPayload() []*models.WithTaskDeleteIscsiLun {
	return o.Payload
}

func (o *DeleteIscsiLunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunBadRequest creates a DeleteIscsiLunBadRequest with default headers values
func NewDeleteIscsiLunBadRequest() *DeleteIscsiLunBadRequest {
	return &DeleteIscsiLunBadRequest{}
}

/* DeleteIscsiLunBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteIscsiLunBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun][%d] deleteIscsiLunBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteIscsiLunBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunNotFound creates a DeleteIscsiLunNotFound with default headers values
func NewDeleteIscsiLunNotFound() *DeleteIscsiLunNotFound {
	return &DeleteIscsiLunNotFound{}
}

/* DeleteIscsiLunNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteIscsiLunNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun][%d] deleteIscsiLunNotFound  %+v", 404, o.Payload)
}
func (o *DeleteIscsiLunNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIscsiLunInternalServerError creates a DeleteIscsiLunInternalServerError with default headers values
func NewDeleteIscsiLunInternalServerError() *DeleteIscsiLunInternalServerError {
	return &DeleteIscsiLunInternalServerError{}
}

/* DeleteIscsiLunInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteIscsiLunInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteIscsiLunInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-iscsi-lun][%d] deleteIscsiLunInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteIscsiLunInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIscsiLunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
