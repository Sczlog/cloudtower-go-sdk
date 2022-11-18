// Code generated by go-swagger; DO NOT EDIT.

package cloud_tower_application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetCloudTowerApplicationsReader is a Reader for the GetCloudTowerApplications structure.
type GetCloudTowerApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudTowerApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudTowerApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudTowerApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudTowerApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudTowerApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudTowerApplicationsOK creates a GetCloudTowerApplicationsOK with default headers values
func NewGetCloudTowerApplicationsOK() *GetCloudTowerApplicationsOK {
	return &GetCloudTowerApplicationsOK{}
}

/* GetCloudTowerApplicationsOK describes a response with status code 200, with default header values.

Ok
*/
type GetCloudTowerApplicationsOK struct {
	Payload []*models.CloudTowerApplication
}

func (o *GetCloudTowerApplicationsOK) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-applications][%d] getCloudTowerApplicationsOK  %+v", 200, o.Payload)
}
func (o *GetCloudTowerApplicationsOK) GetPayload() []*models.CloudTowerApplication {
	return o.Payload
}

func (o *GetCloudTowerApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsBadRequest creates a GetCloudTowerApplicationsBadRequest with default headers values
func NewGetCloudTowerApplicationsBadRequest() *GetCloudTowerApplicationsBadRequest {
	return &GetCloudTowerApplicationsBadRequest{}
}

/* GetCloudTowerApplicationsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCloudTowerApplicationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-applications][%d] getCloudTowerApplicationsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCloudTowerApplicationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsNotFound creates a GetCloudTowerApplicationsNotFound with default headers values
func NewGetCloudTowerApplicationsNotFound() *GetCloudTowerApplicationsNotFound {
	return &GetCloudTowerApplicationsNotFound{}
}

/* GetCloudTowerApplicationsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCloudTowerApplicationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-applications][%d] getCloudTowerApplicationsNotFound  %+v", 404, o.Payload)
}
func (o *GetCloudTowerApplicationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsInternalServerError creates a GetCloudTowerApplicationsInternalServerError with default headers values
func NewGetCloudTowerApplicationsInternalServerError() *GetCloudTowerApplicationsInternalServerError {
	return &GetCloudTowerApplicationsInternalServerError{}
}

/* GetCloudTowerApplicationsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetCloudTowerApplicationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-applications][%d] getCloudTowerApplicationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCloudTowerApplicationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
