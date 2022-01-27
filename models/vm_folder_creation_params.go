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

// VMFolderCreationParams Vm folder creation params
//
// swagger:model VmFolderCreationParams
type VMFolderCreationParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this Vm folder creation params
func (m *VMFolderCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMFolderCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMFolderCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vm folder creation params based on context it is used
func (m *VMFolderCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMFolderCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMFolderCreationParams) UnmarshalBinary(b []byte) error {
	var res VMFolderCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
