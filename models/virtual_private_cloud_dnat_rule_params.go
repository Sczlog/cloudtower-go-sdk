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

// VirtualPrivateCloudDnatRuleParams virtual private cloud dnat rule params
//
// swagger:model VirtualPrivateCloudDnatRuleParams
type VirtualPrivateCloudDnatRuleParams struct {

	// port
	// Required: true
	Port *int32 `json:"port"`

	// target ip
	// Required: true
	TargetIP *string `json:"target_ip"`

	// target port
	// Required: true
	TargetPort *int32 `json:"target_port"`
}

// Validate validates this virtual private cloud dnat rule params
func (m *VirtualPrivateCloudDnatRuleParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudDnatRuleParams) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudDnatRuleParams) validateTargetIP(formats strfmt.Registry) error {

	if err := validate.Required("target_ip", "body", m.TargetIP); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudDnatRuleParams) validateTargetPort(formats strfmt.Registry) error {

	if err := validate.Required("target_port", "body", m.TargetPort); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual private cloud dnat rule params based on context it is used
func (m *VirtualPrivateCloudDnatRuleParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudDnatRuleParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudDnatRuleParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudDnatRuleParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
