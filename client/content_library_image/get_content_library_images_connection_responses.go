// Code generated by go-swagger; DO NOT EDIT.

package content_library_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetContentLibraryImagesConnectionReader is a Reader for the GetContentLibraryImagesConnection structure.
type GetContentLibraryImagesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentLibraryImagesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentLibraryImagesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentLibraryImagesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentLibraryImagesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentLibraryImagesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentLibraryImagesConnectionOK creates a GetContentLibraryImagesConnectionOK with default headers values
func NewGetContentLibraryImagesConnectionOK() *GetContentLibraryImagesConnectionOK {
	return &GetContentLibraryImagesConnectionOK{}
}

/* GetContentLibraryImagesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetContentLibraryImagesConnectionOK struct {
	Payload *models.ContentLibraryImageConnection
}

func (o *GetContentLibraryImagesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images-connection][%d] getContentLibraryImagesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetContentLibraryImagesConnectionOK) GetPayload() *models.ContentLibraryImageConnection {
	return o.Payload
}

func (o *GetContentLibraryImagesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentLibraryImageConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentLibraryImagesConnectionBadRequest creates a GetContentLibraryImagesConnectionBadRequest with default headers values
func NewGetContentLibraryImagesConnectionBadRequest() *GetContentLibraryImagesConnectionBadRequest {
	return &GetContentLibraryImagesConnectionBadRequest{}
}

/* GetContentLibraryImagesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetContentLibraryImagesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentLibraryImagesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images-connection][%d] getContentLibraryImagesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetContentLibraryImagesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentLibraryImagesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentLibraryImagesConnectionNotFound creates a GetContentLibraryImagesConnectionNotFound with default headers values
func NewGetContentLibraryImagesConnectionNotFound() *GetContentLibraryImagesConnectionNotFound {
	return &GetContentLibraryImagesConnectionNotFound{}
}

/* GetContentLibraryImagesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetContentLibraryImagesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentLibraryImagesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images-connection][%d] getContentLibraryImagesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetContentLibraryImagesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentLibraryImagesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentLibraryImagesConnectionInternalServerError creates a GetContentLibraryImagesConnectionInternalServerError with default headers values
func NewGetContentLibraryImagesConnectionInternalServerError() *GetContentLibraryImagesConnectionInternalServerError {
	return &GetContentLibraryImagesConnectionInternalServerError{}
}

/* GetContentLibraryImagesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetContentLibraryImagesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentLibraryImagesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images-connection][%d] getContentLibraryImagesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetContentLibraryImagesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentLibraryImagesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}