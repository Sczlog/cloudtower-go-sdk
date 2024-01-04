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

// ImportVMVolumeParams import Vm volume params
//
// swagger:model ImportVmVolumeParams
type ImportVMVolumeParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// storage policy
	// Required: true
	StoragePolicy *VMVolumeElfStoragePolicyType `json:"storage_policy"`

	// upload task id
	// Required: true
	UploadTaskID *string `json:"upload_task_id"`
}

// Validate validates this import Vm volume params
func (m *ImportVMVolumeParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportVMVolumeParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ImportVMVolumeParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ImportVMVolumeParams) validateStoragePolicy(formats strfmt.Registry) error {

	if err := validate.Required("storage_policy", "body", m.StoragePolicy); err != nil {
		return err
	}

	if err := validate.Required("storage_policy", "body", m.StoragePolicy); err != nil {
		return err
	}

	if m.StoragePolicy != nil {
		if err := m.StoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *ImportVMVolumeParams) validateUploadTaskID(formats strfmt.Registry) error {

	if err := validate.Required("upload_task_id", "body", m.UploadTaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this import Vm volume params based on the context it is used
func (m *ImportVMVolumeParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportVMVolumeParams) contextValidateStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePolicy != nil {
		if err := m.StoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImportVMVolumeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportVMVolumeParams) UnmarshalBinary(b []byte) error {
	var res ImportVMVolumeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}