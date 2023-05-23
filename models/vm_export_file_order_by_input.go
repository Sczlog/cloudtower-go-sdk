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

// VMExportFileOrderByInput Vm export file order by input
//
// swagger:model VmExportFileOrderByInput
type VMExportFileOrderByInput string

func NewVMExportFileOrderByInput(value VMExportFileOrderByInput) *VMExportFileOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMExportFileOrderByInput.
func (m VMExportFileOrderByInput) Pointer() *VMExportFileOrderByInput {
	return &m
}

const (

	// VMExportFileOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VMExportFileOrderByInputCreatedAtASC VMExportFileOrderByInput = "createdAt_ASC"

	// VMExportFileOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VMExportFileOrderByInputCreatedAtDESC VMExportFileOrderByInput = "createdAt_DESC"

	// VMExportFileOrderByInputDamagedASC captures enum value "damaged_ASC"
	VMExportFileOrderByInputDamagedASC VMExportFileOrderByInput = "damaged_ASC"

	// VMExportFileOrderByInputDamagedDESC captures enum value "damaged_DESC"
	VMExportFileOrderByInputDamagedDESC VMExportFileOrderByInput = "damaged_DESC"

	// VMExportFileOrderByInputDataPortIDASC captures enum value "data_port_id_ASC"
	VMExportFileOrderByInputDataPortIDASC VMExportFileOrderByInput = "data_port_id_ASC"

	// VMExportFileOrderByInputDataPortIDDESC captures enum value "data_port_id_DESC"
	VMExportFileOrderByInputDataPortIDDESC VMExportFileOrderByInput = "data_port_id_DESC"

	// VMExportFileOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VMExportFileOrderByInputEntityAsyncStatusASC VMExportFileOrderByInput = "entityAsyncStatus_ASC"

	// VMExportFileOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VMExportFileOrderByInputEntityAsyncStatusDESC VMExportFileOrderByInput = "entityAsyncStatus_DESC"

	// VMExportFileOrderByInputFilesASC captures enum value "files_ASC"
	VMExportFileOrderByInputFilesASC VMExportFileOrderByInput = "files_ASC"

	// VMExportFileOrderByInputFilesDESC captures enum value "files_DESC"
	VMExportFileOrderByInputFilesDESC VMExportFileOrderByInput = "files_DESC"

	// VMExportFileOrderByInputIDASC captures enum value "id_ASC"
	VMExportFileOrderByInputIDASC VMExportFileOrderByInput = "id_ASC"

	// VMExportFileOrderByInputIDDESC captures enum value "id_DESC"
	VMExportFileOrderByInputIDDESC VMExportFileOrderByInput = "id_DESC"

	// VMExportFileOrderByInputStorageClusterIDASC captures enum value "storage_cluster_id_ASC"
	VMExportFileOrderByInputStorageClusterIDASC VMExportFileOrderByInput = "storage_cluster_id_ASC"

	// VMExportFileOrderByInputStorageClusterIDDESC captures enum value "storage_cluster_id_DESC"
	VMExportFileOrderByInputStorageClusterIDDESC VMExportFileOrderByInput = "storage_cluster_id_DESC"

	// VMExportFileOrderByInputTypeASC captures enum value "type_ASC"
	VMExportFileOrderByInputTypeASC VMExportFileOrderByInput = "type_ASC"

	// VMExportFileOrderByInputTypeDESC captures enum value "type_DESC"
	VMExportFileOrderByInputTypeDESC VMExportFileOrderByInput = "type_DESC"
)

// for schema
var vmExportFileOrderByInputEnum []interface{}

func init() {
	var res []VMExportFileOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","damaged_ASC","damaged_DESC","data_port_id_ASC","data_port_id_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","files_ASC","files_DESC","id_ASC","id_DESC","storage_cluster_id_ASC","storage_cluster_id_DESC","type_ASC","type_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmExportFileOrderByInputEnum = append(vmExportFileOrderByInputEnum, v)
	}
}

func (m VMExportFileOrderByInput) validateVMExportFileOrderByInputEnum(path, location string, value VMExportFileOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmExportFileOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm export file order by input
func (m VMExportFileOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMExportFileOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm export file order by input based on context it is used
func (m VMExportFileOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}