// Code generated by go-swagger; DO NOT EDIT.

package namespace_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetNamespaceGroupsReader is a Reader for the GetNamespaceGroups structure.
type GetNamespaceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNamespaceGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNamespaceGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNamespaceGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNamespaceGroupsOK creates a GetNamespaceGroupsOK with default headers values
func NewGetNamespaceGroupsOK() *GetNamespaceGroupsOK {
	return &GetNamespaceGroupsOK{}
}

/* GetNamespaceGroupsOK describes a response with status code 200, with default header values.

Ok
*/
type GetNamespaceGroupsOK struct {
	Payload []*models.NamespaceGroup
}

func (o *GetNamespaceGroupsOK) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups][%d] getNamespaceGroupsOK  %+v", 200, o.Payload)
}
func (o *GetNamespaceGroupsOK) GetPayload() []*models.NamespaceGroup {
	return o.Payload
}

func (o *GetNamespaceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceGroupsBadRequest creates a GetNamespaceGroupsBadRequest with default headers values
func NewGetNamespaceGroupsBadRequest() *GetNamespaceGroupsBadRequest {
	return &GetNamespaceGroupsBadRequest{}
}

/* GetNamespaceGroupsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetNamespaceGroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetNamespaceGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups][%d] getNamespaceGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *GetNamespaceGroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNamespaceGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceGroupsNotFound creates a GetNamespaceGroupsNotFound with default headers values
func NewGetNamespaceGroupsNotFound() *GetNamespaceGroupsNotFound {
	return &GetNamespaceGroupsNotFound{}
}

/* GetNamespaceGroupsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetNamespaceGroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetNamespaceGroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups][%d] getNamespaceGroupsNotFound  %+v", 404, o.Payload)
}
func (o *GetNamespaceGroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNamespaceGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceGroupsInternalServerError creates a GetNamespaceGroupsInternalServerError with default headers values
func NewGetNamespaceGroupsInternalServerError() *GetNamespaceGroupsInternalServerError {
	return &GetNamespaceGroupsInternalServerError{}
}

/* GetNamespaceGroupsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetNamespaceGroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetNamespaceGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups][%d] getNamespaceGroupsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNamespaceGroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNamespaceGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
