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

// DeleteVirtualPrivateCloudSubnetReader is a Reader for the DeleteVirtualPrivateCloudSubnet structure.
type DeleteVirtualPrivateCloudSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVirtualPrivateCloudSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVirtualPrivateCloudSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVirtualPrivateCloudSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVirtualPrivateCloudSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVirtualPrivateCloudSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVirtualPrivateCloudSubnetOK creates a DeleteVirtualPrivateCloudSubnetOK with default headers values
func NewDeleteVirtualPrivateCloudSubnetOK() *DeleteVirtualPrivateCloudSubnetOK {
	return &DeleteVirtualPrivateCloudSubnetOK{}
}

/* DeleteVirtualPrivateCloudSubnetOK describes a response with status code 200, with default header values.

DeleteVirtualPrivateCloudSubnetOK delete virtual private cloud subnet o k
*/
type DeleteVirtualPrivateCloudSubnetOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteVirtualPrivateCloudSubnet
}

func (o *DeleteVirtualPrivateCloudSubnetOK) Error() string {
	return fmt.Sprintf("[POST /delete-virtual-private-cloud-subnet][%d] deleteVirtualPrivateCloudSubnetOK  %+v", 200, o.Payload)
}
func (o *DeleteVirtualPrivateCloudSubnetOK) GetPayload() []*models.WithTaskDeleteVirtualPrivateCloudSubnet {
	return o.Payload
}

func (o *DeleteVirtualPrivateCloudSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVirtualPrivateCloudSubnetBadRequest creates a DeleteVirtualPrivateCloudSubnetBadRequest with default headers values
func NewDeleteVirtualPrivateCloudSubnetBadRequest() *DeleteVirtualPrivateCloudSubnetBadRequest {
	return &DeleteVirtualPrivateCloudSubnetBadRequest{}
}

/* DeleteVirtualPrivateCloudSubnetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVirtualPrivateCloudSubnetBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVirtualPrivateCloudSubnetBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-virtual-private-cloud-subnet][%d] deleteVirtualPrivateCloudSubnetBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVirtualPrivateCloudSubnetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVirtualPrivateCloudSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVirtualPrivateCloudSubnetNotFound creates a DeleteVirtualPrivateCloudSubnetNotFound with default headers values
func NewDeleteVirtualPrivateCloudSubnetNotFound() *DeleteVirtualPrivateCloudSubnetNotFound {
	return &DeleteVirtualPrivateCloudSubnetNotFound{}
}

/* DeleteVirtualPrivateCloudSubnetNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVirtualPrivateCloudSubnetNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVirtualPrivateCloudSubnetNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-virtual-private-cloud-subnet][%d] deleteVirtualPrivateCloudSubnetNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVirtualPrivateCloudSubnetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVirtualPrivateCloudSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVirtualPrivateCloudSubnetInternalServerError creates a DeleteVirtualPrivateCloudSubnetInternalServerError with default headers values
func NewDeleteVirtualPrivateCloudSubnetInternalServerError() *DeleteVirtualPrivateCloudSubnetInternalServerError {
	return &DeleteVirtualPrivateCloudSubnetInternalServerError{}
}

/* DeleteVirtualPrivateCloudSubnetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVirtualPrivateCloudSubnetInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVirtualPrivateCloudSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-virtual-private-cloud-subnet][%d] deleteVirtualPrivateCloudSubnetInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVirtualPrivateCloudSubnetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVirtualPrivateCloudSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
