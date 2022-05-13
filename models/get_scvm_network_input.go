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

// GetScvmNetworkInput get scvm network input
//
// swagger:model GetScvmNetworkInput
type GetScvmNetworkInput struct {

	// hosts
	// Required: true
	Hosts *HostWhereInput `json:"hosts"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// nics
	Nics *NicWhereInput `json:"nics,omitempty"`

	// range
	// Required: true
	Range *string `json:"range"`
}

// Validate validates this get scvm network input
func (m *GetScvmNetworkInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetScvmNetworkInput) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	if m.Hosts != nil {
		if err := m.Hosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts")
			}
			return err
		}
	}

	return nil
}

func (m *GetScvmNetworkInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetScvmNetworkInput) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	if m.Nics != nil {
		if err := m.Nics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics")
			}
			return err
		}
	}

	return nil
}

func (m *GetScvmNetworkInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get scvm network input based on the context it is used
func (m *GetScvmNetworkInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetScvmNetworkInput) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.Hosts != nil {
		if err := m.Hosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts")
			}
			return err
		}
	}

	return nil
}

func (m *GetScvmNetworkInput) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	if m.Nics != nil {
		if err := m.Nics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetScvmNetworkInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetScvmNetworkInput) UnmarshalBinary(b []byte) error {
	var res GetScvmNetworkInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}