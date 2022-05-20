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

// WitnessOrderByInput witness order by input
//
// swagger:model WitnessOrderByInput
type WitnessOrderByInput string

func NewWitnessOrderByInput(value WitnessOrderByInput) *WitnessOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated WitnessOrderByInput.
func (m WitnessOrderByInput) Pointer() *WitnessOrderByInput {
	return &m
}

const (

	// WitnessOrderByInputCPUHzPerCoreASC captures enum value "cpu_hz_per_core_ASC"
	WitnessOrderByInputCPUHzPerCoreASC WitnessOrderByInput = "cpu_hz_per_core_ASC"

	// WitnessOrderByInputCPUHzPerCoreDESC captures enum value "cpu_hz_per_core_DESC"
	WitnessOrderByInputCPUHzPerCoreDESC WitnessOrderByInput = "cpu_hz_per_core_DESC"

	// WitnessOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	WitnessOrderByInputCreatedAtASC WitnessOrderByInput = "createdAt_ASC"

	// WitnessOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	WitnessOrderByInputCreatedAtDESC WitnessOrderByInput = "createdAt_DESC"

	// WitnessOrderByInputDataIPASC captures enum value "data_ip_ASC"
	WitnessOrderByInputDataIPASC WitnessOrderByInput = "data_ip_ASC"

	// WitnessOrderByInputDataIPDESC captures enum value "data_ip_DESC"
	WitnessOrderByInputDataIPDESC WitnessOrderByInput = "data_ip_DESC"

	// WitnessOrderByInputIDASC captures enum value "id_ASC"
	WitnessOrderByInputIDASC WitnessOrderByInput = "id_ASC"

	// WitnessOrderByInputIDDESC captures enum value "id_DESC"
	WitnessOrderByInputIDDESC WitnessOrderByInput = "id_DESC"

	// WitnessOrderByInputLocalIDASC captures enum value "local_id_ASC"
	WitnessOrderByInputLocalIDASC WitnessOrderByInput = "local_id_ASC"

	// WitnessOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	WitnessOrderByInputLocalIDDESC WitnessOrderByInput = "local_id_DESC"

	// WitnessOrderByInputNameASC captures enum value "name_ASC"
	WitnessOrderByInputNameASC WitnessOrderByInput = "name_ASC"

	// WitnessOrderByInputNameDESC captures enum value "name_DESC"
	WitnessOrderByInputNameDESC WitnessOrderByInput = "name_DESC"

	// WitnessOrderByInputSystemDataCapacityASC captures enum value "system_data_capacity_ASC"
	WitnessOrderByInputSystemDataCapacityASC WitnessOrderByInput = "system_data_capacity_ASC"

	// WitnessOrderByInputSystemDataCapacityDESC captures enum value "system_data_capacity_DESC"
	WitnessOrderByInputSystemDataCapacityDESC WitnessOrderByInput = "system_data_capacity_DESC"

	// WitnessOrderByInputSystemUsedDataSpaceASC captures enum value "system_used_data_space_ASC"
	WitnessOrderByInputSystemUsedDataSpaceASC WitnessOrderByInput = "system_used_data_space_ASC"

	// WitnessOrderByInputSystemUsedDataSpaceDESC captures enum value "system_used_data_space_DESC"
	WitnessOrderByInputSystemUsedDataSpaceDESC WitnessOrderByInput = "system_used_data_space_DESC"

	// WitnessOrderByInputTotalCPUCoresASC captures enum value "total_cpu_cores_ASC"
	WitnessOrderByInputTotalCPUCoresASC WitnessOrderByInput = "total_cpu_cores_ASC"

	// WitnessOrderByInputTotalCPUCoresDESC captures enum value "total_cpu_cores_DESC"
	WitnessOrderByInputTotalCPUCoresDESC WitnessOrderByInput = "total_cpu_cores_DESC"

	// WitnessOrderByInputTotalCPUHzASC captures enum value "total_cpu_hz_ASC"
	WitnessOrderByInputTotalCPUHzASC WitnessOrderByInput = "total_cpu_hz_ASC"

	// WitnessOrderByInputTotalCPUHzDESC captures enum value "total_cpu_hz_DESC"
	WitnessOrderByInputTotalCPUHzDESC WitnessOrderByInput = "total_cpu_hz_DESC"

	// WitnessOrderByInputTotalMemoryBytesASC captures enum value "total_memory_bytes_ASC"
	WitnessOrderByInputTotalMemoryBytesASC WitnessOrderByInput = "total_memory_bytes_ASC"

	// WitnessOrderByInputTotalMemoryBytesDESC captures enum value "total_memory_bytes_DESC"
	WitnessOrderByInputTotalMemoryBytesDESC WitnessOrderByInput = "total_memory_bytes_DESC"

	// WitnessOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	WitnessOrderByInputUpdatedAtASC WitnessOrderByInput = "updatedAt_ASC"

	// WitnessOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	WitnessOrderByInputUpdatedAtDESC WitnessOrderByInput = "updatedAt_DESC"
)

// for schema
var witnessOrderByInputEnum []interface{}

func init() {
	var res []WitnessOrderByInput
	if err := json.Unmarshal([]byte(`["cpu_hz_per_core_ASC","cpu_hz_per_core_DESC","createdAt_ASC","createdAt_DESC","data_ip_ASC","data_ip_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","system_data_capacity_ASC","system_data_capacity_DESC","system_used_data_space_ASC","system_used_data_space_DESC","total_cpu_cores_ASC","total_cpu_cores_DESC","total_cpu_hz_ASC","total_cpu_hz_DESC","total_memory_bytes_ASC","total_memory_bytes_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		witnessOrderByInputEnum = append(witnessOrderByInputEnum, v)
	}
}

func (m WitnessOrderByInput) validateWitnessOrderByInputEnum(path, location string, value WitnessOrderByInput) error {
	if err := validate.EnumCase(path, location, value, witnessOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this witness order by input
func (m WitnessOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWitnessOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this witness order by input based on context it is used
func (m WitnessOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
