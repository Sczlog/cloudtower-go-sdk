// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExecutePlan execute plan
//
// swagger:model ExecutePlan
type ExecutePlan struct {

	// typename
	// Enum: [ExecutePlan]
	Typename *string `json:"__typename,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Required: true
	ID *string `json:"id"`

	// period
	// Required: true
	Period *string `json:"period"`

	// retain
	// Required: true
	Retain *int32 `json:"retain"`

	// start at
	// Required: true
	StartAt *string `json:"start_at"`
}

// Validate validates this execute plan
func (m *ExecutePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var executePlanTypeTypenamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ExecutePlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executePlanTypeTypenamePropEnum = append(executePlanTypeTypenamePropEnum, v)
	}
}

const (

	// ExecutePlanTypenameExecutePlan captures enum value "ExecutePlan"
	ExecutePlanTypenameExecutePlan string = "ExecutePlan"
)

// prop value enum
func (m *ExecutePlan) validateTypenameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, executePlanTypeTypenamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExecutePlan) validateTypename(formats strfmt.Registry) error {
	if swag.IsZero(m.Typename) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypenameEnum("__typename", "body", *m.Typename); err != nil {
		return err
	}

	return nil
}

func (m *ExecutePlan) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ExecutePlan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ExecutePlan) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *ExecutePlan) validateRetain(formats strfmt.Registry) error {

	if err := validate.Required("retain", "body", m.Retain); err != nil {
		return err
	}

	return nil
}

func (m *ExecutePlan) validateStartAt(formats strfmt.Registry) error {

	if err := validate.Required("start_at", "body", m.StartAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this execute plan based on context it is used
func (m *ExecutePlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExecutePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutePlan) UnmarshalBinary(b []byte) error {
	var res ExecutePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
