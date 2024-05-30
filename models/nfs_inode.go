// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NfsInode nfs inode
//
// swagger:model NfsInode
type NfsInode struct {

	// assigned size
	// Required: true
	AssignedSize *int64 `json:"assigned_size"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// file
	// Required: true
	File *bool `json:"file"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// local updated at
	// Required: true
	LocalUpdatedAt *string `json:"local_updated_at"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nfs export
	// Required: true
	NfsExport *NestedNfsExport `json:"nfs_export"`

	// parent id
	// Required: true
	ParentID *string `json:"parent_id"`

	// shared size
	// Required: true
	SharedSize *int64 `json:"shared_size"`

	// snapshot num
	// Required: true
	SnapshotNum *int32 `json:"snapshot_num"`

	// unique logical size
	UniqueLogicalSize *float64 `json:"unique_logical_size,omitempty"`

	// unique size
	// Required: true
	UniqueSize *int64 `json:"unique_size"`
}

// Validate validates this nfs inode
func (m *NfsInode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqueSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsInode) validateAssignedSize(formats strfmt.Registry) error {

	if err := validate.Required("assigned_size", "body", m.AssignedSize); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *NfsInode) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NfsInode) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateLocalUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_updated_at", "body", m.LocalUpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateNfsExport(formats strfmt.Registry) error {

	if err := validate.Required("nfs_export", "body", m.NfsExport); err != nil {
		return err
	}

	if m.NfsExport != nil {
		if err := m.NfsExport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

func (m *NfsInode) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateSharedSize(formats strfmt.Registry) error {

	if err := validate.Required("shared_size", "body", m.SharedSize); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateSnapshotNum(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_num", "body", m.SnapshotNum); err != nil {
		return err
	}

	return nil
}

func (m *NfsInode) validateUniqueSize(formats strfmt.Registry) error {

	if err := validate.Required("unique_size", "body", m.UniqueSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nfs inode based on the context it is used
func (m *NfsInode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsExport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsInode) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *NfsInode) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NfsInode) contextValidateNfsExport(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsExport != nil {
		if err := m.NfsExport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsInode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsInode) UnmarshalBinary(b []byte) error {
	var res NfsInode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
