// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupStoreRepository backup store repository
//
// swagger:model BackupStoreRepository
type BackupStoreRepository struct {

	// backup plans
	BackupPlans []*NestedBackupPlan `json:"backup_plans,omitempty"`

	// backup restore points
	BackupRestorePoints []*NestedBackupRestorePoint `json:"backup_restore_points,omitempty"`

	// backup service
	BackupService *NestedBackupService `json:"backup_service,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// iscsi chap name
	IscsiChapName *string `json:"iscsi_chap_name,omitempty"`

	// iscsi chap secret
	IscsiChapSecret *string `json:"iscsi_chap_secret,omitempty"`

	// iscsi ip
	IscsiIP *string `json:"iscsi_ip,omitempty"`

	// iscsi lun id
	IscsiLunID *string `json:"iscsi_lun_id,omitempty"`

	// iscsi port
	IscsiPort *int32 `json:"iscsi_port,omitempty"`

	// iscsi target iqn
	IscsiTargetIqn *string `json:"iscsi_target_iqn,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nfs path
	NfsPath *string `json:"nfs_path,omitempty"`

	// nfs server
	NfsServer *string `json:"nfs_server,omitempty"`

	// status
	// Required: true
	Status *BackupStoreStatus `json:"status"`

	// total capacity
	// Required: true
	TotalCapacity *float64 `json:"total_capacity"`

	// type
	// Required: true
	Type *BackupStoreType `json:"type"`

	// used data space
	// Required: true
	UsedDataSpace *float64 `json:"used_data_space"`

	// valid data space
	ValidDataSpace *float64 `json:"valid_data_space,omitempty"`
}

// Validate validates this backup store repository
func (m *BackupStoreRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupRestorePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedDataSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupStoreRepository) validateBackupPlans(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPlans) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupPlans); i++ {
		if swag.IsZero(m.BackupPlans[i]) { // not required
			continue
		}

		if m.BackupPlans[i] != nil {
			if err := m.BackupPlans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupStoreRepository) validateBackupRestorePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRestorePoints) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupRestorePoints); i++ {
		if swag.IsZero(m.BackupRestorePoints[i]) { // not required
			continue
		}

		if m.BackupRestorePoints[i] != nil {
			if err := m.BackupRestorePoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupStoreRepository) validateBackupService(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupService) { // not required
		return nil
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

func (m *BackupStoreRepository) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *BackupStoreRepository) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupStoreRepository) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BackupStoreRepository) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupStoreRepository) validateTotalCapacity(formats strfmt.Registry) error {

	if err := validate.Required("total_capacity", "body", m.TotalCapacity); err != nil {
		return err
	}

	return nil
}

func (m *BackupStoreRepository) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
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

func (m *BackupStoreRepository) validateUsedDataSpace(formats strfmt.Registry) error {

	if err := validate.Required("used_data_space", "body", m.UsedDataSpace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this backup store repository based on the context it is used
func (m *BackupStoreRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupRestorePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupStoreRepository) contextValidateBackupPlans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupPlans); i++ {

		if m.BackupPlans[i] != nil {
			if err := m.BackupPlans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupStoreRepository) contextValidateBackupRestorePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupRestorePoints); i++ {

		if m.BackupRestorePoints[i] != nil {
			if err := m.BackupRestorePoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupStoreRepository) contextValidateBackupService(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupStoreRepository) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupStoreRepository) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupStoreRepository) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BackupStoreRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupStoreRepository) UnmarshalBinary(b []byte) error {
	var res BackupStoreRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
