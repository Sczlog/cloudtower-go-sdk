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

// NamespaceGroupCreationParams namespace group creation params
//
// swagger:model NamespaceGroupCreationParams
type NamespaceGroupCreationParams struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// nvmf subsystem id
	// Required: true
	NvmfSubsystemID *string `json:"nvmf_subsystem_id"`
}

// Validate validates this namespace group creation params
func (m *NamespaceGroupCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamespaceGroupCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroupCreationParams) validateNvmfSubsystemID(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem_id", "body", m.NvmfSubsystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this namespace group creation params based on context it is used
func (m *NamespaceGroupCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGroupCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGroupCreationParams) UnmarshalBinary(b []byte) error {
	var res NamespaceGroupCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
