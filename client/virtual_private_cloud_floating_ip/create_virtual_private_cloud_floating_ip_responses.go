// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_floating_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVirtualPrivateCloudFloatingIPReader is a Reader for the CreateVirtualPrivateCloudFloatingIP structure.
type CreateVirtualPrivateCloudFloatingIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualPrivateCloudFloatingIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVirtualPrivateCloudFloatingIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVirtualPrivateCloudFloatingIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVirtualPrivateCloudFloatingIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVirtualPrivateCloudFloatingIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualPrivateCloudFloatingIPOK creates a CreateVirtualPrivateCloudFloatingIPOK with default headers values
func NewCreateVirtualPrivateCloudFloatingIPOK() *CreateVirtualPrivateCloudFloatingIPOK {
	return &CreateVirtualPrivateCloudFloatingIPOK{}
}

/* CreateVirtualPrivateCloudFloatingIPOK describes a response with status code 200, with default header values.

CreateVirtualPrivateCloudFloatingIPOK create virtual private cloud floating Ip o k
*/
type CreateVirtualPrivateCloudFloatingIPOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudFloatingIP
}

func (o *CreateVirtualPrivateCloudFloatingIPOK) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-floating-ip][%d] createVirtualPrivateCloudFloatingIpOK  %+v", 200, o.Payload)
}
func (o *CreateVirtualPrivateCloudFloatingIPOK) GetPayload() []*models.WithTaskVirtualPrivateCloudFloatingIP {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudFloatingIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudFloatingIPBadRequest creates a CreateVirtualPrivateCloudFloatingIPBadRequest with default headers values
func NewCreateVirtualPrivateCloudFloatingIPBadRequest() *CreateVirtualPrivateCloudFloatingIPBadRequest {
	return &CreateVirtualPrivateCloudFloatingIPBadRequest{}
}

/* CreateVirtualPrivateCloudFloatingIPBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVirtualPrivateCloudFloatingIPBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudFloatingIPBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-floating-ip][%d] createVirtualPrivateCloudFloatingIpBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualPrivateCloudFloatingIPBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudFloatingIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudFloatingIPNotFound creates a CreateVirtualPrivateCloudFloatingIPNotFound with default headers values
func NewCreateVirtualPrivateCloudFloatingIPNotFound() *CreateVirtualPrivateCloudFloatingIPNotFound {
	return &CreateVirtualPrivateCloudFloatingIPNotFound{}
}

/* CreateVirtualPrivateCloudFloatingIPNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVirtualPrivateCloudFloatingIPNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudFloatingIPNotFound) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-floating-ip][%d] createVirtualPrivateCloudFloatingIpNotFound  %+v", 404, o.Payload)
}
func (o *CreateVirtualPrivateCloudFloatingIPNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudFloatingIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudFloatingIPInternalServerError creates a CreateVirtualPrivateCloudFloatingIPInternalServerError with default headers values
func NewCreateVirtualPrivateCloudFloatingIPInternalServerError() *CreateVirtualPrivateCloudFloatingIPInternalServerError {
	return &CreateVirtualPrivateCloudFloatingIPInternalServerError{}
}

/* CreateVirtualPrivateCloudFloatingIPInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVirtualPrivateCloudFloatingIPInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudFloatingIPInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-floating-ip][%d] createVirtualPrivateCloudFloatingIpInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVirtualPrivateCloudFloatingIPInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudFloatingIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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