// Code generated by go-swagger; DO NOT EDIT.

package zone_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetZoneTopoesReader is a Reader for the GetZoneTopoes structure.
type GetZoneTopoesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneTopoesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZoneTopoesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetZoneTopoesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetZoneTopoesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetZoneTopoesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZoneTopoesOK creates a GetZoneTopoesOK with default headers values
func NewGetZoneTopoesOK() *GetZoneTopoesOK {
	return &GetZoneTopoesOK{}
}

/* GetZoneTopoesOK describes a response with status code 200, with default header values.

GetZoneTopoesOK get zone topoes o k
*/
type GetZoneTopoesOK struct {
	XTowerRequestID string

	Payload []*models.ZoneTopo
}

func (o *GetZoneTopoesOK) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes][%d] getZoneTopoesOK  %+v", 200, o.Payload)
}
func (o *GetZoneTopoesOK) GetPayload() []*models.ZoneTopo {
	return o.Payload
}

func (o *GetZoneTopoesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZoneTopoesBadRequest creates a GetZoneTopoesBadRequest with default headers values
func NewGetZoneTopoesBadRequest() *GetZoneTopoesBadRequest {
	return &GetZoneTopoesBadRequest{}
}

/* GetZoneTopoesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetZoneTopoesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes][%d] getZoneTopoesBadRequest  %+v", 400, o.Payload)
}
func (o *GetZoneTopoesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZoneTopoesNotFound creates a GetZoneTopoesNotFound with default headers values
func NewGetZoneTopoesNotFound() *GetZoneTopoesNotFound {
	return &GetZoneTopoesNotFound{}
}

/* GetZoneTopoesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetZoneTopoesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes][%d] getZoneTopoesNotFound  %+v", 404, o.Payload)
}
func (o *GetZoneTopoesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZoneTopoesInternalServerError creates a GetZoneTopoesInternalServerError with default headers values
func NewGetZoneTopoesInternalServerError() *GetZoneTopoesInternalServerError {
	return &GetZoneTopoesInternalServerError{}
}

/* GetZoneTopoesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetZoneTopoesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes][%d] getZoneTopoesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetZoneTopoesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
