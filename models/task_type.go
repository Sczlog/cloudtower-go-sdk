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

// TaskType task type
//
// swagger:model TaskType
type TaskType string

func NewTaskType(value TaskType) *TaskType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TaskType.
func (m TaskType) Pointer() *TaskType {
	return &m
}

const (

	// TaskTypeAPPLICATION captures enum value "APPLICATION"
	TaskTypeAPPLICATION TaskType = "APPLICATION"

	// TaskTypeBACKUP captures enum value "BACKUP"
	TaskTypeBACKUP TaskType = "BACKUP"

	// TaskTypeRESOLVER captures enum value "RESOLVER"
	TaskTypeRESOLVER TaskType = "RESOLVER"
)

// for schema
var taskTypeEnum []interface{}

func init() {
	var res []TaskType
	if err := json.Unmarshal([]byte(`["APPLICATION","BACKUP","RESOLVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskTypeEnum = append(taskTypeEnum, v)
	}
}

func (m TaskType) validateTaskTypeEnum(path, location string, value TaskType) error {
	if err := validate.EnumCase(path, location, value, taskTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this task type
func (m TaskType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaskTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this task type based on context it is used
func (m TaskType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}