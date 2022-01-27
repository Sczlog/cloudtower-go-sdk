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

// EverouteCluster everoute cluster
//
// swagger:model EverouteCluster
type EverouteCluster struct {

	// agent elf clusters
	AgentElfClusters []*NestedCluster `json:"agent_elf_clusters,omitempty"`

	// agent elf vdses
	AgentElfVdses []*NestedVds `json:"agent_elf_vdses,omitempty"`

	// controller instances
	// Required: true
	ControllerInstances []*NestedEverouteControllerInstance `json:"controller_instances"`

	// controller template
	// Required: true
	ControllerTemplate *NestedEverouteControllerTemplate `json:"controller_template"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// global default action
	// Required: true
	GlobalDefaultAction *GlobalPolicyAction `json:"global_default_action"`

	// global whitelist
	GlobalWhitelist *NestedEverouteClusterWhitelist `json:"global_whitelist,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// installed
	Installed *bool `json:"installed,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// phase
	Phase *EverouteClusterPhase `json:"phase,omitempty"`

	// status
	// Required: true
	Status *NestedEverouteClusterStatus `json:"status"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this everoute cluster
func (m *EverouteCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentElfClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgentElfVdses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalDefaultAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteCluster) validateAgentElfClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentElfClusters); i++ {
		if swag.IsZero(m.AgentElfClusters[i]) { // not required
			continue
		}

		if m.AgentElfClusters[i] != nil {
			if err := m.AgentElfClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agent_elf_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agent_elf_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) validateAgentElfVdses(formats strfmt.Registry) error {
	if swag.IsZero(m.AgentElfVdses) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentElfVdses); i++ {
		if swag.IsZero(m.AgentElfVdses[i]) { // not required
			continue
		}

		if m.AgentElfVdses[i] != nil {
			if err := m.AgentElfVdses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agent_elf_vdses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agent_elf_vdses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) validateControllerInstances(formats strfmt.Registry) error {

	if err := validate.Required("controller_instances", "body", m.ControllerInstances); err != nil {
		return err
	}

	for i := 0; i < len(m.ControllerInstances); i++ {
		if swag.IsZero(m.ControllerInstances[i]) { // not required
			continue
		}

		if m.ControllerInstances[i] != nil {
			if err := m.ControllerInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controller_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controller_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) validateControllerTemplate(formats strfmt.Registry) error {

	if err := validate.Required("controller_template", "body", m.ControllerTemplate); err != nil {
		return err
	}

	if m.ControllerTemplate != nil {
		if err := m.ControllerTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controller_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controller_template")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteCluster) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *EverouteCluster) validateGlobalDefaultAction(formats strfmt.Registry) error {

	if err := validate.Required("global_default_action", "body", m.GlobalDefaultAction); err != nil {
		return err
	}

	if err := validate.Required("global_default_action", "body", m.GlobalDefaultAction); err != nil {
		return err
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

func (m *EverouteCluster) validateGlobalWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalWhitelist) { // not required
		return nil
	}

	if m.GlobalWhitelist != nil {
		if err := m.GlobalWhitelist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_whitelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_whitelist")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EverouteCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EverouteCluster) validatePhase(formats strfmt.Registry) error {
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

func (m *EverouteCluster) validateStatus(formats strfmt.Registry) error {

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

func (m *EverouteCluster) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this everoute cluster based on the context it is used
func (m *EverouteCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgentElfClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentElfVdses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalDefaultAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteCluster) contextValidateAgentElfClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AgentElfClusters); i++ {

		if m.AgentElfClusters[i] != nil {
			if err := m.AgentElfClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agent_elf_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agent_elf_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) contextValidateAgentElfVdses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AgentElfVdses); i++ {

		if m.AgentElfVdses[i] != nil {
			if err := m.AgentElfVdses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agent_elf_vdses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agent_elf_vdses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) contextValidateControllerInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ControllerInstances); i++ {

		if m.ControllerInstances[i] != nil {
			if err := m.ControllerInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controller_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controller_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteCluster) contextValidateControllerTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ControllerTemplate != nil {
		if err := m.ControllerTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controller_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controller_template")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteCluster) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EverouteCluster) contextValidateGlobalDefaultAction(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EverouteCluster) contextValidateGlobalWhitelist(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalWhitelist != nil {
		if err := m.GlobalWhitelist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global_whitelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global_whitelist")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteCluster) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EverouteCluster) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *EverouteCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EverouteCluster) UnmarshalBinary(b []byte) error {
	var res EverouteCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
