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

// MgtIPMapper mgt Ip mapper
//
// swagger:model MgtIpMapper
type MgtIPMapper struct {

	// host id
	// Required: true
	HostID *string `json:"host_id"`

	// ip
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this mgt Ip mapper
func (m *MgtIPMapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MgtIPMapper) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *MgtIPMapper) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mgt Ip mapper based on context it is used
func (m *MgtIPMapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MgtIPMapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MgtIPMapper) UnmarshalBinary(b []byte) error {
	var res MgtIPMapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
