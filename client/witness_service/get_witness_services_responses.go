// Code generated by go-swagger; DO NOT EDIT.

package witness_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetWitnessServicesReader is a Reader for the GetWitnessServices structure.
type GetWitnessServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWitnessServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWitnessServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWitnessServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWitnessServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWitnessServicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWitnessServicesOK creates a GetWitnessServicesOK with default headers values
func NewGetWitnessServicesOK() *GetWitnessServicesOK {
	return &GetWitnessServicesOK{}
}

/* GetWitnessServicesOK describes a response with status code 200, with default header values.

Ok
*/
type GetWitnessServicesOK struct {
	Payload []*models.WitnessService
}

func (o *GetWitnessServicesOK) Error() string {
	return fmt.Sprintf("[POST /get-witness-services][%d] getWitnessServicesOK  %+v", 200, o.Payload)
}
func (o *GetWitnessServicesOK) GetPayload() []*models.WitnessService {
	return o.Payload
}

func (o *GetWitnessServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessServicesBadRequest creates a GetWitnessServicesBadRequest with default headers values
func NewGetWitnessServicesBadRequest() *GetWitnessServicesBadRequest {
	return &GetWitnessServicesBadRequest{}
}

/* GetWitnessServicesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWitnessServicesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-witness-services][%d] getWitnessServicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetWitnessServicesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessServicesNotFound creates a GetWitnessServicesNotFound with default headers values
func NewGetWitnessServicesNotFound() *GetWitnessServicesNotFound {
	return &GetWitnessServicesNotFound{}
}

/* GetWitnessServicesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetWitnessServicesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessServicesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-witness-services][%d] getWitnessServicesNotFound  %+v", 404, o.Payload)
}
func (o *GetWitnessServicesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessServicesInternalServerError creates a GetWitnessServicesInternalServerError with default headers values
func NewGetWitnessServicesInternalServerError() *GetWitnessServicesInternalServerError {
	return &GetWitnessServicesInternalServerError{}
}

/* GetWitnessServicesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetWitnessServicesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessServicesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-witness-services][%d] getWitnessServicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWitnessServicesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessServicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
