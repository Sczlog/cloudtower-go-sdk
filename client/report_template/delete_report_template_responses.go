// Code generated by go-swagger; DO NOT EDIT.

package report_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteReportTemplateReader is a Reader for the DeleteReportTemplate structure.
type DeleteReportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteReportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteReportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteReportTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteReportTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteReportTemplateOK creates a DeleteReportTemplateOK with default headers values
func NewDeleteReportTemplateOK() *DeleteReportTemplateOK {
	return &DeleteReportTemplateOK{}
}

/* DeleteReportTemplateOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteReportTemplateOK struct {
	Payload []*models.WithTaskDeleteReportTemplate
}

func (o *DeleteReportTemplateOK) Error() string {
	return fmt.Sprintf("[POST /delete-report-template][%d] deleteReportTemplateOK  %+v", 200, o.Payload)
}
func (o *DeleteReportTemplateOK) GetPayload() []*models.WithTaskDeleteReportTemplate {
	return o.Payload
}

func (o *DeleteReportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportTemplateBadRequest creates a DeleteReportTemplateBadRequest with default headers values
func NewDeleteReportTemplateBadRequest() *DeleteReportTemplateBadRequest {
	return &DeleteReportTemplateBadRequest{}
}

/* DeleteReportTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteReportTemplateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteReportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-report-template][%d] deleteReportTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteReportTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteReportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportTemplateNotFound creates a DeleteReportTemplateNotFound with default headers values
func NewDeleteReportTemplateNotFound() *DeleteReportTemplateNotFound {
	return &DeleteReportTemplateNotFound{}
}

/* DeleteReportTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteReportTemplateNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteReportTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-report-template][%d] deleteReportTemplateNotFound  %+v", 404, o.Payload)
}
func (o *DeleteReportTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteReportTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportTemplateInternalServerError creates a DeleteReportTemplateInternalServerError with default headers values
func NewDeleteReportTemplateInternalServerError() *DeleteReportTemplateInternalServerError {
	return &DeleteReportTemplateInternalServerError{}
}

/* DeleteReportTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteReportTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteReportTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-report-template][%d] deleteReportTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteReportTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteReportTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
