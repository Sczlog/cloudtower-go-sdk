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
)

// EverouteClusterWhereInput everoute cluster where input
//
// swagger:model EverouteClusterWhereInput
type EverouteClusterWhereInput struct {

	// a n d
	AND []*EverouteClusterWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*EverouteClusterWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*EverouteClusterWhereInput `json:"OR,omitempty"`

	// agent elf clusters every
	AgentElfClustersEvery *ClusterWhereInput `json:"agent_elf_clusters_every,omitempty"`

	// agent elf clusters none
	AgentElfClustersNone *ClusterWhereInput `json:"agent_elf_clusters_none,omitempty"`

	// agent elf clusters some
	AgentElfClustersSome *ClusterWhereInput `json:"agent_elf_clusters_some,omitempty"`

	// agent elf vdses every
	AgentElfVdsesEvery *VdsWhereInput `json:"agent_elf_vdses_every,omitempty"`

	// agent elf vdses none
	AgentElfVdsesNone *VdsWhereInput `json:"agent_elf_vdses_none,omitempty"`

	// agent elf vdses some
	AgentElfVdsesSome *VdsWhereInput `json:"agent_elf_vdses_some,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// global default action
	GlobalDefaultAction *GlobalPolicyAction `json:"global_default_action,omitempty"`

	// global default action in
	GlobalDefaultActionIn []GlobalPolicyAction `json:"global_default_action_in,omitempty"`

	// global default action not
	GlobalDefaultActionNot *GlobalPolicyAction `json:"global_default_action_not,omitempty"`

	// global default action not in
	GlobalDefaultActionNotIn []GlobalPolicyAction `json:"global_default_action_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// installed
	Installed *bool `json:"installed,omitempty"`

	// installed not
	InstalledNot *bool `json:"installed_not,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// phase
	Phase *EverouteClusterPhase `json:"phase,omitempty"`

	// phase in
	PhaseIn []EverouteClusterPhase `json:"phase_in,omitempty"`

	// phase not
	PhaseNot *EverouteClusterPhase `json:"phase_not,omitempty"`

	// phase not in
	PhaseNotIn []EverouteClusterPhase `json:"phase_not_in,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	// version contains
	VersionContains *string `json:"version_contains,omitempty"`

	// version ends with
	VersionEndsWith *string `json:"version_ends_with,omitempty"`

	// version gt
	VersionGt *string `json:"version_gt,omitempty"`

	// version gte
	VersionGte *string `json:"version_gte,omitempty"`

	// version in
	VersionIn []string `json:"version_in,omitempty"`

	// version lt
	VersionLt *string `json:"version_lt,omitempty"`

	// version lte
	VersionLte *string `json:"version_lte,omitempty"`

	// version not
	VersionNot *string `json:"version_not,omitempty"`

	// version not contains
	VersionNotContains *string `json:"version_not_contains,omitempty"`

	// version not ends with
	VersionNotEndsWith *string `json:"version_not_ends_with,omitempty"`

	// version not in
	VersionNotIn []string `json:"version_not_in,omitempty"`

	// version not starts with
	VersionNotStartsWith *string `json:"version_not_starts_with,omitempty"`

	// version starts with
	VersionStartsWith *string `json:"version_starts_with,omitempty"`
}

// Validate validates this everoute cluster where input
func (m *EverouteClusterWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfClustersEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfClustersNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfClustersSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfVdsesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfVdsesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfVdsesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalDefaultAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalDefaultActionIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalDefaultActionNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalDefaultActionNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteClusterWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfClustersEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfClustersEvery) { // not required
		return nil
	}

	if m.AgentElfClustersEvery != nil {
		if err := m.AgentElfClustersEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfClustersNone(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfClustersNone) { // not required
		return nil
	}

	if m.AgentElfClustersNone != nil {
		if err := m.AgentElfClustersNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfClustersSome(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfClustersSome) { // not required
		return nil
	}

	if m.AgentElfClustersSome != nil {
		if err := m.AgentElfClustersSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfVdsesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfVdsesEvery) { // not required
		return nil
	}

	if m.AgentElfVdsesEvery != nil {
		if err := m.AgentElfVdsesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_every")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfVdsesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfVdsesNone) { // not required
		return nil
	}

	if m.AgentElfVdsesNone != nil {
		if err := m.AgentElfVdsesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_none")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateAgentElfVdsesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfVdsesSome) { // not required
		return nil
	}

	if m.AgentElfVdsesSome != nil {
		if err := m.AgentElfVdsesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_some")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *EverouteClusterWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateGlobalDefaultAction(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalDefaultAction) { // not required
		return nil
	}

	if m.GlobalDefaultAction != nil {
		if err := m.GlobalDefaultAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateGlobalDefaultActionIn(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalDefaultActionIn) { // not required
		return nil
	}

	for i := 0; i < len(m.GlobalDefaultActionIn); i++ {

		if err := m.GlobalDefaultActionIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validateGlobalDefaultActionNot(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalDefaultActionNot) { // not required
		return nil
	}

	if m.GlobalDefaultActionNot != nil {
		if err := m.GlobalDefaultActionNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validateGlobalDefaultActionNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalDefaultActionNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.GlobalDefaultActionNotIn); i++ {

		if err := m.GlobalDefaultActionNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validatePhaseIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PhaseIn); i++ {

		if err := m.PhaseIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) validatePhaseNot(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseNot) { // not required
		return nil
	}

	if m.PhaseNot != nil {
		if err := m.PhaseNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) validatePhaseNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PhaseNotIn); i++ {

		if err := m.PhaseNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this everoute cluster where input based on the context it is used
func (m *EverouteClusterWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfClustersEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfClustersNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfClustersSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfVdsesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfVdsesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfVdsesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalDefaultAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalDefaultActionIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalDefaultActionNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalDefaultActionNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfClustersEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfClustersEvery != nil {
		if err := m.AgentElfClustersEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfClustersNone(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfClustersNone != nil {
		if err := m.AgentElfClustersNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfClustersSome(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfClustersSome != nil {
		if err := m.AgentElfClustersSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfVdsesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfVdsesEvery != nil {
		if err := m.AgentElfVdsesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_every")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfVdsesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfVdsesNone != nil {
		if err := m.AgentElfVdsesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_none")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateAgentElfVdsesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.AgentElfVdsesSome != nil {
		if err := m.AgentElfVdsesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent_elf_vdses_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent_elf_vdses_some")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EverouteClusterWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateGlobalDefaultAction(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalDefaultAction != nil {
		if err := m.GlobalDefaultAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateGlobalDefaultActionIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GlobalDefaultActionIn); i++ {

		if err := m.GlobalDefaultActionIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateGlobalDefaultActionNot(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalDefaultActionNot != nil {
		if err := m.GlobalDefaultActionNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidateGlobalDefaultActionNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GlobalDefaultActionNotIn); i++ {

		if err := m.GlobalDefaultActionNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_default_action_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_default_action_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidatePhaseIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhaseIn); i++ {

		if err := m.PhaseIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidatePhaseNot(ctx context.Context, formats strfmt.Registry) error {

	if m.PhaseNot != nil {
		if err := m.PhaseNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteClusterWhereInput) contextValidatePhaseNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhaseNotIn); i++ {

		if err := m.PhaseNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EverouteClusterWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EverouteClusterWhereInput) UnmarshalBinary(b []byte) error {
	var res EverouteClusterWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
