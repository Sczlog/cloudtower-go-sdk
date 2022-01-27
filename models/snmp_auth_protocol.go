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

// SnmpAuthProtocol snmp auth protocol
//
// swagger:model SnmpAuthProtocol
type SnmpAuthProtocol string

func NewSnmpAuthProtocol(value SnmpAuthProtocol) *SnmpAuthProtocol {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SnmpAuthProtocol.
func (m SnmpAuthProtocol) Pointer() *SnmpAuthProtocol {
	return &m
}

const (

	// SnmpAuthProtocolMD5 captures enum value "MD5"
	SnmpAuthProtocolMD5 SnmpAuthProtocol = "MD5"

	// SnmpAuthProtocolSHA captures enum value "SHA"
	SnmpAuthProtocolSHA SnmpAuthProtocol = "SHA"
)

// for schema
var snmpAuthProtocolEnum []interface{}

func init() {
	var res []SnmpAuthProtocol
	if err := json.Unmarshal([]byte(`["MD5","SHA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpAuthProtocolEnum = append(snmpAuthProtocolEnum, v)
	}
}

func (m SnmpAuthProtocol) validateSnmpAuthProtocolEnum(path, location string, value SnmpAuthProtocol) error {
	if err := validate.EnumCase(path, location, value, snmpAuthProtocolEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snmp auth protocol
func (m SnmpAuthProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnmpAuthProtocolEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snmp auth protocol based on context it is used
func (m SnmpAuthProtocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
