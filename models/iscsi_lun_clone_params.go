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

// IscsiLunCloneParams iscsi lun clone params
//
// swagger:model IscsiLunCloneParams
type IscsiLunCloneParams struct {

	// iscsi target id
	// Required: true
	IscsiTargetID *string `json:"iscsi_target_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// snapshot id
	// Required: true
	SnapshotID *string `json:"snapshot_id"`
}

// Validate validates this iscsi lun clone params
func (m *IscsiLunCloneParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIscsiTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunCloneParams) validateIscsiTargetID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target_id", "body", m.IscsiTargetID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunCloneParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunCloneParams) validateSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_id", "body", m.SnapshotID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi lun clone params based on context it is used
func (m *IscsiLunCloneParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunCloneParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunCloneParams) UnmarshalBinary(b []byte) error {
	var res IscsiLunCloneParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
