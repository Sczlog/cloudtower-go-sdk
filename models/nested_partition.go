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

// NestedPartition nested partition
//
// swagger:model NestedPartition
type NestedPartition struct {

	// name
	Name *string `json:"name,omitempty"`

	// path
	Path *string `json:"path,omitempty"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// usage
	// Required: true
	Usage *PartitionUsage `json:"usage"`

	// used size
	// Required: true
	UsedSize *float64 `json:"used_size"`
}

// Validate validates this nested partition
func (m *NestedPartition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedPartition) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *NestedPartition) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

func (m *NestedPartition) validateUsedSize(formats strfmt.Registry) error {

	if err := validate.Required("used_size", "body", m.UsedSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested partition based on the context it is used
func (m *NestedPartition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedPartition) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {
		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedPartition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedPartition) UnmarshalBinary(b []byte) error {
	var res NestedPartition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
