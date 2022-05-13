// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateSnapshotPlanReader is a Reader for the CreateSnapshotPlan structure.
type CreateSnapshotPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSnapshotPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSnapshotPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSnapshotPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateSnapshotPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSnapshotPlanOK creates a CreateSnapshotPlanOK with default headers values
func NewCreateSnapshotPlanOK() *CreateSnapshotPlanOK {
	return &CreateSnapshotPlanOK{}
}

/* CreateSnapshotPlanOK describes a response with status code 200, with default header values.

Ok
*/
type CreateSnapshotPlanOK struct {
	Payload []*models.WithTaskSnapshotPlan
}

func (o *CreateSnapshotPlanOK) Error() string {
	return fmt.Sprintf("[POST /create-snapshot-plan][%d] createSnapshotPlanOK  %+v", 200, o.Payload)
}
func (o *CreateSnapshotPlanOK) GetPayload() []*models.WithTaskSnapshotPlan {
	return o.Payload
}

func (o *CreateSnapshotPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotPlanBadRequest creates a CreateSnapshotPlanBadRequest with default headers values
func NewCreateSnapshotPlanBadRequest() *CreateSnapshotPlanBadRequest {
	return &CreateSnapshotPlanBadRequest{}
}

/* CreateSnapshotPlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSnapshotPlanBadRequest struct {
	Payload *models.ErrorBody
}

func (o *CreateSnapshotPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-snapshot-plan][%d] createSnapshotPlanBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSnapshotPlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSnapshotPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotPlanNotFound creates a CreateSnapshotPlanNotFound with default headers values
func NewCreateSnapshotPlanNotFound() *CreateSnapshotPlanNotFound {
	return &CreateSnapshotPlanNotFound{}
}

/* CreateSnapshotPlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateSnapshotPlanNotFound struct {
	Payload *models.ErrorBody
}

func (o *CreateSnapshotPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /create-snapshot-plan][%d] createSnapshotPlanNotFound  %+v", 404, o.Payload)
}
func (o *CreateSnapshotPlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSnapshotPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotPlanInternalServerError creates a CreateSnapshotPlanInternalServerError with default headers values
func NewCreateSnapshotPlanInternalServerError() *CreateSnapshotPlanInternalServerError {
	return &CreateSnapshotPlanInternalServerError{}
}

/* CreateSnapshotPlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateSnapshotPlanInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *CreateSnapshotPlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-snapshot-plan][%d] createSnapshotPlanInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateSnapshotPlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSnapshotPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
