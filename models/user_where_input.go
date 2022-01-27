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

// UserWhereInput user where input
// Example: {"AND":"UserWhereInput[]","NOT":"UserWhereInput[]","OR":"UserWhereInput[]","email_address":"string","email_address_contains":"string","email_address_ends_with":"string","email_address_gt":"string","email_address_gte":"string","email_address_in":["string"],"email_address_lt":"string","email_address_lte":"string","email_address_not":"string","email_address_not_contains":"string","email_address_not_ends_with":"string","email_address_not_in":["string"],"email_address_not_starts_with":"string","email_address_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"ldap_dn":"string","ldap_dn_contains":"string","ldap_dn_ends_with":"string","ldap_dn_gt":"string","ldap_dn_gte":"string","ldap_dn_in":["string"],"ldap_dn_lt":"string","ldap_dn_lte":"string","ldap_dn_not":"string","ldap_dn_not_contains":"string","ldap_dn_not_ends_with":"string","ldap_dn_not_in":["string"],"ldap_dn_not_starts_with":"string","ldap_dn_starts_with":"string","mobile_phone":"string","mobile_phone_contains":"string","mobile_phone_ends_with":"string","mobile_phone_gt":"string","mobile_phone_gte":"string","mobile_phone_in":["string"],"mobile_phone_lt":"string","mobile_phone_lte":"string","mobile_phone_not":"string","mobile_phone_not_contains":"string","mobile_phone_not_ends_with":"string","mobile_phone_not_in":["string"],"mobile_phone_not_starts_with":"string","mobile_phone_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","password":"string","password_contains":"string","password_ends_with":"string","password_gt":"string","password_gte":"string","password_in":["string"],"password_lt":"string","password_lte":"string","password_not":"string","password_not_contains":"string","password_not_ends_with":"string","password_not_in":["string"],"password_not_starts_with":"string","password_starts_with":"string","role":"ADMIN","role_in":["ADMIN"],"role_not":"ADMIN","role_not_in":["ADMIN"],"roles_every":"UserRoleNextWhereInput","roles_none":"UserRoleNextWhereInput","roles_some":"UserRoleNextWhereInput","source":"LDAP","source_in":["LDAP"],"source_not":"LDAP","source_not_in":["LDAP"],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string"}
//
// swagger:model UserWhereInput
type UserWhereInput struct {

	// a n d
	AND []*UserWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*UserWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*UserWhereInput `json:"OR,omitempty"`

	// email address
	EmailAddress *string `json:"email_address,omitempty"`

	// email address contains
	EmailAddressContains *string `json:"email_address_contains,omitempty"`

	// email address ends with
	EmailAddressEndsWith *string `json:"email_address_ends_with,omitempty"`

	// email address gt
	EmailAddressGt *string `json:"email_address_gt,omitempty"`

	// email address gte
	EmailAddressGte *string `json:"email_address_gte,omitempty"`

	// email address in
	EmailAddressIn []string `json:"email_address_in,omitempty"`

	// email address lt
	EmailAddressLt *string `json:"email_address_lt,omitempty"`

	// email address lte
	EmailAddressLte *string `json:"email_address_lte,omitempty"`

	// email address not
	EmailAddressNot *string `json:"email_address_not,omitempty"`

	// email address not contains
	EmailAddressNotContains *string `json:"email_address_not_contains,omitempty"`

	// email address not ends with
	EmailAddressNotEndsWith *string `json:"email_address_not_ends_with,omitempty"`

	// email address not in
	EmailAddressNotIn []string `json:"email_address_not_in,omitempty"`

	// email address not starts with
	EmailAddressNotStartsWith *string `json:"email_address_not_starts_with,omitempty"`

	// email address starts with
	EmailAddressStartsWith *string `json:"email_address_starts_with,omitempty"`

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

	// internal
	Internal *bool `json:"internal,omitempty"`

	// internal not
	InternalNot *bool `json:"internal_not,omitempty"`

	// ldap dn
	LdapDn *string `json:"ldap_dn,omitempty"`

	// ldap dn contains
	LdapDnContains *string `json:"ldap_dn_contains,omitempty"`

	// ldap dn ends with
	LdapDnEndsWith *string `json:"ldap_dn_ends_with,omitempty"`

	// ldap dn gt
	LdapDnGt *string `json:"ldap_dn_gt,omitempty"`

	// ldap dn gte
	LdapDnGte *string `json:"ldap_dn_gte,omitempty"`

	// ldap dn in
	LdapDnIn []string `json:"ldap_dn_in,omitempty"`

	// ldap dn lt
	LdapDnLt *string `json:"ldap_dn_lt,omitempty"`

	// ldap dn lte
	LdapDnLte *string `json:"ldap_dn_lte,omitempty"`

	// ldap dn not
	LdapDnNot *string `json:"ldap_dn_not,omitempty"`

	// ldap dn not contains
	LdapDnNotContains *string `json:"ldap_dn_not_contains,omitempty"`

	// ldap dn not ends with
	LdapDnNotEndsWith *string `json:"ldap_dn_not_ends_with,omitempty"`

	// ldap dn not in
	LdapDnNotIn []string `json:"ldap_dn_not_in,omitempty"`

	// ldap dn not starts with
	LdapDnNotStartsWith *string `json:"ldap_dn_not_starts_with,omitempty"`

	// ldap dn starts with
	LdapDnStartsWith *string `json:"ldap_dn_starts_with,omitempty"`

	// mobile phone
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// mobile phone contains
	MobilePhoneContains *string `json:"mobile_phone_contains,omitempty"`

	// mobile phone ends with
	MobilePhoneEndsWith *string `json:"mobile_phone_ends_with,omitempty"`

	// mobile phone gt
	MobilePhoneGt *string `json:"mobile_phone_gt,omitempty"`

	// mobile phone gte
	MobilePhoneGte *string `json:"mobile_phone_gte,omitempty"`

	// mobile phone in
	MobilePhoneIn []string `json:"mobile_phone_in,omitempty"`

	// mobile phone lt
	MobilePhoneLt *string `json:"mobile_phone_lt,omitempty"`

	// mobile phone lte
	MobilePhoneLte *string `json:"mobile_phone_lte,omitempty"`

	// mobile phone not
	MobilePhoneNot *string `json:"mobile_phone_not,omitempty"`

	// mobile phone not contains
	MobilePhoneNotContains *string `json:"mobile_phone_not_contains,omitempty"`

	// mobile phone not ends with
	MobilePhoneNotEndsWith *string `json:"mobile_phone_not_ends_with,omitempty"`

	// mobile phone not in
	MobilePhoneNotIn []string `json:"mobile_phone_not_in,omitempty"`

	// mobile phone not starts with
	MobilePhoneNotStartsWith *string `json:"mobile_phone_not_starts_with,omitempty"`

	// mobile phone starts with
	MobilePhoneStartsWith *string `json:"mobile_phone_starts_with,omitempty"`

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

	// password
	Password *string `json:"password,omitempty"`

	// password contains
	PasswordContains *string `json:"password_contains,omitempty"`

	// password ends with
	PasswordEndsWith *string `json:"password_ends_with,omitempty"`

	// password gt
	PasswordGt *string `json:"password_gt,omitempty"`

	// password gte
	PasswordGte *string `json:"password_gte,omitempty"`

	// password in
	PasswordIn []string `json:"password_in,omitempty"`

	// password lt
	PasswordLt *string `json:"password_lt,omitempty"`

	// password lte
	PasswordLte *string `json:"password_lte,omitempty"`

	// password not
	PasswordNot *string `json:"password_not,omitempty"`

	// password not contains
	PasswordNotContains *string `json:"password_not_contains,omitempty"`

	// password not ends with
	PasswordNotEndsWith *string `json:"password_not_ends_with,omitempty"`

	// password not in
	PasswordNotIn []string `json:"password_not_in,omitempty"`

	// password not starts with
	PasswordNotStartsWith *string `json:"password_not_starts_with,omitempty"`

	// password starts with
	PasswordStartsWith *string `json:"password_starts_with,omitempty"`

	// role
	Role *UserRole `json:"role,omitempty"`

	// role in
	RoleIn []UserRole `json:"role_in,omitempty"`

	// role not
	RoleNot *UserRole `json:"role_not,omitempty"`

	// role not in
	RoleNotIn []UserRole `json:"role_not_in,omitempty"`

	// roles every
	RolesEvery *UserRoleNextWhereInput `json:"roles_every,omitempty"`

	// roles none
	RolesNone *UserRoleNextWhereInput `json:"roles_none,omitempty"`

	// roles some
	RolesSome *UserRoleNextWhereInput `json:"roles_some,omitempty"`

	// source
	Source *UserSource `json:"source,omitempty"`

	// source in
	SourceIn []UserSource `json:"source_in,omitempty"`

	// source not
	SourceNot *UserSource `json:"source_not,omitempty"`

	// source not in
	SourceNotIn []UserSource `json:"source_not_in,omitempty"`

	// username
	Username *string `json:"username,omitempty"`

	// username contains
	UsernameContains *string `json:"username_contains,omitempty"`

	// username ends with
	UsernameEndsWith *string `json:"username_ends_with,omitempty"`

	// username gt
	UsernameGt *string `json:"username_gt,omitempty"`

	// username gte
	UsernameGte *string `json:"username_gte,omitempty"`

	// username in
	UsernameIn []string `json:"username_in,omitempty"`

	// username lt
	UsernameLt *string `json:"username_lt,omitempty"`

	// username lte
	UsernameLte *string `json:"username_lte,omitempty"`

	// username not
	UsernameNot *string `json:"username_not,omitempty"`

	// username not contains
	UsernameNotContains *string `json:"username_not_contains,omitempty"`

	// username not ends with
	UsernameNotEndsWith *string `json:"username_not_ends_with,omitempty"`

	// username not in
	UsernameNotIn []string `json:"username_not_in,omitempty"`

	// username not starts with
	UsernameNotStartsWith *string `json:"username_not_starts_with,omitempty"`

	// username starts with
	UsernameStartsWith *string `json:"username_starts_with,omitempty"`
}

// Validate validates this user where input
func (m *UserWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *UserWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *UserWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *UserWhereInput) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateRoleIn(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleIn) { // not required
		return nil
	}

	for i := 0; i < len(m.RoleIn); i++ {

		if err := m.RoleIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) validateRoleNot(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleNot) { // not required
		return nil
	}

	if m.RoleNot != nil {
		if err := m.RoleNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateRoleNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.RoleNotIn); i++ {

		if err := m.RoleNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) validateRolesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.RolesEvery) { // not required
		return nil
	}

	if m.RolesEvery != nil {
		if err := m.RolesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_every")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateRolesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.RolesNone) { // not required
		return nil
	}

	if m.RolesNone != nil {
		if err := m.RolesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_none")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateRolesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.RolesSome) { // not required
		return nil
	}

	if m.RolesSome != nil {
		if err := m.RolesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_some")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateSourceIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceIn); i++ {

		if err := m.SourceIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) validateSourceNot(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceNot) { // not required
		return nil
	}

	if m.SourceNot != nil {
		if err := m.SourceNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) validateSourceNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceNotIn); i++ {

		if err := m.SourceNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this user where input based on the context it is used
func (m *UserWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRolesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRolesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRolesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserWhereInput) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateRoleIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoleIn); i++ {

		if err := m.RoleIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) contextValidateRoleNot(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleNot != nil {
		if err := m.RoleNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateRoleNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoleNotIn); i++ {

		if err := m.RoleNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) contextValidateRolesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.RolesEvery != nil {
		if err := m.RolesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_every")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateRolesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.RolesNone != nil {
		if err := m.RolesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_none")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateRolesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.RolesSome != nil {
		if err := m.RolesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles_some")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateSourceIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceIn); i++ {

		if err := m.SourceIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserWhereInput) contextValidateSourceNot(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceNot != nil {
		if err := m.SourceNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserWhereInput) contextValidateSourceNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceNotIn); i++ {

		if err := m.SourceNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserWhereInput) UnmarshalBinary(b []byte) error {
	var res UserWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
