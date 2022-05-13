// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateVMNicAdvanceInfoReader is a Reader for the UpdateVMNicAdvanceInfo structure.
type UpdateVMNicAdvanceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMNicAdvanceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMNicAdvanceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewUpdateVMNicAdvanceInfoNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewUpdateVMNicAdvanceInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMNicAdvanceInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVMNicAdvanceInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMNicAdvanceInfoOK creates a UpdateVMNicAdvanceInfoOK with default headers values
func NewUpdateVMNicAdvanceInfoOK() *UpdateVMNicAdvanceInfoOK {
	return &UpdateVMNicAdvanceInfoOK{}
}

/* UpdateVMNicAdvanceInfoOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVMNicAdvanceInfoOK struct {
	Payload []*UpdateVMNicAdvanceInfoOKBodyItems0
}

func (o *UpdateVMNicAdvanceInfoOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-advance-info][%d] updateVmNicAdvanceInfoOK  %+v", 200, o.Payload)
}
func (o *UpdateVMNicAdvanceInfoOK) GetPayload() []*UpdateVMNicAdvanceInfoOKBodyItems0 {
	return o.Payload
}

func (o *UpdateVMNicAdvanceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicAdvanceInfoNotModified creates a UpdateVMNicAdvanceInfoNotModified with default headers values
func NewUpdateVMNicAdvanceInfoNotModified() *UpdateVMNicAdvanceInfoNotModified {
	return &UpdateVMNicAdvanceInfoNotModified{}
}

/* UpdateVMNicAdvanceInfoNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateVMNicAdvanceInfoNotModified struct {
}

func (o *UpdateVMNicAdvanceInfoNotModified) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-advance-info][%d] updateVmNicAdvanceInfoNotModified ", 304)
}

func (o *UpdateVMNicAdvanceInfoNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMNicAdvanceInfoBadRequest creates a UpdateVMNicAdvanceInfoBadRequest with default headers values
func NewUpdateVMNicAdvanceInfoBadRequest() *UpdateVMNicAdvanceInfoBadRequest {
	return &UpdateVMNicAdvanceInfoBadRequest{}
}

/* UpdateVMNicAdvanceInfoBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMNicAdvanceInfoBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicAdvanceInfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-advance-info][%d] updateVmNicAdvanceInfoBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMNicAdvanceInfoBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicAdvanceInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicAdvanceInfoNotFound creates a UpdateVMNicAdvanceInfoNotFound with default headers values
func NewUpdateVMNicAdvanceInfoNotFound() *UpdateVMNicAdvanceInfoNotFound {
	return &UpdateVMNicAdvanceInfoNotFound{}
}

/* UpdateVMNicAdvanceInfoNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMNicAdvanceInfoNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicAdvanceInfoNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-advance-info][%d] updateVmNicAdvanceInfoNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMNicAdvanceInfoNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicAdvanceInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicAdvanceInfoInternalServerError creates a UpdateVMNicAdvanceInfoInternalServerError with default headers values
func NewUpdateVMNicAdvanceInfoInternalServerError() *UpdateVMNicAdvanceInfoInternalServerError {
	return &UpdateVMNicAdvanceInfoInternalServerError{}
}

/* UpdateVMNicAdvanceInfoInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMNicAdvanceInfoInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicAdvanceInfoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-advance-info][%d] updateVmNicAdvanceInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMNicAdvanceInfoInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicAdvanceInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateVMNicAdvanceInfoOKBodyItems0 update VM nic advance info o k body items0
swagger:model UpdateVMNicAdvanceInfoOKBodyItems0
*/
type UpdateVMNicAdvanceInfoOKBodyItems0 struct {

	// data
	// Required: true
	Data interface{} `json:"data"`

	// task id
	// Required: true
	TaskID interface{} `json:"task_id"`
}

// Validate validates this update VM nic advance info o k body items0
func (o *UpdateVMNicAdvanceInfoOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVMNicAdvanceInfoOKBodyItems0) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("data", "body", nil)
	}

	return nil
}

func (o *UpdateVMNicAdvanceInfoOKBodyItems0) validateTaskID(formats strfmt.Registry) error {

	if o.TaskID == nil {
		return errors.Required("task_id", "body", nil)
	}

	return nil
}

// ContextValidate validates this update VM nic advance info o k body items0 based on context it is used
func (o *UpdateVMNicAdvanceInfoOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVMNicAdvanceInfoOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVMNicAdvanceInfoOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UpdateVMNicAdvanceInfoOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
