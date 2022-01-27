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

// AlertNotifierWhereInput alert notifier where input
// Example: {"AND":"AlertNotifierWhereInput[]","NOT":"AlertNotifierWhereInput[]","OR":"AlertNotifierWhereInput[]","disabled":false,"disabled_not":false,"email_from":"string","email_from_contains":"string","email_from_ends_with":"string","email_from_gt":"string","email_from_gte":"string","email_from_in":["string"],"email_from_lt":"string","email_from_lte":"string","email_from_not":"string","email_from_not_contains":"string","email_from_not_ends_with":"string","email_from_not_in":["string"],"email_from_not_starts_with":"string","email_from_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","language_code":"EN_US","language_code_in":["EN_US"],"language_code_not":"EN_US","language_code_not_in":["EN_US"],"password":"string","password_contains":"string","password_ends_with":"string","password_gt":"string","password_gte":"string","password_in":["string"],"password_lt":"string","password_lte":"string","password_not":"string","password_not_contains":"string","password_not_ends_with":"string","password_not_in":["string"],"password_not_starts_with":"string","password_starts_with":"string","security_mode":"SSL","security_mode_in":["SSL"],"security_mode_not":"SSL","security_mode_not_in":["SSL"],"smtp_server_host":"string","smtp_server_host_contains":"string","smtp_server_host_ends_with":"string","smtp_server_host_gt":"string","smtp_server_host_gte":"string","smtp_server_host_in":["string"],"smtp_server_host_lt":"string","smtp_server_host_lte":"string","smtp_server_host_not":"string","smtp_server_host_not_contains":"string","smtp_server_host_not_ends_with":"string","smtp_server_host_not_in":["string"],"smtp_server_host_not_starts_with":"string","smtp_server_host_starts_with":"string","smtp_server_port":0,"smtp_server_port_gt":0,"smtp_server_port_gte":0,"smtp_server_port_in":[0],"smtp_server_port_lt":0,"smtp_server_port_lte":0,"smtp_server_port_not":0,"smtp_server_port_not_in":[0],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string"}
//
// swagger:model AlertNotifierWhereInput
type AlertNotifierWhereInput struct {

	// a n d
	AND []*AlertNotifierWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*AlertNotifierWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*AlertNotifierWhereInput `json:"OR,omitempty"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disabled not
	DisabledNot *bool `json:"disabled_not,omitempty"`

	// email from
	EmailFrom *string `json:"email_from,omitempty"`

	// email from contains
	EmailFromContains *string `json:"email_from_contains,omitempty"`

	// email from ends with
	EmailFromEndsWith *string `json:"email_from_ends_with,omitempty"`

	// email from gt
	EmailFromGt *string `json:"email_from_gt,omitempty"`

	// email from gte
	EmailFromGte *string `json:"email_from_gte,omitempty"`

	// email from in
	EmailFromIn []string `json:"email_from_in,omitempty"`

	// email from lt
	EmailFromLt *string `json:"email_from_lt,omitempty"`

	// email from lte
	EmailFromLte *string `json:"email_from_lte,omitempty"`

	// email from not
	EmailFromNot *string `json:"email_from_not,omitempty"`

	// email from not contains
	EmailFromNotContains *string `json:"email_from_not_contains,omitempty"`

	// email from not ends with
	EmailFromNotEndsWith *string `json:"email_from_not_ends_with,omitempty"`

	// email from not in
	EmailFromNotIn []string `json:"email_from_not_in,omitempty"`

	// email from not starts with
	EmailFromNotStartsWith *string `json:"email_from_not_starts_with,omitempty"`

	// email from starts with
	EmailFromStartsWith *string `json:"email_from_starts_with,omitempty"`

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

	// language code
	LanguageCode *NotifierLanguageCode `json:"language_code,omitempty"`

	// language code in
	LanguageCodeIn []NotifierLanguageCode `json:"language_code_in,omitempty"`

	// language code not
	LanguageCodeNot *NotifierLanguageCode `json:"language_code_not,omitempty"`

	// language code not in
	LanguageCodeNotIn []NotifierLanguageCode `json:"language_code_not_in,omitempty"`

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

	// security mode
	SecurityMode *NotifierSecurityMode `json:"security_mode,omitempty"`

	// security mode in
	SecurityModeIn []NotifierSecurityMode `json:"security_mode_in,omitempty"`

	// security mode not
	SecurityModeNot *NotifierSecurityMode `json:"security_mode_not,omitempty"`

	// security mode not in
	SecurityModeNotIn []NotifierSecurityMode `json:"security_mode_not_in,omitempty"`

	// smtp server host
	SMTPServerHost *string `json:"smtp_server_host,omitempty"`

	// smtp server host contains
	SMTPServerHostContains *string `json:"smtp_server_host_contains,omitempty"`

	// smtp server host ends with
	SMTPServerHostEndsWith *string `json:"smtp_server_host_ends_with,omitempty"`

	// smtp server host gt
	SMTPServerHostGt *string `json:"smtp_server_host_gt,omitempty"`

	// smtp server host gte
	SMTPServerHostGte *string `json:"smtp_server_host_gte,omitempty"`

	// smtp server host in
	SMTPServerHostIn []string `json:"smtp_server_host_in,omitempty"`

	// smtp server host lt
	SMTPServerHostLt *string `json:"smtp_server_host_lt,omitempty"`

	// smtp server host lte
	SMTPServerHostLte *string `json:"smtp_server_host_lte,omitempty"`

	// smtp server host not
	SMTPServerHostNot *string `json:"smtp_server_host_not,omitempty"`

	// smtp server host not contains
	SMTPServerHostNotContains *string `json:"smtp_server_host_not_contains,omitempty"`

	// smtp server host not ends with
	SMTPServerHostNotEndsWith *string `json:"smtp_server_host_not_ends_with,omitempty"`

	// smtp server host not in
	SMTPServerHostNotIn []string `json:"smtp_server_host_not_in,omitempty"`

	// smtp server host not starts with
	SMTPServerHostNotStartsWith *string `json:"smtp_server_host_not_starts_with,omitempty"`

	// smtp server host starts with
	SMTPServerHostStartsWith *string `json:"smtp_server_host_starts_with,omitempty"`

	// smtp server port
	SMTPServerPort *int32 `json:"smtp_server_port,omitempty"`

	// smtp server port gt
	SMTPServerPortGt *int32 `json:"smtp_server_port_gt,omitempty"`

	// smtp server port gte
	SMTPServerPortGte *int32 `json:"smtp_server_port_gte,omitempty"`

	// smtp server port in
	SMTPServerPortIn []int32 `json:"smtp_server_port_in,omitempty"`

	// smtp server port lt
	SMTPServerPortLt *int32 `json:"smtp_server_port_lt,omitempty"`

	// smtp server port lte
	SMTPServerPortLte *int32 `json:"smtp_server_port_lte,omitempty"`

	// smtp server port not
	SMTPServerPortNot *int32 `json:"smtp_server_port_not,omitempty"`

	// smtp server port not in
	SMTPServerPortNotIn []int32 `json:"smtp_server_port_not_in,omitempty"`

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

// Validate validates this alert notifier where input
func (m *AlertNotifierWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *AlertNotifierWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *AlertNotifierWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *AlertNotifierWhereInput) validateLanguageCode(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCode) { // not required
		return nil
	}

	if m.LanguageCode != nil {
		if err := m.LanguageCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeNot) { // not required
		return nil
	}

	if m.LanguageCodeNot != nil {
		if err := m.LanguageCodeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityMode) { // not required
		return nil
	}

	if m.SecurityMode != nil {
		if err := m.SecurityMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityModeIn); i++ {

		if err := m.SecurityModeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeNot) { // not required
		return nil
	}

	if m.SecurityModeNot != nil {
		if err := m.SecurityModeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityModeNotIn); i++ {

		if err := m.SecurityModeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this alert notifier where input based on the context it is used
func (m *AlertNotifierWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateLanguageCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertNotifierWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertNotifierWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertNotifierWhereInput) contextValidateLanguageCode(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCode != nil {
		if err := m.LanguageCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCodeNot != nil {
		if err := m.LanguageCodeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityMode(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityMode != nil {
		if err := m.SecurityMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityModeIn); i++ {

		if err := m.SecurityModeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityModeNot != nil {
		if err := m.SecurityModeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityModeNotIn); i++ {

		if err := m.SecurityModeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertNotifierWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertNotifierWhereInput) UnmarshalBinary(b []byte) error {
	var res AlertNotifierWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
