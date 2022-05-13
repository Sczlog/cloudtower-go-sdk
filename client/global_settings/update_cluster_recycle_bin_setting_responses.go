// Code generated by go-swagger; DO NOT EDIT.

package global_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateClusterRecycleBinSettingReader is a Reader for the UpdateClusterRecycleBinSetting structure.
type UpdateClusterRecycleBinSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterRecycleBinSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterRecycleBinSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterRecycleBinSettingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterRecycleBinSettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterRecycleBinSettingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterRecycleBinSettingOK creates a UpdateClusterRecycleBinSettingOK with default headers values
func NewUpdateClusterRecycleBinSettingOK() *UpdateClusterRecycleBinSettingOK {
	return &UpdateClusterRecycleBinSettingOK{}
}

/* UpdateClusterRecycleBinSettingOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterRecycleBinSettingOK struct {
	Payload []*models.WithTaskClusterSettings
}

func (o *UpdateClusterRecycleBinSettingOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster-recycle-bin-setting][%d] updateClusterRecycleBinSettingOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterRecycleBinSettingOK) GetPayload() []*models.WithTaskClusterSettings {
	return o.Payload
}

func (o *UpdateClusterRecycleBinSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterRecycleBinSettingBadRequest creates a UpdateClusterRecycleBinSettingBadRequest with default headers values
func NewUpdateClusterRecycleBinSettingBadRequest() *UpdateClusterRecycleBinSettingBadRequest {
	return &UpdateClusterRecycleBinSettingBadRequest{}
}

/* UpdateClusterRecycleBinSettingBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateClusterRecycleBinSettingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterRecycleBinSettingBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster-recycle-bin-setting][%d] updateClusterRecycleBinSettingBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterRecycleBinSettingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterRecycleBinSettingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterRecycleBinSettingNotFound creates a UpdateClusterRecycleBinSettingNotFound with default headers values
func NewUpdateClusterRecycleBinSettingNotFound() *UpdateClusterRecycleBinSettingNotFound {
	return &UpdateClusterRecycleBinSettingNotFound{}
}

/* UpdateClusterRecycleBinSettingNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateClusterRecycleBinSettingNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterRecycleBinSettingNotFound) Error() string {
	return fmt.Sprintf("[POST /update-cluster-recycle-bin-setting][%d] updateClusterRecycleBinSettingNotFound  %+v", 404, o.Payload)
}
func (o *UpdateClusterRecycleBinSettingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterRecycleBinSettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterRecycleBinSettingInternalServerError creates a UpdateClusterRecycleBinSettingInternalServerError with default headers values
func NewUpdateClusterRecycleBinSettingInternalServerError() *UpdateClusterRecycleBinSettingInternalServerError {
	return &UpdateClusterRecycleBinSettingInternalServerError{}
}

/* UpdateClusterRecycleBinSettingInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateClusterRecycleBinSettingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterRecycleBinSettingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-cluster-recycle-bin-setting][%d] updateClusterRecycleBinSettingInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateClusterRecycleBinSettingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterRecycleBinSettingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
