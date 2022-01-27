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

// SnmpVersion snmp version
//
// swagger:model SnmpVersion
type SnmpVersion string

func NewSnmpVersion(value SnmpVersion) *SnmpVersion {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SnmpVersion.
func (m SnmpVersion) Pointer() *SnmpVersion {
	return &m
}

const (

	// SnmpVersionV2C captures enum value "V2C"
	SnmpVersionV2C SnmpVersion = "V2C"

	// SnmpVersionV3 captures enum value "V3"
	SnmpVersionV3 SnmpVersion = "V3"
)

// for schema
var snmpVersionEnum []interface{}

func init() {
	var res []SnmpVersion
	if err := json.Unmarshal([]byte(`["V2C","V3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpVersionEnum = append(snmpVersionEnum, v)
	}
}

func (m SnmpVersion) validateSnmpVersionEnum(path, location string, value SnmpVersion) error {
	if err := validate.EnumCase(path, location, value, snmpVersionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snmp version
func (m SnmpVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnmpVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snmp version based on context it is used
func (m SnmpVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
