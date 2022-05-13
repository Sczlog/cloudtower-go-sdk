// Code generated by go-swagger; DO NOT EDIT.

package everoute_license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetEverouteLicensesReader is a Reader for the GetEverouteLicenses structure.
type GetEverouteLicensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEverouteLicensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEverouteLicensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEverouteLicensesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEverouteLicensesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEverouteLicensesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEverouteLicensesOK creates a GetEverouteLicensesOK with default headers values
func NewGetEverouteLicensesOK() *GetEverouteLicensesOK {
	return &GetEverouteLicensesOK{}
}

/* GetEverouteLicensesOK describes a response with status code 200, with default header values.

Ok
*/
type GetEverouteLicensesOK struct {
	Payload []*models.EverouteLicense
}

func (o *GetEverouteLicensesOK) Error() string {
	return fmt.Sprintf("[POST /get-everoute-licenses][%d] getEverouteLicensesOK  %+v", 200, o.Payload)
}
func (o *GetEverouteLicensesOK) GetPayload() []*models.EverouteLicense {
	return o.Payload
}

func (o *GetEverouteLicensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEverouteLicensesBadRequest creates a GetEverouteLicensesBadRequest with default headers values
func NewGetEverouteLicensesBadRequest() *GetEverouteLicensesBadRequest {
	return &GetEverouteLicensesBadRequest{}
}

/* GetEverouteLicensesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetEverouteLicensesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetEverouteLicensesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-everoute-licenses][%d] getEverouteLicensesBadRequest  %+v", 400, o.Payload)
}
func (o *GetEverouteLicensesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteLicensesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEverouteLicensesNotFound creates a GetEverouteLicensesNotFound with default headers values
func NewGetEverouteLicensesNotFound() *GetEverouteLicensesNotFound {
	return &GetEverouteLicensesNotFound{}
}

/* GetEverouteLicensesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEverouteLicensesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetEverouteLicensesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-everoute-licenses][%d] getEverouteLicensesNotFound  %+v", 404, o.Payload)
}
func (o *GetEverouteLicensesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteLicensesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEverouteLicensesInternalServerError creates a GetEverouteLicensesInternalServerError with default headers values
func NewGetEverouteLicensesInternalServerError() *GetEverouteLicensesInternalServerError {
	return &GetEverouteLicensesInternalServerError{}
}

/* GetEverouteLicensesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetEverouteLicensesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetEverouteLicensesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-everoute-licenses][%d] getEverouteLicensesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEverouteLicensesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteLicensesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
