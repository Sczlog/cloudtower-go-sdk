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

// AllocatableEcStorageCapacity allocatable ec storage capacity
//
// swagger:model AllocatableEcStorageCapacity
type AllocatableEcStorageCapacity struct {

	// capacity
	// Required: true
	Capacity *float64 `json:"capacity"`

	// k
	// Required: true
	K *int32 `json:"k"`

	// m
	// Required: true
	M *int32 `json:"m"`
}

// Validate validates this allocatable ec storage capacity
func (m *AllocatableEcStorageCapacity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatableEcStorageCapacity) validateCapacity(formats strfmt.Registry) error {

	if err := validate.Required("capacity", "body", m.Capacity); err != nil {
		return err
	}

	return nil
}

func (m *AllocatableEcStorageCapacity) validateK(formats strfmt.Registry) error {

	if err := validate.Required("k", "body", m.K); err != nil {
		return err
	}

	return nil
}

func (m *AllocatableEcStorageCapacity) validateM(formats strfmt.Registry) error {

	if err := validate.Required("m", "body", m.M); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this allocatable ec storage capacity based on context it is used
func (m *AllocatableEcStorageCapacity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllocatableEcStorageCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatableEcStorageCapacity) UnmarshalBinary(b []byte) error {
	var res AllocatableEcStorageCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}