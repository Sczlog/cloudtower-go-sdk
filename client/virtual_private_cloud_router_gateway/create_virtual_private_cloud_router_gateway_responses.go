// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_router_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVirtualPrivateCloudRouterGatewayReader is a Reader for the CreateVirtualPrivateCloudRouterGateway structure.
type CreateVirtualPrivateCloudRouterGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualPrivateCloudRouterGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVirtualPrivateCloudRouterGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVirtualPrivateCloudRouterGatewayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVirtualPrivateCloudRouterGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVirtualPrivateCloudRouterGatewayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualPrivateCloudRouterGatewayOK creates a CreateVirtualPrivateCloudRouterGatewayOK with default headers values
func NewCreateVirtualPrivateCloudRouterGatewayOK() *CreateVirtualPrivateCloudRouterGatewayOK {
	return &CreateVirtualPrivateCloudRouterGatewayOK{}
}

/* CreateVirtualPrivateCloudRouterGatewayOK describes a response with status code 200, with default header values.

CreateVirtualPrivateCloudRouterGatewayOK create virtual private cloud router gateway o k
*/
type CreateVirtualPrivateCloudRouterGatewayOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudRouterGateway
}

func (o *CreateVirtualPrivateCloudRouterGatewayOK) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-router-gateway][%d] createVirtualPrivateCloudRouterGatewayOK  %+v", 200, o.Payload)
}
func (o *CreateVirtualPrivateCloudRouterGatewayOK) GetPayload() []*models.WithTaskVirtualPrivateCloudRouterGateway {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudRouterGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudRouterGatewayBadRequest creates a CreateVirtualPrivateCloudRouterGatewayBadRequest with default headers values
func NewCreateVirtualPrivateCloudRouterGatewayBadRequest() *CreateVirtualPrivateCloudRouterGatewayBadRequest {
	return &CreateVirtualPrivateCloudRouterGatewayBadRequest{}
}

/* CreateVirtualPrivateCloudRouterGatewayBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVirtualPrivateCloudRouterGatewayBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudRouterGatewayBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-router-gateway][%d] createVirtualPrivateCloudRouterGatewayBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualPrivateCloudRouterGatewayBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudRouterGatewayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudRouterGatewayNotFound creates a CreateVirtualPrivateCloudRouterGatewayNotFound with default headers values
func NewCreateVirtualPrivateCloudRouterGatewayNotFound() *CreateVirtualPrivateCloudRouterGatewayNotFound {
	return &CreateVirtualPrivateCloudRouterGatewayNotFound{}
}

/* CreateVirtualPrivateCloudRouterGatewayNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVirtualPrivateCloudRouterGatewayNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudRouterGatewayNotFound) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-router-gateway][%d] createVirtualPrivateCloudRouterGatewayNotFound  %+v", 404, o.Payload)
}
func (o *CreateVirtualPrivateCloudRouterGatewayNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudRouterGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudRouterGatewayInternalServerError creates a CreateVirtualPrivateCloudRouterGatewayInternalServerError with default headers values
func NewCreateVirtualPrivateCloudRouterGatewayInternalServerError() *CreateVirtualPrivateCloudRouterGatewayInternalServerError {
	return &CreateVirtualPrivateCloudRouterGatewayInternalServerError{}
}

/* CreateVirtualPrivateCloudRouterGatewayInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVirtualPrivateCloudRouterGatewayInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudRouterGatewayInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-router-gateway][%d] createVirtualPrivateCloudRouterGatewayInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVirtualPrivateCloudRouterGatewayInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudRouterGatewayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
