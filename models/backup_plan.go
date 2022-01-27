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

// BackupPlan backup plan
//
// swagger:model BackupPlan
type BackupPlan struct {

	// backup delay option
	BackupDelayOption *BackupPlanDelayOption `json:"backup_delay_option,omitempty"`

	// backup plan executions
	BackupPlanExecutions []*NestedBackupPlanExecution `json:"backup_plan_executions,omitempty"`

	// backup restore point count
	BackupRestorePointCount *int32 `json:"backup_restore_point_count,omitempty"`

	// backup restore points
	BackupRestorePoints []*NestedBackupRestorePoint `json:"backup_restore_points,omitempty"`

	// backup service
	// Required: true
	BackupService *NestedBackupService `json:"backup_service"`

	// backup store repository
	// Required: true
	BackupStoreRepository *NestedBackupStoreRepository `json:"backup_store_repository"`

	// backup total size
	BackupTotalSize *float64 `json:"backup_total_size,omitempty"`

	// compression
	Compression *bool `json:"compression,omitempty"`

	// delete strategy
	DeleteStrategy *BackupPlanDeleteStrategy `json:"delete_strategy,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// enable window
	// Required: true
	EnableWindow *bool `json:"enable_window"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// full interval
	// Required: true
	FullInterval *int32 `json:"full_interval"`

	// full period
	// Required: true
	FullPeriod *BackupPlanPeriod `json:"full_period"`

	// full time point
	// Required: true
	FullTimePoint *NestedBackupPlanTimePoint `json:"full_time_point"`

	// id
	// Required: true
	ID *string `json:"id"`

	// incremental period
	// Required: true
	IncrementalPeriod *BackupPlanPeriod `json:"incremental_period"`

	// incremental time points
	// Required: true
	IncrementalTimePoints []*NestedBackupPlanTimePoint `json:"incremental_time_points"`

	// incremental weekdays
	IncrementalWeekdays []WeekdayTypeEnum `json:"incremental_weekdays,omitempty"`

	// keep policy
	KeepPolicy *BackupPlanKeepPolicy `json:"keep_policy,omitempty"`

	// keep policy value
	KeepPolicyValue *int32 `json:"keep_policy_value,omitempty"`

	// last execute status
	// Required: true
	LastExecuteStatus *BackupPlanExecutionStatus `json:"last_execute_status"`

	// last execute status message
	LastExecuteStatusMessage *string `json:"last_execute_status_message,omitempty"`

	// last execute success job count
	LastExecuteSuccessJobCount *int32 `json:"last_execute_success_job_count,omitempty"`

	// last execute total job count
	LastExecuteTotalJobCount *int32 `json:"last_execute_total_job_count,omitempty"`

	// last executed at
	LastExecutedAt *string `json:"last_executed_at,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace
	Namespace *string `json:"namespace,omitempty"`

	// next execute time
	NextExecuteTime *string `json:"next_execute_time,omitempty"`

	// status
	// Required: true
	Status *BackupPlanStatus `json:"status"`

	// vms
	Vms []*NestedVM `json:"vms,omitempty"`

	// window end
	WindowEnd *string `json:"window_end,omitempty"`

	// window start
	WindowStart *string `json:"window_start,omitempty"`
}

// Validate validates this backup plan
func (m *BackupPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupDelayOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupPlanExecutions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupRestorePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupStoreRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullTimePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalTimePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalWeekdays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeepPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastExecuteStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlan) validateBackupDelayOption(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupDelayOption) { // not required
		return nil
	}

	if m.BackupDelayOption != nil {
		if err := m.BackupDelayOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_delay_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_delay_option")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateBackupPlanExecutions(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPlanExecutions) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupPlanExecutions); i++ {
		if swag.IsZero(m.BackupPlanExecutions[i]) { // not required
			continue
		}

		if m.BackupPlanExecutions[i] != nil {
			if err := m.BackupPlanExecutions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_plan_executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_plan_executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlan) validateBackupRestorePoints(formats strfmt.Registry) error {
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

func (m *BackupPlan) validateBackupService(formats strfmt.Registry) error {

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

func (m *BackupPlan) validateBackupStoreRepository(formats strfmt.Registry) error {

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

func (m *BackupPlan) validateDeleteStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteStrategy) { // not required
		return nil
	}

	if m.DeleteStrategy != nil {
		if err := m.DeleteStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delete_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delete_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateEnableWindow(formats strfmt.Registry) error {

	if err := validate.Required("enable_window", "body", m.EnableWindow); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlan) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *BackupPlan) validateFullInterval(formats strfmt.Registry) error {

	if err := validate.Required("full_interval", "body", m.FullInterval); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlan) validateFullPeriod(formats strfmt.Registry) error {

	if err := validate.Required("full_period", "body", m.FullPeriod); err != nil {
		return err
	}

	if err := validate.Required("full_period", "body", m.FullPeriod); err != nil {
		return err
	}

	if m.FullPeriod != nil {
		if err := m.FullPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateFullTimePoint(formats strfmt.Registry) error {

	if err := validate.Required("full_time_point", "body", m.FullTimePoint); err != nil {
		return err
	}

	if m.FullTimePoint != nil {
		if err := m.FullTimePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full_time_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full_time_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlan) validateIncrementalPeriod(formats strfmt.Registry) error {

	if err := validate.Required("incremental_period", "body", m.IncrementalPeriod); err != nil {
		return err
	}

	if err := validate.Required("incremental_period", "body", m.IncrementalPeriod); err != nil {
		return err
	}

	if m.IncrementalPeriod != nil {
		if err := m.IncrementalPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateIncrementalTimePoints(formats strfmt.Registry) error {

	if err := validate.Required("incremental_time_points", "body", m.IncrementalTimePoints); err != nil {
		return err
	}

	for i := 0; i < len(m.IncrementalTimePoints); i++ {
		if swag.IsZero(m.IncrementalTimePoints[i]) { // not required
			continue
		}

		if m.IncrementalTimePoints[i] != nil {
			if err := m.IncrementalTimePoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incremental_time_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incremental_time_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlan) validateIncrementalWeekdays(formats strfmt.Registry) error {
	if swag.IsZero(m.IncrementalWeekdays) { // not required
		return nil
	}

	for i := 0; i < len(m.IncrementalWeekdays); i++ {

		if err := m.IncrementalWeekdays[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental_weekdays" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental_weekdays" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupPlan) validateKeepPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.KeepPolicy) { // not required
		return nil
	}

	if m.KeepPolicy != nil {
		if err := m.KeepPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keep_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keep_policy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateLastExecuteStatus(formats strfmt.Registry) error {

	if err := validate.Required("last_execute_status", "body", m.LastExecuteStatus); err != nil {
		return err
	}

	if err := validate.Required("last_execute_status", "body", m.LastExecuteStatus); err != nil {
		return err
	}

	if m.LastExecuteStatus != nil {
		if err := m.LastExecuteStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_execute_status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlan) validateStatus(formats strfmt.Registry) error {

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

func (m *BackupPlan) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	for i := 0; i < len(m.Vms); i++ {
		if swag.IsZero(m.Vms[i]) { // not required
			continue
		}

		if m.Vms[i] != nil {
			if err := m.Vms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup plan based on the context it is used
func (m *BackupPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupDelayOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupPlanExecutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupRestorePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupStoreRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullTimePoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalTimePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalWeekdays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeepPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastExecuteStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlan) contextValidateBackupDelayOption(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupDelayOption != nil {
		if err := m.BackupDelayOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_delay_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_delay_option")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateBackupPlanExecutions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupPlanExecutions); i++ {

		if m.BackupPlanExecutions[i] != nil {
			if err := m.BackupPlanExecutions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_plan_executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_plan_executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlan) contextValidateBackupRestorePoints(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupPlan) contextValidateBackupService(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupPlan) contextValidateBackupStoreRepository(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupPlan) contextValidateDeleteStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteStrategy != nil {
		if err := m.DeleteStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delete_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delete_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupPlan) contextValidateFullPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.FullPeriod != nil {
		if err := m.FullPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateFullTimePoint(ctx context.Context, formats strfmt.Registry) error {

	if m.FullTimePoint != nil {
		if err := m.FullTimePoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full_time_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full_time_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateIncrementalPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.IncrementalPeriod != nil {
		if err := m.IncrementalPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateIncrementalTimePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncrementalTimePoints); i++ {

		if m.IncrementalTimePoints[i] != nil {
			if err := m.IncrementalTimePoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incremental_time_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incremental_time_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlan) contextValidateIncrementalWeekdays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncrementalWeekdays); i++ {

		if err := m.IncrementalWeekdays[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental_weekdays" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental_weekdays" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupPlan) contextValidateKeepPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.KeepPolicy != nil {
		if err := m.KeepPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keep_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keep_policy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateLastExecuteStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LastExecuteStatus != nil {
		if err := m.LastExecuteStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_execute_status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlan) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupPlan) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vms); i++ {

		if m.Vms[i] != nil {
			if err := m.Vms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlan) UnmarshalBinary(b []byte) error {
	var res BackupPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
