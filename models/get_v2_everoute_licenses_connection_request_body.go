// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetV2EverouteLicensesConnectionRequestBody get v2 everoute licenses connection request body
//
// swagger:model GetV2EverouteLicensesConnectionRequestBody
type GetV2EverouteLicensesConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *V2EverouteLicenseOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *V2EverouteLicenseWhereInput `json:"where,omitempty"`
}

// Validate validates this get v2 everoute licenses connection request body
func (m *GetV2EverouteLicensesConnectionRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
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

func (m *GetV2EverouteLicensesConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetV2EverouteLicensesConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
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

// ContextValidate validate this get v2 everoute licenses connection request body based on the context it is used
func (m *GetV2EverouteLicensesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
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

func (m *GetV2EverouteLicensesConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetV2EverouteLicensesConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetV2EverouteLicensesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetV2EverouteLicensesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetV2EverouteLicensesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
