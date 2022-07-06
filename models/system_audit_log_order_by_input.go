// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SystemAuditLogOrderByInput system audit log order by input
//
// swagger:model SystemAuditLogOrderByInput
type SystemAuditLogOrderByInput string

func NewSystemAuditLogOrderByInput(value SystemAuditLogOrderByInput) *SystemAuditLogOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SystemAuditLogOrderByInput.
func (m SystemAuditLogOrderByInput) Pointer() *SystemAuditLogOrderByInput {
	return &m
}

const (

	// SystemAuditLogOrderByInputActionASC captures enum value "action_ASC"
	SystemAuditLogOrderByInputActionASC SystemAuditLogOrderByInput = "action_ASC"

	// SystemAuditLogOrderByInputActionDESC captures enum value "action_DESC"
	SystemAuditLogOrderByInputActionDESC SystemAuditLogOrderByInput = "action_DESC"

	// SystemAuditLogOrderByInputFinishedAtASC captures enum value "finished_at_ASC"
	SystemAuditLogOrderByInputFinishedAtASC SystemAuditLogOrderByInput = "finished_at_ASC"

	// SystemAuditLogOrderByInputFinishedAtDESC captures enum value "finished_at_DESC"
	SystemAuditLogOrderByInputFinishedAtDESC SystemAuditLogOrderByInput = "finished_at_DESC"

	// SystemAuditLogOrderByInputIDASC captures enum value "id_ASC"
	SystemAuditLogOrderByInputIDASC SystemAuditLogOrderByInput = "id_ASC"

	// SystemAuditLogOrderByInputIDDESC captures enum value "id_DESC"
	SystemAuditLogOrderByInputIDDESC SystemAuditLogOrderByInput = "id_DESC"

	// SystemAuditLogOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	SystemAuditLogOrderByInputLocalCreatedAtASC SystemAuditLogOrderByInput = "local_created_at_ASC"

	// SystemAuditLogOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	SystemAuditLogOrderByInputLocalCreatedAtDESC SystemAuditLogOrderByInput = "local_created_at_DESC"

	// SystemAuditLogOrderByInputLocalIDASC captures enum value "local_id_ASC"
	SystemAuditLogOrderByInputLocalIDASC SystemAuditLogOrderByInput = "local_id_ASC"

	// SystemAuditLogOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	SystemAuditLogOrderByInputLocalIDDESC SystemAuditLogOrderByInput = "local_id_DESC"

	// SystemAuditLogOrderByInputMessageASC captures enum value "message_ASC"
	SystemAuditLogOrderByInputMessageASC SystemAuditLogOrderByInput = "message_ASC"

	// SystemAuditLogOrderByInputMessageDESC captures enum value "message_DESC"
	SystemAuditLogOrderByInputMessageDESC SystemAuditLogOrderByInput = "message_DESC"

	// SystemAuditLogOrderByInputResourceIDASC captures enum value "resource_id_ASC"
	SystemAuditLogOrderByInputResourceIDASC SystemAuditLogOrderByInput = "resource_id_ASC"

	// SystemAuditLogOrderByInputResourceIDDESC captures enum value "resource_id_DESC"
	SystemAuditLogOrderByInputResourceIDDESC SystemAuditLogOrderByInput = "resource_id_DESC"

	// SystemAuditLogOrderByInputStatusASC captures enum value "status_ASC"
	SystemAuditLogOrderByInputStatusASC SystemAuditLogOrderByInput = "status_ASC"

	// SystemAuditLogOrderByInputStatusDESC captures enum value "status_DESC"
	SystemAuditLogOrderByInputStatusDESC SystemAuditLogOrderByInput = "status_DESC"
)

// for schema
var systemAuditLogOrderByInputEnum []interface{}

func init() {
	var res []SystemAuditLogOrderByInput
	if err := json.Unmarshal([]byte(`["action_ASC","action_DESC","finished_at_ASC","finished_at_DESC","id_ASC","id_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","message_ASC","message_DESC","resource_id_ASC","resource_id_DESC","status_ASC","status_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemAuditLogOrderByInputEnum = append(systemAuditLogOrderByInputEnum, v)
	}
}

func (m SystemAuditLogOrderByInput) validateSystemAuditLogOrderByInputEnum(path, location string, value SystemAuditLogOrderByInput) error {
	if err := validate.EnumCase(path, location, value, systemAuditLogOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this system audit log order by input
func (m SystemAuditLogOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSystemAuditLogOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this system audit log order by input based on context it is used
func (m SystemAuditLogOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
