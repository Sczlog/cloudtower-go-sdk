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

// ContentLibraryImage content library image
//
// swagger:model ContentLibraryImage
type ContentLibraryImage struct {

	// clusters
	Clusters []*NestedCluster `json:"clusters,omitempty"`

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// description
	// Required: true
	Description *string `json:"description"`

	// elf image uuids
	// Required: true
	ElfImageUuids []string `json:"elf_image_uuids"`

	// elf images
	ElfImages []*NestedElfImage `json:"elf_images,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// path
	// Required: true
	Path *string `json:"path"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// vm disks
	VMDisks []*NestedVMDisk `json:"vm_disks,omitempty"`

	// vm snapshots
	VMSnapshots []*NestedVMSnapshot `json:"vm_snapshots,omitempty"`

	// vm templates
	VMTemplates []*NestedVMTemplate `json:"vm_templates,omitempty"`
}

// Validate validates this content library image
func (m *ContentLibraryImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImageUuids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryImage) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateElfImageUuids(formats strfmt.Registry) error {

	if err := validate.Required("elf_image_uuids", "body", m.ElfImageUuids); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateElfImages(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfImages) { // not required
		return nil
	}

	for i := 0; i < len(m.ElfImages); i++ {
		if swag.IsZero(m.ElfImages[i]) { // not required
			continue
		}

		if m.ElfImages[i] != nil {
			if err := m.ElfImages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elf_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elf_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *ContentLibraryImage) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryImage) validateVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.VMDisks); i++ {
		if swag.IsZero(m.VMDisks[i]) { // not required
			continue
		}

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) validateVMSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.VMSnapshots); i++ {
		if swag.IsZero(m.VMSnapshots[i]) { // not required
			continue
		}

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) validateVMTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VMTemplates); i++ {
		if swag.IsZero(m.VMTemplates[i]) { // not required
			continue
		}

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this content library image based on the context it is used
func (m *ContentLibraryImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryImage) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) contextValidateElfImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ElfImages); i++ {

		if m.ElfImages[i] != nil {
			if err := m.ElfImages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elf_images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elf_images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryImage) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMDisks); i++ {

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) contextValidateVMSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMSnapshots); i++ {

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryImage) contextValidateVMTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMTemplates); i++ {

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryImage) UnmarshalBinary(b []byte) error {
	var res ContentLibraryImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}