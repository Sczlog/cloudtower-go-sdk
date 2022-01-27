// Code generated by go-swagger; DO NOT EDIT.

package nfs_inode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetNfsInodesReader is a Reader for the GetNfsInodes structure.
type GetNfsInodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsInodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsInodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNfsInodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNfsInodesOK creates a GetNfsInodesOK with default headers values
func NewGetNfsInodesOK() *GetNfsInodesOK {
	return &GetNfsInodesOK{}
}

/* GetNfsInodesOK describes a response with status code 200, with default header values.

Ok
*/
type GetNfsInodesOK struct {
	Payload []*models.NfsInode
}

func (o *GetNfsInodesOK) Error() string {
	return fmt.Sprintf("[POST /get-nfs-inodes][%d] getNfsInodesOK  %+v", 200, o.Payload)
}
func (o *GetNfsInodesOK) GetPayload() []*models.NfsInode {
	return o.Payload
}

func (o *GetNfsInodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsInodesBadRequest creates a GetNfsInodesBadRequest with default headers values
func NewGetNfsInodesBadRequest() *GetNfsInodesBadRequest {
	return &GetNfsInodesBadRequest{}
}

/* GetNfsInodesBadRequest describes a response with status code 400, with default header values.

GetNfsInodesBadRequest get nfs inodes bad request
*/
type GetNfsInodesBadRequest struct {
	Payload string
}

func (o *GetNfsInodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nfs-inodes][%d] getNfsInodesBadRequest  %+v", 400, o.Payload)
}
func (o *GetNfsInodesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNfsInodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
