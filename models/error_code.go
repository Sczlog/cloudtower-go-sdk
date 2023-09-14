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

// ErrorCode error code
//
// swagger:model ErrorCode
type ErrorCode string

func NewErrorCode(value ErrorCode) *ErrorCode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ErrorCode.
func (m ErrorCode) Pointer() *ErrorCode {
	return &m
}

const (

	// ErrorCodeResourceLocked captures enum value "ResourceLocked"
	ErrorCodeResourceLocked ErrorCode = "ResourceLocked"
)

// for schema
var errorCodeEnum []interface{}

func init() {
	var res []ErrorCode
	if err := json.Unmarshal([]byte(`["ResourceLocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorCodeEnum = append(errorCodeEnum, v)
	}
}

func (m ErrorCode) validateErrorCodeEnum(path, location string, value ErrorCode) error {
	if err := validate.EnumCase(path, location, value, errorCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this error code
func (m ErrorCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateErrorCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this error code based on context it is used
func (m ErrorCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
