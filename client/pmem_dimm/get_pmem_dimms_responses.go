// Code generated by go-swagger; DO NOT EDIT.

package pmem_dimm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetPmemDimmsReader is a Reader for the GetPmemDimms structure.
type GetPmemDimmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPmemDimmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPmemDimmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPmemDimmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPmemDimmsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPmemDimmsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPmemDimmsOK creates a GetPmemDimmsOK with default headers values
func NewGetPmemDimmsOK() *GetPmemDimmsOK {
	return &GetPmemDimmsOK{}
}

/* GetPmemDimmsOK describes a response with status code 200, with default header values.

Ok
*/
type GetPmemDimmsOK struct {
	Payload []*models.PmemDimm
}

func (o *GetPmemDimmsOK) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsOK  %+v", 200, o.Payload)
}
func (o *GetPmemDimmsOK) GetPayload() []*models.PmemDimm {
	return o.Payload
}

func (o *GetPmemDimmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPmemDimmsBadRequest creates a GetPmemDimmsBadRequest with default headers values
func NewGetPmemDimmsBadRequest() *GetPmemDimmsBadRequest {
	return &GetPmemDimmsBadRequest{}
}

/* GetPmemDimmsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetPmemDimmsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetPmemDimmsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsBadRequest  %+v", 400, o.Payload)
}
func (o *GetPmemDimmsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPmemDimmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPmemDimmsNotFound creates a GetPmemDimmsNotFound with default headers values
func NewGetPmemDimmsNotFound() *GetPmemDimmsNotFound {
	return &GetPmemDimmsNotFound{}
}

/* GetPmemDimmsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPmemDimmsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetPmemDimmsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsNotFound  %+v", 404, o.Payload)
}
func (o *GetPmemDimmsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPmemDimmsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPmemDimmsInternalServerError creates a GetPmemDimmsInternalServerError with default headers values
func NewGetPmemDimmsInternalServerError() *GetPmemDimmsInternalServerError {
	return &GetPmemDimmsInternalServerError{}
}

/* GetPmemDimmsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetPmemDimmsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetPmemDimmsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPmemDimmsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPmemDimmsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
