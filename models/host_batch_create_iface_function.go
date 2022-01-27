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

// HostBatchCreateIfaceFunction host batch create iface function
//
// swagger:model HostBatchCreateIfaceFunction
type HostBatchCreateIfaceFunction string

func NewHostBatchCreateIfaceFunction(value HostBatchCreateIfaceFunction) *HostBatchCreateIfaceFunction {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HostBatchCreateIfaceFunction.
func (m HostBatchCreateIfaceFunction) Pointer() *HostBatchCreateIfaceFunction {
	return &m
}

const (

	// HostBatchCreateIfaceFunctionACCESS captures enum value "ACCESS"
	HostBatchCreateIfaceFunctionACCESS HostBatchCreateIfaceFunction = "ACCESS"

	// HostBatchCreateIfaceFunctionMANAGEMENT captures enum value "MANAGEMENT"
	HostBatchCreateIfaceFunctionMANAGEMENT HostBatchCreateIfaceFunction = "MANAGEMENT"

	// HostBatchCreateIfaceFunctionSTORAGE captures enum value "STORAGE"
	HostBatchCreateIfaceFunctionSTORAGE HostBatchCreateIfaceFunction = "STORAGE"

	// HostBatchCreateIfaceFunctionVMWAREACCESS captures enum value "VMWARE_ACCESS"
	HostBatchCreateIfaceFunctionVMWAREACCESS HostBatchCreateIfaceFunction = "VMWARE_ACCESS"
)

// for schema
var hostBatchCreateIfaceFunctionEnum []interface{}

func init() {
	var res []HostBatchCreateIfaceFunction
	if err := json.Unmarshal([]byte(`["ACCESS","MANAGEMENT","STORAGE","VMWARE_ACCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostBatchCreateIfaceFunctionEnum = append(hostBatchCreateIfaceFunctionEnum, v)
	}
}

func (m HostBatchCreateIfaceFunction) validateHostBatchCreateIfaceFunctionEnum(path, location string, value HostBatchCreateIfaceFunction) error {
	if err := validate.EnumCase(path, location, value, hostBatchCreateIfaceFunctionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this host batch create iface function
func (m HostBatchCreateIfaceFunction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHostBatchCreateIfaceFunctionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this host batch create iface function based on context it is used
func (m HostBatchCreateIfaceFunction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
