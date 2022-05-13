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

// GetVMPlacementGroupsConnectionReader is a Reader for the GetVMPlacementGroupsConnection structure.
type GetVMPlacementGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMPlacementGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMPlacementGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMPlacementGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMPlacementGroupsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVMPlacementGroupsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMPlacementGroupsConnectionOK creates a GetVMPlacementGroupsConnectionOK with default headers values
func NewGetVMPlacementGroupsConnectionOK() *GetVMPlacementGroupsConnectionOK {
	return &GetVMPlacementGroupsConnectionOK{}
}

/* GetVMPlacementGroupsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMPlacementGroupsConnectionOK struct {
	Payload *models.VMPlacementGroupConnection
}

func (o *GetVMPlacementGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionOK) GetPayload() *models.VMPlacementGroupConnection {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMPlacementGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMPlacementGroupsConnectionBadRequest creates a GetVMPlacementGroupsConnectionBadRequest with default headers values
func NewGetVMPlacementGroupsConnectionBadRequest() *GetVMPlacementGroupsConnectionBadRequest {
	return &GetVMPlacementGroupsConnectionBadRequest{}
}

/* GetVMPlacementGroupsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVMPlacementGroupsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVMPlacementGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMPlacementGroupsConnectionNotFound creates a GetVMPlacementGroupsConnectionNotFound with default headers values
func NewGetVMPlacementGroupsConnectionNotFound() *GetVMPlacementGroupsConnectionNotFound {
	return &GetVMPlacementGroupsConnectionNotFound{}
}

/* GetVMPlacementGroupsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVMPlacementGroupsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVMPlacementGroupsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMPlacementGroupsConnectionInternalServerError creates a GetVMPlacementGroupsConnectionInternalServerError with default headers values
func NewGetVMPlacementGroupsConnectionInternalServerError() *GetVMPlacementGroupsConnectionInternalServerError {
	return &GetVMPlacementGroupsConnectionInternalServerError{}
}

/* GetVMPlacementGroupsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVMPlacementGroupsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVMPlacementGroupsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
