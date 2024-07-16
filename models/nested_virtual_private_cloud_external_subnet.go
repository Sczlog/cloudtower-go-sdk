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

// NestedVirtualPrivateCloudExternalSubnet nested virtual private cloud external subnet
//
// swagger:model NestedVirtualPrivateCloudExternalSubnet
type NestedVirtualPrivateCloudExternalSubnet struct {

	// floating ip cidr
	FloatingIPCidr *string `json:"floating_ip_cidr,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nested virtual private cloud external subnet
func (m *NestedVirtualPrivateCloudExternalSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnet) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested virtual private cloud external subnet based on context it is used
func (m *NestedVirtualPrivateCloudExternalSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudExternalSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudExternalSubnet) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudExternalSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
