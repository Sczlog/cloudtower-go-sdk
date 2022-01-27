// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UploadTask upload task
//
// swagger:model UploadTask
type UploadTask struct {

	// args
	// Required: true
	Args interface{} `json:"args"`

	// chunk size
	// Required: true
	ChunkSize *float64 `json:"chunk_size"`

	// current chunk
	// Required: true
	CurrentChunk *int32 `json:"current_chunk"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// resource type
	// Required: true
	ResourceType *UploadResourceType `json:"resource_type"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// status
	// Required: true
	Status *UploadTaskStatus `json:"status"`

	// updated at
	// Required: true
	UpdatedAt *string `json:"updatedAt"`
}

// Validate validates this upload task
func (m *UploadTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChunkSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentChunk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTask) validateArgs(formats strfmt.Registry) error {

	if m.Args == nil {
		return errors.Required("args", "body", nil)
	}

	return nil
}

func (m *UploadTask) validateChunkSize(formats strfmt.Registry) error {

	if err := validate.Required("chunk_size", "body", m.ChunkSize); err != nil {
		return err
	}

	return nil
}

func (m *UploadTask) validateCurrentChunk(formats strfmt.Registry) error {

	if err := validate.Required("current_chunk", "body", m.CurrentChunk); err != nil {
		return err
	}

	return nil
}

func (m *UploadTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *UploadTask) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTask) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *UploadTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTask) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this upload task based on the context it is used
func (m *UploadTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTask) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceType != nil {
		if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTask) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UploadTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadTask) UnmarshalBinary(b []byte) error {
	var res UploadTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
