// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVirtualPrivateCloudSubnetReader is a Reader for the CreateVirtualPrivateCloudSubnet structure.
type CreateVirtualPrivateCloudSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualPrivateCloudSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVirtualPrivateCloudSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVirtualPrivateCloudSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVirtualPrivateCloudSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVirtualPrivateCloudSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualPrivateCloudSubnetOK creates a CreateVirtualPrivateCloudSubnetOK with default headers values
func NewCreateVirtualPrivateCloudSubnetOK() *CreateVirtualPrivateCloudSubnetOK {
	return &CreateVirtualPrivateCloudSubnetOK{}
}

/* CreateVirtualPrivateCloudSubnetOK describes a response with status code 200, with default header values.

CreateVirtualPrivateCloudSubnetOK create virtual private cloud subnet o k
*/
type CreateVirtualPrivateCloudSubnetOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudSubnet
}

func (o *CreateVirtualPrivateCloudSubnetOK) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-subnet][%d] createVirtualPrivateCloudSubnetOK  %+v", 200, o.Payload)
}
func (o *CreateVirtualPrivateCloudSubnetOK) GetPayload() []*models.WithTaskVirtualPrivateCloudSubnet {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualPrivateCloudSubnetBadRequest creates a CreateVirtualPrivateCloudSubnetBadRequest with default headers values
func NewCreateVirtualPrivateCloudSubnetBadRequest() *CreateVirtualPrivateCloudSubnetBadRequest {
	return &CreateVirtualPrivateCloudSubnetBadRequest{}
}

/* CreateVirtualPrivateCloudSubnetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVirtualPrivateCloudSubnetBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSubnetBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-subnet][%d] createVirtualPrivateCloudSubnetBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualPrivateCloudSubnetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualPrivateCloudSubnetNotFound creates a CreateVirtualPrivateCloudSubnetNotFound with default headers values
func NewCreateVirtualPrivateCloudSubnetNotFound() *CreateVirtualPrivateCloudSubnetNotFound {
	return &CreateVirtualPrivateCloudSubnetNotFound{}
}

/* CreateVirtualPrivateCloudSubnetNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVirtualPrivateCloudSubnetNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSubnetNotFound) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-subnet][%d] createVirtualPrivateCloudSubnetNotFound  %+v", 404, o.Payload)
}
func (o *CreateVirtualPrivateCloudSubnetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualPrivateCloudSubnetInternalServerError creates a CreateVirtualPrivateCloudSubnetInternalServerError with default headers values
func NewCreateVirtualPrivateCloudSubnetInternalServerError() *CreateVirtualPrivateCloudSubnetInternalServerError {
	return &CreateVirtualPrivateCloudSubnetInternalServerError{}
}

/* CreateVirtualPrivateCloudSubnetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVirtualPrivateCloudSubnetInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-subnet][%d] createVirtualPrivateCloudSubnetInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVirtualPrivateCloudSubnetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}