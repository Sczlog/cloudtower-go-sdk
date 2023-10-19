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

// GpuDeviceSriovSwitchParams gpu device sriov switch params
//
// swagger:model GpuDeviceSriovSwitchParams
type GpuDeviceSriovSwitchParams struct {

	// data
	// Required: true
	Data *GpuDeviceSriovSwitchParamsData `json:"data"`

	// where
	// Required: true
	Where *GpuDeviceWhereInput `json:"where"`
}

// Validate validates this gpu device sriov switch params
func (m *GpuDeviceSriovSwitchParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
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

func (m *GpuDeviceSriovSwitchParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *GpuDeviceSriovSwitchParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this gpu device sriov switch params based on the context it is used
func (m *GpuDeviceSriovSwitchParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuDeviceSriovSwitchParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *GpuDeviceSriovSwitchParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GpuDeviceSriovSwitchParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpuDeviceSriovSwitchParams) UnmarshalBinary(b []byte) error {
	var res GpuDeviceSriovSwitchParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GpuDeviceSriovSwitchParamsData gpu device sriov switch params data
//
// swagger:model GpuDeviceSriovSwitchParamsData
type GpuDeviceSriovSwitchParamsData struct {

	// enable
	// Required: true
	Enable *bool `json:"enable"`
}

// Validate validates this gpu device sriov switch params data
func (m *GpuDeviceSriovSwitchParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuDeviceSriovSwitchParamsData) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"enable", "body", m.Enable); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gpu device sriov switch params data based on context it is used
func (m *GpuDeviceSriovSwitchParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GpuDeviceSriovSwitchParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpuDeviceSriovSwitchParamsData) UnmarshalBinary(b []byte) error {
	var res GpuDeviceSriovSwitchParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
