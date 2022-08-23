// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NestedStep nested step
//
// swagger:model NestedStep
type NestedStep struct {

	// current
	Current *float64 `json:"current,omitempty"`

	// finished
	Finished *bool `json:"finished,omitempty"`

	// key
	Key *string `json:"key,omitempty"`

	// per second
	PerSecond *float64 `json:"per_second,omitempty"`

	// total
	Total *float64 `json:"total,omitempty"`

	// unit
	Unit *StepUnit `json:"unit,omitempty"`
}

// Validate validates this nested step
func (m *NestedStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedStep) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested step based on the context it is used
func (m *NestedStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedStep) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedStep) UnmarshalBinary(b []byte) error {
	var res NestedStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}