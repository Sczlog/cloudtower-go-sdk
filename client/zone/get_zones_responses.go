// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetZonesReader is a Reader for the GetZones structure.
type GetZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetZonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetZonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetZonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZonesOK creates a GetZonesOK with default headers values
func NewGetZonesOK() *GetZonesOK {
	return &GetZonesOK{}
}

/* GetZonesOK describes a response with status code 200, with default header values.

Ok
*/
type GetZonesOK struct {
	Payload []*models.Zone
}

func (o *GetZonesOK) Error() string {
	return fmt.Sprintf("[POST /get-zones][%d] getZonesOK  %+v", 200, o.Payload)
}
func (o *GetZonesOK) GetPayload() []*models.Zone {
	return o.Payload
}

func (o *GetZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesBadRequest creates a GetZonesBadRequest with default headers values
func NewGetZonesBadRequest() *GetZonesBadRequest {
	return &GetZonesBadRequest{}
}

/* GetZonesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetZonesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetZonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-zones][%d] getZonesBadRequest  %+v", 400, o.Payload)
}
func (o *GetZonesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesNotFound creates a GetZonesNotFound with default headers values
func NewGetZonesNotFound() *GetZonesNotFound {
	return &GetZonesNotFound{}
}

/* GetZonesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetZonesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetZonesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-zones][%d] getZonesNotFound  %+v", 404, o.Payload)
}
func (o *GetZonesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesInternalServerError creates a GetZonesInternalServerError with default headers values
func NewGetZonesInternalServerError() *GetZonesInternalServerError {
	return &GetZonesInternalServerError{}
}

/* GetZonesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetZonesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetZonesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-zones][%d] getZonesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetZonesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
