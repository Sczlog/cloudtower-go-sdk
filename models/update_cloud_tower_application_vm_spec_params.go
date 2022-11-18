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

// UpdateCloudTowerApplicationVMSpecParams update cloud tower application Vm spec params
//
// swagger:model UpdateCloudTowerApplicationVmSpecParams
type UpdateCloudTowerApplicationVMSpecParams struct {

	// data
	// Required: true
	Data *UpdateCloudTowerApplicationVMSpecParamsData `json:"data"`

	// where
	// Required: true
	Where *CloudTowerApplicationWhereUniqueInput `json:"where"`
}

// Validate validates this update cloud tower application Vm spec params
func (m *UpdateCloudTowerApplicationVMSpecParams) Validate(formats strfmt.Registry) error {
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

func (m *UpdateCloudTowerApplicationVMSpecParams) validateData(formats strfmt.Registry) error {

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

func (m *UpdateCloudTowerApplicationVMSpecParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this update cloud tower application Vm spec params based on the context it is used
func (m *UpdateCloudTowerApplicationVMSpecParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *UpdateCloudTowerApplicationVMSpecParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UpdateCloudTowerApplicationVMSpecParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UpdateCloudTowerApplicationVMSpecParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudTowerApplicationVMSpecParams) UnmarshalBinary(b []byte) error {
	var res UpdateCloudTowerApplicationVMSpecParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateCloudTowerApplicationVMSpecParamsData update cloud tower application VM spec params data
//
// swagger:model UpdateCloudTowerApplicationVMSpecParamsData
type UpdateCloudTowerApplicationVMSpecParamsData struct {

	// vm spec
	// Required: true
	VMSpec *ApplicationVMSpecDefinition `json:"vmSpec"`
}

// Validate validates this update cloud tower application VM spec params data
func (m *UpdateCloudTowerApplicationVMSpecParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVMSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCloudTowerApplicationVMSpecParamsData) validateVMSpec(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"vmSpec", "body", m.VMSpec); err != nil {
		return err
	}

	if m.VMSpec != nil {
		if err := m.VMSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update cloud tower application VM spec params data based on the context it is used
func (m *UpdateCloudTowerApplicationVMSpecParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVMSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCloudTowerApplicationVMSpecParamsData) contextValidateVMSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.VMSpec != nil {
		if err := m.VMSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vmSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vmSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCloudTowerApplicationVMSpecParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudTowerApplicationVMSpecParamsData) UnmarshalBinary(b []byte) error {
	var res UpdateCloudTowerApplicationVMSpecParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
