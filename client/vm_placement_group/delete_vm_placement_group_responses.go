// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteVMPlacementGroupReader is a Reader for the DeleteVMPlacementGroup structure.
type DeleteVMPlacementGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMPlacementGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMPlacementGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMPlacementGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMPlacementGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMPlacementGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMPlacementGroupOK creates a DeleteVMPlacementGroupOK with default headers values
func NewDeleteVMPlacementGroupOK() *DeleteVMPlacementGroupOK {
	return &DeleteVMPlacementGroupOK{}
}

/* DeleteVMPlacementGroupOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMPlacementGroupOK struct {
	Payload []*models.WithTaskDeleteVMPlacementGroup
}

func (o *DeleteVMPlacementGroupOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-placement-group][%d] deleteVmPlacementGroupOK  %+v", 200, o.Payload)
}
func (o *DeleteVMPlacementGroupOK) GetPayload() []*models.WithTaskDeleteVMPlacementGroup {
	return o.Payload
}

func (o *DeleteVMPlacementGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMPlacementGroupBadRequest creates a DeleteVMPlacementGroupBadRequest with default headers values
func NewDeleteVMPlacementGroupBadRequest() *DeleteVMPlacementGroupBadRequest {
	return &DeleteVMPlacementGroupBadRequest{}
}

/* DeleteVMPlacementGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMPlacementGroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMPlacementGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-placement-group][%d] deleteVmPlacementGroupBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMPlacementGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMPlacementGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMPlacementGroupNotFound creates a DeleteVMPlacementGroupNotFound with default headers values
func NewDeleteVMPlacementGroupNotFound() *DeleteVMPlacementGroupNotFound {
	return &DeleteVMPlacementGroupNotFound{}
}

/* DeleteVMPlacementGroupNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMPlacementGroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMPlacementGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm-placement-group][%d] deleteVmPlacementGroupNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMPlacementGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMPlacementGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMPlacementGroupInternalServerError creates a DeleteVMPlacementGroupInternalServerError with default headers values
func NewDeleteVMPlacementGroupInternalServerError() *DeleteVMPlacementGroupInternalServerError {
	return &DeleteVMPlacementGroupInternalServerError{}
}

/* DeleteVMPlacementGroupInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMPlacementGroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteVMPlacementGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm-placement-group][%d] deleteVmPlacementGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMPlacementGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMPlacementGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
