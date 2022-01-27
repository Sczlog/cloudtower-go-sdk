// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// RestartVMReader is a Reader for the RestartVM structure.
type RestartVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestartVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartVMOK creates a RestartVMOK with default headers values
func NewRestartVMOK() *RestartVMOK {
	return &RestartVMOK{}
}

/* RestartVMOK describes a response with status code 200, with default header values.

Ok
*/
type RestartVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *RestartVMOK) Error() string {
	return fmt.Sprintf("[POST /restart-vm][%d] restartVmOK  %+v", 200, o.Payload)
}
func (o *RestartVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *RestartVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartVMBadRequest creates a RestartVMBadRequest with default headers values
func NewRestartVMBadRequest() *RestartVMBadRequest {
	return &RestartVMBadRequest{}
}

/* RestartVMBadRequest describes a response with status code 400, with default header values.

RestartVMBadRequest restart Vm bad request
*/
type RestartVMBadRequest struct {
	Payload string
}

func (o *RestartVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /restart-vm][%d] restartVmBadRequest  %+v", 400, o.Payload)
}
func (o *RestartVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RestartVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
