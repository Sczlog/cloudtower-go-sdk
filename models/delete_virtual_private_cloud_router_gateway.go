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

// DeleteVirtualPrivateCloudRouterGateway delete virtual private cloud router gateway
//
// swagger:model DeleteVirtualPrivateCloudRouterGateway
type DeleteVirtualPrivateCloudRouterGateway struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this delete virtual private cloud router gateway
func (m *DeleteVirtualPrivateCloudRouterGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteVirtualPrivateCloudRouterGateway) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete virtual private cloud router gateway based on context it is used
func (m *DeleteVirtualPrivateCloudRouterGateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudRouterGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudRouterGateway) UnmarshalBinary(b []byte) error {
	var res DeleteVirtualPrivateCloudRouterGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
