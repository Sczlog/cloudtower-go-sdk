// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// UpdateDeployReader is a Reader for the UpdateDeploy structure.
type UpdateDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDeployBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDeployOK creates a UpdateDeployOK with default headers values
func NewUpdateDeployOK() *UpdateDeployOK {
	return &UpdateDeployOK{}
}

/* UpdateDeployOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateDeployOK struct {
	Payload *models.WithTaskLicense
}

func (o *UpdateDeployOK) Error() string {
	return fmt.Sprintf("[POST /update-license][%d] updateDeployOK  %+v", 200, o.Payload)
}
func (o *UpdateDeployOK) GetPayload() *models.WithTaskLicense {
	return o.Payload
}

func (o *UpdateDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WithTaskLicense)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeployBadRequest creates a UpdateDeployBadRequest with default headers values
func NewUpdateDeployBadRequest() *UpdateDeployBadRequest {
	return &UpdateDeployBadRequest{}
}

/* UpdateDeployBadRequest describes a response with status code 400, with default header values.

UpdateDeployBadRequest update deploy bad request
*/
type UpdateDeployBadRequest struct {
	Payload string
}

func (o *UpdateDeployBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-license][%d] updateDeployBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDeployBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateDeployBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
