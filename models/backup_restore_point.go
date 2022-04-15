// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupRestorePoint backup restore point
//
// swagger:model BackupRestorePoint
type BackupRestorePoint struct {

	// backup plan
	// Required: true
	BackupPlan *NestedBackupPlan `json:"backup_plan"`

	// backup service
	// Required: true
	BackupService *NestedBackupService `json:"backup_service"`

	// backup store repository
	// Required: true
	BackupStoreRepository *NestedBackupStoreRepository `json:"backup_store_repository"`

	// backup target execution
	// Required: true
	BackupTargetExecution *NestedBackupTargetExecution `json:"backup_target_execution"`

	// cluster local id
	ClusterLocalID *string `json:"cluster_local_id,omitempty"`

	// compression ratio
	CompressionRatio *float64 `json:"compression_ratio,omitempty"`

	// creation
	Creation *BackupRestorePointCreation `json:"creation,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// logical size
	LogicalSize *int64 `json:"logical_size,omitempty"`

	// parent restore point
	ParentRestorePoint *string `json:"parent_restore_point,omitempty"`

	// physical size
	PhysicalSize *int64 `json:"physical_size,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// slice
	// Required: true
	Slice *string `json:"slice"`

	// type
	Type *BackupRestorePointType `json:"type,omitempty"`

	// valid capacity
	ValidCapacity *int64 `json:"valid_capacity,omitempty"`

	// valid size
	ValidSize *int64 `json:"valid_size,omitempty"`

	// vm
	VM *NestedVM `json:"vm,omitempty"`

	// vm local id
	VMLocalID *string `json:"vm_local_id,omitempty"`

	// vm name
	VMName *string `json:"vm_name,omitempty"`
}

// Validate validates this backup restore point
func (m *BackupRestorePoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupStoreRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupTargetExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestorePoint) validateBackupPlan(formats strfmt.Registry) error {

	if err := validate.Required("backup_plan", "body", m.BackupPlan); err != nil {
		return err
	}

	if m.BackupPlan != nil {
		if err := m.BackupPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_plan")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateBackupService(formats strfmt.Registry) error {

	if err := validate.Required("backup_service", "body", m.BackupService); err != nil {
		return err
	}

	if m.BackupService != nil {
		if err := m.BackupService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_service")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateBackupStoreRepository(formats strfmt.Registry) error {

	if err := validate.Required("backup_store_repository", "body", m.BackupStoreRepository); err != nil {
		return err
	}

	if m.BackupStoreRepository != nil {
		if err := m.BackupStoreRepository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_store_repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_store_repository")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateBackupTargetExecution(formats strfmt.Registry) error {

	if err := validate.Required("backup_target_execution", "body", m.BackupTargetExecution); err != nil {
		return err
	}

	if m.BackupTargetExecution != nil {
		if err := m.BackupTargetExecution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_target_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_target_execution")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateCreation(formats strfmt.Registry) error {
	if swag.IsZero(m.Creation) { // not required
		return nil
	}

	if m.Creation != nil {
		if err := m.Creation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creation")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestorePoint) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestorePoint) validateSlice(formats strfmt.Registry) error {

	if err := validate.Required("slice", "body", m.Slice); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestorePoint) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) validateVM(formats strfmt.Registry) error {
	if swag.IsZero(m.VM) { // not required
		return nil
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup restore point based on the context it is used
func (m *BackupRestorePoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupStoreRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupTargetExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestorePoint) contextValidateBackupPlan(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupPlan != nil {
		if err := m.BackupPlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_plan")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateBackupService(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupService != nil {
		if err := m.BackupService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_service")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateBackupStoreRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupStoreRepository != nil {
		if err := m.BackupStoreRepository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_store_repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_store_repository")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateBackupTargetExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupTargetExecution != nil {
		if err := m.BackupTargetExecution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_target_execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_target_execution")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateCreation(ctx context.Context, formats strfmt.Registry) error {

	if m.Creation != nil {
		if err := m.Creation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creation")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestorePoint) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupRestorePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestorePoint) UnmarshalBinary(b []byte) error {
	var res BackupRestorePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
