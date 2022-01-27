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

// GetSystemAuditLogsConnectionRequestBody get system audit logs connection request body
// Example: {"after":"systemAuditLogsConnection-id-string","before":"systemAuditLogsConnection-id-string","first":0,"last":0,"orderBy":"action_ASC","skip":0,"where":{"AND":"SystemAuditLogWhereInput[]","NOT":"SystemAuditLogWhereInput[]","OR":"SystemAuditLogWhereInput[]","action":"string","action_contains":"string","action_ends_with":"string","action_gt":"string","action_gte":"string","action_in":["string"],"action_lt":"string","action_lte":"string","action_not":"string","action_not_contains":"string","action_not_ends_with":"string","action_not_in":["string"],"action_not_starts_with":"string","action_starts_with":"string","cluster":"ClusterWhereInput","finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","message":"string","message_contains":"string","message_ends_with":"string","message_gt":"string","message_gte":"string","message_in":["string"],"message_lt":"string","message_lte":"string","message_not":"string","message_not_contains":"string","message_not_ends_with":"string","message_not_in":["string"],"message_not_starts_with":"string","message_starts_with":"string","resource_id":"string","resource_id_contains":"string","resource_id_ends_with":"string","resource_id_gt":"string","resource_id_gte":"string","resource_id_in":["string"],"resource_id_lt":"string","resource_id_lte":"string","resource_id_not":"string","resource_id_not_contains":"string","resource_id_not_ends_with":"string","resource_id_not_in":["string"],"resource_id_not_starts_with":"string","resource_id_starts_with":"string","status":"FAILED","status_in":["FAILED"],"status_not":"FAILED","status_not_in":["FAILED"]}}
//
// swagger:model GetSystemAuditLogsConnectionRequestBody
type GetSystemAuditLogsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *SystemAuditLogOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *SystemAuditLogWhereInput `json:"where,omitempty"`
}

// Validate validates this get system audit logs connection request body
func (m *GetSystemAuditLogsConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetSystemAuditLogsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetSystemAuditLogsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get system audit logs connection request body based on the context it is used
func (m *GetSystemAuditLogsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetSystemAuditLogsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetSystemAuditLogsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetSystemAuditLogsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSystemAuditLogsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetSystemAuditLogsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
