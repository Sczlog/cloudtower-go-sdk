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

// ClusterUpgradeHistoryOrderByInput cluster upgrade history order by input
//
// swagger:model ClusterUpgradeHistoryOrderByInput
type ClusterUpgradeHistoryOrderByInput string

func NewClusterUpgradeHistoryOrderByInput(value ClusterUpgradeHistoryOrderByInput) *ClusterUpgradeHistoryOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterUpgradeHistoryOrderByInput.
func (m ClusterUpgradeHistoryOrderByInput) Pointer() *ClusterUpgradeHistoryOrderByInput {
	return &m
}

const (

	// ClusterUpgradeHistoryOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ClusterUpgradeHistoryOrderByInputCreatedAtASC ClusterUpgradeHistoryOrderByInput = "createdAt_ASC"

	// ClusterUpgradeHistoryOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ClusterUpgradeHistoryOrderByInputCreatedAtDESC ClusterUpgradeHistoryOrderByInput = "createdAt_DESC"

	// ClusterUpgradeHistoryOrderByInputDateASC captures enum value "date_ASC"
	ClusterUpgradeHistoryOrderByInputDateASC ClusterUpgradeHistoryOrderByInput = "date_ASC"

	// ClusterUpgradeHistoryOrderByInputDateDESC captures enum value "date_DESC"
	ClusterUpgradeHistoryOrderByInputDateDESC ClusterUpgradeHistoryOrderByInput = "date_DESC"

	// ClusterUpgradeHistoryOrderByInputIDASC captures enum value "id_ASC"
	ClusterUpgradeHistoryOrderByInputIDASC ClusterUpgradeHistoryOrderByInput = "id_ASC"

	// ClusterUpgradeHistoryOrderByInputIDDESC captures enum value "id_DESC"
	ClusterUpgradeHistoryOrderByInputIDDESC ClusterUpgradeHistoryOrderByInput = "id_DESC"

	// ClusterUpgradeHistoryOrderByInputLocalIDASC captures enum value "local_id_ASC"
	ClusterUpgradeHistoryOrderByInputLocalIDASC ClusterUpgradeHistoryOrderByInput = "local_id_ASC"

	// ClusterUpgradeHistoryOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	ClusterUpgradeHistoryOrderByInputLocalIDDESC ClusterUpgradeHistoryOrderByInput = "local_id_DESC"

	// ClusterUpgradeHistoryOrderByInputProgressASC captures enum value "progress_ASC"
	ClusterUpgradeHistoryOrderByInputProgressASC ClusterUpgradeHistoryOrderByInput = "progress_ASC"

	// ClusterUpgradeHistoryOrderByInputProgressDESC captures enum value "progress_DESC"
	ClusterUpgradeHistoryOrderByInputProgressDESC ClusterUpgradeHistoryOrderByInput = "progress_DESC"

	// ClusterUpgradeHistoryOrderByInputResultASC captures enum value "result_ASC"
	ClusterUpgradeHistoryOrderByInputResultASC ClusterUpgradeHistoryOrderByInput = "result_ASC"

	// ClusterUpgradeHistoryOrderByInputResultDESC captures enum value "result_DESC"
	ClusterUpgradeHistoryOrderByInputResultDESC ClusterUpgradeHistoryOrderByInput = "result_DESC"

	// ClusterUpgradeHistoryOrderByInputTaskUUIDASC captures enum value "task_uuid_ASC"
	ClusterUpgradeHistoryOrderByInputTaskUUIDASC ClusterUpgradeHistoryOrderByInput = "task_uuid_ASC"

	// ClusterUpgradeHistoryOrderByInputTaskUUIDDESC captures enum value "task_uuid_DESC"
	ClusterUpgradeHistoryOrderByInputTaskUUIDDESC ClusterUpgradeHistoryOrderByInput = "task_uuid_DESC"

	// ClusterUpgradeHistoryOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ClusterUpgradeHistoryOrderByInputUpdatedAtASC ClusterUpgradeHistoryOrderByInput = "updatedAt_ASC"

	// ClusterUpgradeHistoryOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ClusterUpgradeHistoryOrderByInputUpdatedAtDESC ClusterUpgradeHistoryOrderByInput = "updatedAt_DESC"

	// ClusterUpgradeHistoryOrderByInputVersionASC captures enum value "version_ASC"
	ClusterUpgradeHistoryOrderByInputVersionASC ClusterUpgradeHistoryOrderByInput = "version_ASC"

	// ClusterUpgradeHistoryOrderByInputVersionDESC captures enum value "version_DESC"
	ClusterUpgradeHistoryOrderByInputVersionDESC ClusterUpgradeHistoryOrderByInput = "version_DESC"
)

// for schema
var clusterUpgradeHistoryOrderByInputEnum []interface{}

func init() {
	var res []ClusterUpgradeHistoryOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","date_ASC","date_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","progress_ASC","progress_DESC","result_ASC","result_DESC","task_uuid_ASC","task_uuid_DESC","updatedAt_ASC","updatedAt_DESC","version_ASC","version_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterUpgradeHistoryOrderByInputEnum = append(clusterUpgradeHistoryOrderByInputEnum, v)
	}
}

func (m ClusterUpgradeHistoryOrderByInput) validateClusterUpgradeHistoryOrderByInputEnum(path, location string, value ClusterUpgradeHistoryOrderByInput) error {
	if err := validate.EnumCase(path, location, value, clusterUpgradeHistoryOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster upgrade history order by input
func (m ClusterUpgradeHistoryOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterUpgradeHistoryOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster upgrade history order by input based on context it is used
func (m ClusterUpgradeHistoryOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
