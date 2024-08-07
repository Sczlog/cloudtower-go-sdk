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

// UpdateVirtualPrivateCloudSubnetReader is a Reader for the UpdateVirtualPrivateCloudSubnet structure.
type UpdateVirtualPrivateCloudSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVirtualPrivateCloudSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVirtualPrivateCloudSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVirtualPrivateCloudSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVirtualPrivateCloudSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVirtualPrivateCloudSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVirtualPrivateCloudSubnetOK creates a UpdateVirtualPrivateCloudSubnetOK with default headers values
func NewUpdateVirtualPrivateCloudSubnetOK() *UpdateVirtualPrivateCloudSubnetOK {
	return &UpdateVirtualPrivateCloudSubnetOK{}
}

/* UpdateVirtualPrivateCloudSubnetOK describes a response with status code 200, with default header values.

UpdateVirtualPrivateCloudSubnetOK update virtual private cloud subnet o k
*/
type UpdateVirtualPrivateCloudSubnetOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudSubnet
}

func (o *UpdateVirtualPrivateCloudSubnetOK) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-subnet][%d] updateVirtualPrivateCloudSubnetOK  %+v", 200, o.Payload)
}
func (o *UpdateVirtualPrivateCloudSubnetOK) GetPayload() []*models.WithTaskVirtualPrivateCloudSubnet {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudSubnetBadRequest creates a UpdateVirtualPrivateCloudSubnetBadRequest with default headers values
func NewUpdateVirtualPrivateCloudSubnetBadRequest() *UpdateVirtualPrivateCloudSubnetBadRequest {
	return &UpdateVirtualPrivateCloudSubnetBadRequest{}
}

/* UpdateVirtualPrivateCloudSubnetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVirtualPrivateCloudSubnetBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudSubnetBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-subnet][%d] updateVirtualPrivateCloudSubnetBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVirtualPrivateCloudSubnetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudSubnetNotFound creates a UpdateVirtualPrivateCloudSubnetNotFound with default headers values
func NewUpdateVirtualPrivateCloudSubnetNotFound() *UpdateVirtualPrivateCloudSubnetNotFound {
	return &UpdateVirtualPrivateCloudSubnetNotFound{}
}

/* UpdateVirtualPrivateCloudSubnetNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVirtualPrivateCloudSubnetNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudSubnetNotFound) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-subnet][%d] updateVirtualPrivateCloudSubnetNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVirtualPrivateCloudSubnetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudSubnetInternalServerError creates a UpdateVirtualPrivateCloudSubnetInternalServerError with default headers values
func NewUpdateVirtualPrivateCloudSubnetInternalServerError() *UpdateVirtualPrivateCloudSubnetInternalServerError {
	return &UpdateVirtualPrivateCloudSubnetInternalServerError{}
}

/* UpdateVirtualPrivateCloudSubnetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVirtualPrivateCloudSubnetInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-subnet][%d] updateVirtualPrivateCloudSubnetInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVirtualPrivateCloudSubnetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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