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

// NestedVirtualPrivateCloudNatGateway nested virtual private cloud nat gateway
//
// swagger:model NestedVirtualPrivateCloudNatGateway
type NestedVirtualPrivateCloudNatGateway struct {

	// external ip
	ExternalIP *string `json:"external_ip,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nested virtual private cloud nat gateway
func (m *NestedVirtualPrivateCloudNatGateway) Validate(formats strfmt.Registry) error {
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

func (m *NestedVirtualPrivateCloudNatGateway) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNatGateway) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested virtual private cloud nat gateway based on context it is used
func (m *NestedVirtualPrivateCloudNatGateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudNatGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudNatGateway) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudNatGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
