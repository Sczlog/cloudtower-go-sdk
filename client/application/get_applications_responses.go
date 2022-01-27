// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetApplicationsReader is a Reader for the GetApplications structure.
type GetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationsOK creates a GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {
	return &GetApplicationsOK{}
}

/* GetApplicationsOK describes a response with status code 200, with default header values.

Ok
*/
type GetApplicationsOK struct {
	Payload []*models.Application
}

func (o *GetApplicationsOK) Error() string {
	return fmt.Sprintf("[POST /get-applications][%d] getApplicationsOK  %+v", 200, o.Payload)
}
func (o *GetApplicationsOK) GetPayload() []*models.Application {
	return o.Payload
}

func (o *GetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsBadRequest creates a GetApplicationsBadRequest with default headers values
func NewGetApplicationsBadRequest() *GetApplicationsBadRequest {
	return &GetApplicationsBadRequest{}
}

/* GetApplicationsBadRequest describes a response with status code 400, with default header values.

GetApplicationsBadRequest get applications bad request
*/
type GetApplicationsBadRequest struct {
	Payload string
}

func (o *GetApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-applications][%d] getApplicationsBadRequest  %+v", 400, o.Payload)
}
func (o *GetApplicationsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
