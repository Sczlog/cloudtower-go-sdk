// Code generated by go-swagger; DO NOT EDIT.

package iscsi_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetIscsiTargetsReader is a Reader for the GetIscsiTargets structure.
type GetIscsiTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIscsiTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIscsiTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIscsiTargetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIscsiTargetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIscsiTargetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIscsiTargetsOK creates a GetIscsiTargetsOK with default headers values
func NewGetIscsiTargetsOK() *GetIscsiTargetsOK {
	return &GetIscsiTargetsOK{}
}

/* GetIscsiTargetsOK describes a response with status code 200, with default header values.

GetIscsiTargetsOK get iscsi targets o k
*/
type GetIscsiTargetsOK struct {
	XTowerRequestID string

	Payload []*models.IscsiTarget
}

func (o *GetIscsiTargetsOK) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsOK  %+v", 200, o.Payload)
}
func (o *GetIscsiTargetsOK) GetPayload() []*models.IscsiTarget {
	return o.Payload
}

func (o *GetIscsiTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIscsiTargetsBadRequest creates a GetIscsiTargetsBadRequest with default headers values
func NewGetIscsiTargetsBadRequest() *GetIscsiTargetsBadRequest {
	return &GetIscsiTargetsBadRequest{}
}

/* GetIscsiTargetsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIscsiTargetsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetIscsiTargetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsBadRequest  %+v", 400, o.Payload)
}
func (o *GetIscsiTargetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIscsiTargetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIscsiTargetsNotFound creates a GetIscsiTargetsNotFound with default headers values
func NewGetIscsiTargetsNotFound() *GetIscsiTargetsNotFound {
	return &GetIscsiTargetsNotFound{}
}

/* GetIscsiTargetsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetIscsiTargetsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetIscsiTargetsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsNotFound  %+v", 404, o.Payload)
}
func (o *GetIscsiTargetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIscsiTargetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIscsiTargetsInternalServerError creates a GetIscsiTargetsInternalServerError with default headers values
func NewGetIscsiTargetsInternalServerError() *GetIscsiTargetsInternalServerError {
	return &GetIscsiTargetsInternalServerError{}
}

/* GetIscsiTargetsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetIscsiTargetsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetIscsiTargetsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetIscsiTargetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIscsiTargetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
