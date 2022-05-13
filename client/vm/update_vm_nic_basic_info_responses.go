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

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateVMNicBasicInfoReader is a Reader for the UpdateVMNicBasicInfo structure.
type UpdateVMNicBasicInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMNicBasicInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMNicBasicInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewUpdateVMNicBasicInfoNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewUpdateVMNicBasicInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMNicBasicInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVMNicBasicInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMNicBasicInfoOK creates a UpdateVMNicBasicInfoOK with default headers values
func NewUpdateVMNicBasicInfoOK() *UpdateVMNicBasicInfoOK {
	return &UpdateVMNicBasicInfoOK{}
}

/* UpdateVMNicBasicInfoOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVMNicBasicInfoOK struct {
	Payload []*UpdateVMNicBasicInfoOKBodyItems0
}

func (o *UpdateVMNicBasicInfoOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-basic-info][%d] updateVmNicBasicInfoOK  %+v", 200, o.Payload)
}
func (o *UpdateVMNicBasicInfoOK) GetPayload() []*UpdateVMNicBasicInfoOKBodyItems0 {
	return o.Payload
}

func (o *UpdateVMNicBasicInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicBasicInfoNotModified creates a UpdateVMNicBasicInfoNotModified with default headers values
func NewUpdateVMNicBasicInfoNotModified() *UpdateVMNicBasicInfoNotModified {
	return &UpdateVMNicBasicInfoNotModified{}
}

/* UpdateVMNicBasicInfoNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateVMNicBasicInfoNotModified struct {
}

func (o *UpdateVMNicBasicInfoNotModified) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-basic-info][%d] updateVmNicBasicInfoNotModified ", 304)
}

func (o *UpdateVMNicBasicInfoNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMNicBasicInfoBadRequest creates a UpdateVMNicBasicInfoBadRequest with default headers values
func NewUpdateVMNicBasicInfoBadRequest() *UpdateVMNicBasicInfoBadRequest {
	return &UpdateVMNicBasicInfoBadRequest{}
}

/* UpdateVMNicBasicInfoBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMNicBasicInfoBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicBasicInfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-basic-info][%d] updateVmNicBasicInfoBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMNicBasicInfoBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicBasicInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicBasicInfoNotFound creates a UpdateVMNicBasicInfoNotFound with default headers values
func NewUpdateVMNicBasicInfoNotFound() *UpdateVMNicBasicInfoNotFound {
	return &UpdateVMNicBasicInfoNotFound{}
}

/* UpdateVMNicBasicInfoNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMNicBasicInfoNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicBasicInfoNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-basic-info][%d] updateVmNicBasicInfoNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMNicBasicInfoNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicBasicInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicBasicInfoInternalServerError creates a UpdateVMNicBasicInfoInternalServerError with default headers values
func NewUpdateVMNicBasicInfoInternalServerError() *UpdateVMNicBasicInfoInternalServerError {
	return &UpdateVMNicBasicInfoInternalServerError{}
}

/* UpdateVMNicBasicInfoInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMNicBasicInfoInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateVMNicBasicInfoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic-basic-info][%d] updateVmNicBasicInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMNicBasicInfoInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicBasicInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateVMNicBasicInfoOKBodyItems0 update VM nic basic info o k body items0
swagger:model UpdateVMNicBasicInfoOKBodyItems0
*/
type UpdateVMNicBasicInfoOKBodyItems0 struct {

	// data
	// Required: true
	Data interface{} `json:"data"`

	// task id
	// Required: true
	TaskID interface{} `json:"task_id"`
}

// Validate validates this update VM nic basic info o k body items0
func (o *UpdateVMNicBasicInfoOKBodyItems0) Validate(formats strfmt.Registry) error {
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

func (o *UpdateVMNicBasicInfoOKBodyItems0) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("data", "body", nil)
	}

	return nil
}

func (o *UpdateVMNicBasicInfoOKBodyItems0) validateTaskID(formats strfmt.Registry) error {

	if o.TaskID == nil {
		return errors.Required("task_id", "body", nil)
	}

	return nil
}

// ContextValidate validates this update VM nic basic info o k body items0 based on context it is used
func (o *UpdateVMNicBasicInfoOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVMNicBasicInfoOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVMNicBasicInfoOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res UpdateVMNicBasicInfoOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
