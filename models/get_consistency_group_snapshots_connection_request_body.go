// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetConsistencyGroupSnapshotsConnectionRequestBody get consistency group snapshots connection request body
// Example: {"after":"consistencyGroupSnapshotsConnection-id-string","before":"consistencyGroupSnapshotsConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"ConsistencyGroupSnapshotWhereInput[]","Iscsi_lun_snapshots_every":"IscsiLunSnapshotWhereInput","Iscsi_lun_snapshots_none":"IscsiLunSnapshotWhereInput","Iscsi_lun_snapshots_some":"IscsiLunSnapshotWhereInput","NOT":"ConsistencyGroupSnapshotWhereInput[]","OR":"ConsistencyGroupSnapshotWhereInput[]","consistency_group":"ConsistencyGroupWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","nvmf_namespace_snapshots_every":"NvmfNamespaceSnapshotWhereInput","nvmf_namespace_snapshots_none":"NvmfNamespaceSnapshotWhereInput","nvmf_namespace_snapshots_some":"NvmfNamespaceSnapshotWhereInput","unique_size":0,"unique_size_gt":0,"unique_size_gte":0,"unique_size_in":[0],"unique_size_lt":0,"unique_size_lte":0,"unique_size_not":0,"unique_size_not_in":[0]}}
//
// swagger:model GetConsistencyGroupSnapshotsConnectionRequestBody
type GetConsistencyGroupSnapshotsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *ConsistencyGroupSnapshotOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *ConsistencyGroupSnapshotWhereInput `json:"where,omitempty"`
}

// Validate validates this get consistency group snapshots connection request body
func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get consistency group snapshots connection request body based on the context it is used
func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConsistencyGroupSnapshotsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetConsistencyGroupSnapshotsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
