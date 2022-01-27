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

// NestedSnapshotGroupVMDiskInfo nested snapshot group Vm disk info
//
// swagger:model NestedSnapshotGroupVmDiskInfo
type NestedSnapshotGroupVMDiskInfo struct {

	// disk id
	// Required: true
	DiskID *string `json:"disk_id"`

	// disk snapshot status
	// Required: true
	DiskSnapshotStatus *ProtectSnapshotStatus `json:"disk_snapshot_status"`
}

// Validate validates this nested snapshot group Vm disk info
func (m *NestedSnapshotGroupVMDiskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSnapshotStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotGroupVMDiskInfo) validateDiskID(formats strfmt.Registry) error {

	if err := validate.Required("disk_id", "body", m.DiskID); err != nil {
		return err
	}

	return nil
}

func (m *NestedSnapshotGroupVMDiskInfo) validateDiskSnapshotStatus(formats strfmt.Registry) error {

	if err := validate.Required("disk_snapshot_status", "body", m.DiskSnapshotStatus); err != nil {
		return err
	}

	if err := validate.Required("disk_snapshot_status", "body", m.DiskSnapshotStatus); err != nil {
		return err
	}

	if m.DiskSnapshotStatus != nil {
		if err := m.DiskSnapshotStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_snapshot_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_snapshot_status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested snapshot group Vm disk info based on the context it is used
func (m *NestedSnapshotGroupVMDiskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiskSnapshotStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotGroupVMDiskInfo) contextValidateDiskSnapshotStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskSnapshotStatus != nil {
		if err := m.DiskSnapshotStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_snapshot_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_snapshot_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedSnapshotGroupVMDiskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedSnapshotGroupVMDiskInfo) UnmarshalBinary(b []byte) error {
	var res NestedSnapshotGroupVMDiskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
