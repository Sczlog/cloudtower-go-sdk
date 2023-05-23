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

// UserCreationParams user creation params
//
// swagger:model UserCreationParams
type UserCreationParams struct {

	// auth config id
	AuthConfigID *string `json:"auth_config_id,omitempty"`

	// email address
	EmailAddress *string `json:"email_address,omitempty"`

	// internal
	Internal *bool `json:"internal,omitempty"`

	// ldap dn
	LdapDn *string `json:"ldap_dn,omitempty"`

	// mobile phone
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// password
	Password *string `json:"password,omitempty"`

	// role id
	// Required: true
	RoleID *string `json:"role_id"`

	// source
	Source *UserSource `json:"source,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user creation params
func (m *UserCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserCreationParams) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("role_id", "body", m.RoleID); err != nil {
		return err
	}

	return nil
}

func (m *UserCreationParams) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *UserCreationParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user creation params based on the context it is used
func (m *UserCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCreationParams) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCreationParams) UnmarshalBinary(b []byte) error {
	var res UserCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
