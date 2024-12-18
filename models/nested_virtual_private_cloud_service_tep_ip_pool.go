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

// NestedVirtualPrivateCloudServiceTepIPPool nested virtual private cloud service tep Ip pool
//
// swagger:model NestedVirtualPrivateCloudServiceTepIpPool
type NestedVirtualPrivateCloudServiceTepIPPool struct {

	// display name
	// Required: true
	DisplayName *string `json:"display_name"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// ip end
	// Required: true
	IPEnd *string `json:"ip_end"`

	// ip start
	// Required: true
	IPStart *string `json:"ip_start"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`
}

// Validate validates this nested virtual private cloud service tep Ip pool
func (m *NestedVirtualPrivateCloudServiceTepIPPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudServiceTepIPPool) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudServiceTepIPPool) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudServiceTepIPPool) validateIPEnd(formats strfmt.Registry) error {

	if err := validate.Required("ip_end", "body", m.IPEnd); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudServiceTepIPPool) validateIPStart(formats strfmt.Registry) error {

	if err := validate.Required("ip_start", "body", m.IPStart); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudServiceTepIPPool) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested virtual private cloud service tep Ip pool based on context it is used
func (m *NestedVirtualPrivateCloudServiceTepIPPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudServiceTepIPPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudServiceTepIPPool) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudServiceTepIPPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
