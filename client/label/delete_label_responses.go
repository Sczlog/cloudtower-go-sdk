// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// DeleteLabelReader is a Reader for the DeleteLabel structure.
type DeleteLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLabelOK creates a DeleteLabelOK with default headers values
func NewDeleteLabelOK() *DeleteLabelOK {
	return &DeleteLabelOK{}
}

/* DeleteLabelOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteLabelOK struct {
	Payload []*models.WithTaskDeleteLabel
}

func (o *DeleteLabelOK) Error() string {
	return fmt.Sprintf("[POST /delete-label][%d] deleteLabelOK  %+v", 200, o.Payload)
}
func (o *DeleteLabelOK) GetPayload() []*models.WithTaskDeleteLabel {
	return o.Payload
}

func (o *DeleteLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLabelBadRequest creates a DeleteLabelBadRequest with default headers values
func NewDeleteLabelBadRequest() *DeleteLabelBadRequest {
	return &DeleteLabelBadRequest{}
}

/* DeleteLabelBadRequest describes a response with status code 400, with default header values.

DeleteLabelBadRequest delete label bad request
*/
type DeleteLabelBadRequest struct {
	Payload string
}

func (o *DeleteLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-label][%d] deleteLabelBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteLabelBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
