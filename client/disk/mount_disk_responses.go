// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// MountDiskReader is a Reader for the MountDisk structure.
type MountDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MountDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMountDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMountDiskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMountDiskOK creates a MountDiskOK with default headers values
func NewMountDiskOK() *MountDiskOK {
	return &MountDiskOK{}
}

/* MountDiskOK describes a response with status code 200, with default header values.

Ok
*/
type MountDiskOK struct {
	Payload []*models.WithTaskDisk
}

func (o *MountDiskOK) Error() string {
	return fmt.Sprintf("[POST /mount-disk][%d] mountDiskOK  %+v", 200, o.Payload)
}
func (o *MountDiskOK) GetPayload() []*models.WithTaskDisk {
	return o.Payload
}

func (o *MountDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMountDiskBadRequest creates a MountDiskBadRequest with default headers values
func NewMountDiskBadRequest() *MountDiskBadRequest {
	return &MountDiskBadRequest{}
}

/* MountDiskBadRequest describes a response with status code 400, with default header values.

MountDiskBadRequest mount disk bad request
*/
type MountDiskBadRequest struct {
	Payload string
}

func (o *MountDiskBadRequest) Error() string {
	return fmt.Sprintf("[POST /mount-disk][%d] mountDiskBadRequest  %+v", 400, o.Payload)
}
func (o *MountDiskBadRequest) GetPayload() string {
	return o.Payload
}

func (o *MountDiskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
