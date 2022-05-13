// Code generated by go-swagger; DO NOT EDIT.

package witness

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetWitnessesReader is a Reader for the GetWitnesses structure.
type GetWitnessesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWitnessesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWitnessesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWitnessesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWitnessesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWitnessesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWitnessesOK creates a GetWitnessesOK with default headers values
func NewGetWitnessesOK() *GetWitnessesOK {
	return &GetWitnessesOK{}
}

/* GetWitnessesOK describes a response with status code 200, with default header values.

Ok
*/
type GetWitnessesOK struct {
	Payload []*models.Witness
}

func (o *GetWitnessesOK) Error() string {
	return fmt.Sprintf("[POST /get-witnesses][%d] getWitnessesOK  %+v", 200, o.Payload)
}
func (o *GetWitnessesOK) GetPayload() []*models.Witness {
	return o.Payload
}

func (o *GetWitnessesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessesBadRequest creates a GetWitnessesBadRequest with default headers values
func NewGetWitnessesBadRequest() *GetWitnessesBadRequest {
	return &GetWitnessesBadRequest{}
}

/* GetWitnessesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWitnessesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-witnesses][%d] getWitnessesBadRequest  %+v", 400, o.Payload)
}
func (o *GetWitnessesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessesNotFound creates a GetWitnessesNotFound with default headers values
func NewGetWitnessesNotFound() *GetWitnessesNotFound {
	return &GetWitnessesNotFound{}
}

/* GetWitnessesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetWitnessesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-witnesses][%d] getWitnessesNotFound  %+v", 404, o.Payload)
}
func (o *GetWitnessesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWitnessesInternalServerError creates a GetWitnessesInternalServerError with default headers values
func NewGetWitnessesInternalServerError() *GetWitnessesInternalServerError {
	return &GetWitnessesInternalServerError{}
}

/* GetWitnessesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetWitnessesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWitnessesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-witnesses][%d] getWitnessesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWitnessesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWitnessesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
