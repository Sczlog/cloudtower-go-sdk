// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// CreateVMPlacementGroupReader is a Reader for the CreateVMPlacementGroup structure.
type CreateVMPlacementGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMPlacementGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMPlacementGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMPlacementGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMPlacementGroupOK creates a CreateVMPlacementGroupOK with default headers values
func NewCreateVMPlacementGroupOK() *CreateVMPlacementGroupOK {
	return &CreateVMPlacementGroupOK{}
}

/* CreateVMPlacementGroupOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVMPlacementGroupOK struct {
	Payload []*models.WithTaskVMPlacementGroup
}

func (o *CreateVMPlacementGroupOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-placement-group][%d] createVmPlacementGroupOK  %+v", 200, o.Payload)
}
func (o *CreateVMPlacementGroupOK) GetPayload() []*models.WithTaskVMPlacementGroup {
	return o.Payload
}

func (o *CreateVMPlacementGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVMPlacementGroupBadRequest creates a CreateVMPlacementGroupBadRequest with default headers values
func NewCreateVMPlacementGroupBadRequest() *CreateVMPlacementGroupBadRequest {
	return &CreateVMPlacementGroupBadRequest{}
}

/* CreateVMPlacementGroupBadRequest describes a response with status code 400, with default header values.

CreateVMPlacementGroupBadRequest create Vm placement group bad request
*/
type CreateVMPlacementGroupBadRequest struct {
	Payload string
}

func (o *CreateVMPlacementGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-placement-group][%d] createVmPlacementGroupBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMPlacementGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVMPlacementGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
