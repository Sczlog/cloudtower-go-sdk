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

// ClusterEnableIscsiUpdationParams cluster enable iscsi updation params
//
// swagger:model ClusterEnableIscsiUpdationParams
type ClusterEnableIscsiUpdationParams struct {

	// enable iscsi
	// Required: true
	EnableIscsi *bool `json:"enable_iscsi"`

	// where
	// Required: true
	Where *ClusterWhereInput `json:"where"`
}

// Validate validates this cluster enable iscsi updation params
func (m *ClusterEnableIscsiUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnableIscsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterEnableIscsiUpdationParams) validateEnableIscsi(formats strfmt.Registry) error {

	if err := validate.Required("enable_iscsi", "body", m.EnableIscsi); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEnableIscsiUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster enable iscsi updation params based on the context it is used
func (m *ClusterEnableIscsiUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterEnableIscsiUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterEnableIscsiUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterEnableIscsiUpdationParams) UnmarshalBinary(b []byte) error {
	var res ClusterEnableIscsiUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
