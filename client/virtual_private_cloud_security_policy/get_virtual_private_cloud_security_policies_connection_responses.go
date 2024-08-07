// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_security_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVirtualPrivateCloudSecurityPoliciesConnectionReader is a Reader for the GetVirtualPrivateCloudSecurityPoliciesConnection structure.
type GetVirtualPrivateCloudSecurityPoliciesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualPrivateCloudSecurityPoliciesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVirtualPrivateCloudSecurityPoliciesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVirtualPrivateCloudSecurityPoliciesConnectionOK creates a GetVirtualPrivateCloudSecurityPoliciesConnectionOK with default headers values
func NewGetVirtualPrivateCloudSecurityPoliciesConnectionOK() *GetVirtualPrivateCloudSecurityPoliciesConnectionOK {
	return &GetVirtualPrivateCloudSecurityPoliciesConnectionOK{}
}

/* GetVirtualPrivateCloudSecurityPoliciesConnectionOK describes a response with status code 200, with default header values.

GetVirtualPrivateCloudSecurityPoliciesConnectionOK get virtual private cloud security policies connection o k
*/
type GetVirtualPrivateCloudSecurityPoliciesConnectionOK struct {
	XTowerRequestID string

	Payload *models.VirtualPrivateCloudSecurityPolicyConnection
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-security-policies-connection][%d] getVirtualPrivateCloudSecurityPoliciesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionOK) GetPayload() *models.VirtualPrivateCloudSecurityPolicyConnection {
	return o.Payload
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.VirtualPrivateCloudSecurityPolicyConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest creates a GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest with default headers values
func NewGetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest() *GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest {
	return &GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest{}
}

/* GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-security-policies-connection][%d] getVirtualPrivateCloudSecurityPoliciesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudSecurityPoliciesConnectionNotFound creates a GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound with default headers values
func NewGetVirtualPrivateCloudSecurityPoliciesConnectionNotFound() *GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound {
	return &GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound{}
}

/* GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-security-policies-connection][%d] getVirtualPrivateCloudSecurityPoliciesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError creates a GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError with default headers values
func NewGetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError() *GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError {
	return &GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError{}
}

/* GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-security-policies-connection][%d] getVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudSecurityPoliciesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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