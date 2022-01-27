// Code generated by go-swagger; DO NOT EDIT.

package consistency_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// UpdateConsistencyGroupReader is a Reader for the UpdateConsistencyGroup structure.
type UpdateConsistencyGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConsistencyGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConsistencyGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConsistencyGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConsistencyGroupOK creates a UpdateConsistencyGroupOK with default headers values
func NewUpdateConsistencyGroupOK() *UpdateConsistencyGroupOK {
	return &UpdateConsistencyGroupOK{}
}

/* UpdateConsistencyGroupOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateConsistencyGroupOK struct {
	Payload []*models.WithTaskConsistencyGroup
}

func (o *UpdateConsistencyGroupOK) Error() string {
	return fmt.Sprintf("[POST /update-consistency-group][%d] updateConsistencyGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateConsistencyGroupOK) GetPayload() []*models.WithTaskConsistencyGroup {
	return o.Payload
}

func (o *UpdateConsistencyGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsistencyGroupBadRequest creates a UpdateConsistencyGroupBadRequest with default headers values
func NewUpdateConsistencyGroupBadRequest() *UpdateConsistencyGroupBadRequest {
	return &UpdateConsistencyGroupBadRequest{}
}

/* UpdateConsistencyGroupBadRequest describes a response with status code 400, with default header values.

UpdateConsistencyGroupBadRequest update consistency group bad request
*/
type UpdateConsistencyGroupBadRequest struct {
	Payload string
}

func (o *UpdateConsistencyGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-consistency-group][%d] updateConsistencyGroupBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateConsistencyGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateConsistencyGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
