// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// ForceStopLogCollectionReader is a Reader for the ForceStopLogCollection structure.
type ForceStopLogCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForceStopLogCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForceStopLogCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewForceStopLogCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForceStopLogCollectionOK creates a ForceStopLogCollectionOK with default headers values
func NewForceStopLogCollectionOK() *ForceStopLogCollectionOK {
	return &ForceStopLogCollectionOK{}
}

/* ForceStopLogCollectionOK describes a response with status code 200, with default header values.

Ok
*/
type ForceStopLogCollectionOK struct {
	Payload []*models.WithTaskLogCollection
}

func (o *ForceStopLogCollectionOK) Error() string {
	return fmt.Sprintf("[POST /force-stop-log-collection][%d] forceStopLogCollectionOK  %+v", 200, o.Payload)
}
func (o *ForceStopLogCollectionOK) GetPayload() []*models.WithTaskLogCollection {
	return o.Payload
}

func (o *ForceStopLogCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForceStopLogCollectionBadRequest creates a ForceStopLogCollectionBadRequest with default headers values
func NewForceStopLogCollectionBadRequest() *ForceStopLogCollectionBadRequest {
	return &ForceStopLogCollectionBadRequest{}
}

/* ForceStopLogCollectionBadRequest describes a response with status code 400, with default header values.

ForceStopLogCollectionBadRequest force stop log collection bad request
*/
type ForceStopLogCollectionBadRequest struct {
	Payload string
}

func (o *ForceStopLogCollectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /force-stop-log-collection][%d] forceStopLogCollectionBadRequest  %+v", 400, o.Payload)
}
func (o *ForceStopLogCollectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ForceStopLogCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
