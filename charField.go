package forms

import (
	"errors"
	"regexp"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type charField struct {
	field
}

func (f *charField) Assign(val interface{}) error {
	if _, ok := val.(string); !ok {
		return typeMismatchError
	}
	f.value = val
	f.bound = true
	return nil
}

func (f *charField) Validate(lc *i18n.Localizer, val interface{}) []string {
	var errs = make([]string, 0)
	if !f.required && val == nil {
		return errs
	}
	if f.required && val == nil {
		errs = append(errs, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fieldRequired,
		}))
		return errs
	}
	if f.Assign(val) != nil {
		errs = append(errs, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errs
	}
	return f.runValidators(lc, errs)
}

func CharField(name string, required bool, vs ...Validator) Field {
	return &charField{
		field{
			name:     name,
			required: required,
			ftype:    "String",
			vs:       vs,
		},
	}
}

var emailRE = regexp.MustCompile("^\\S{4,20}@\\S{2,15}\\.\\S{2,6}$")

func EmailField(name string, required bool, vs ...Validator) Field {
	var validator = func(lc *i18n.Localizer, val interface{}) error {
		if !emailRE.Match([]byte(val.(string))) {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "invalid_email",
				DefaultMessage: &i18n.Message{
					ID:    "invalid_email",
					Other: "invalid email",
				}}))
		}
		return nil
	}
	var f = &charField{
		field{
			name:     name,
			required: required,
			ftype:    "String",
			vs:       []Validator{validator},
		},
	}
	f.vs = append(f.vs, vs...)
	return f
}

var uuidRE = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func UUIDField(name string, required bool, vs ...Validator) Field {
	var validator = func(lc *i18n.Localizer, val interface{}) error {
		if !uuidRE.Match([]byte(val.(string))) {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "invalid_uuid",
				DefaultMessage: &i18n.Message{
					ID:    "invalid_uuid",
					Other: "invalid uuid",
				}}))
		}
		return nil
	}
	var f = &charField{
		field{
			name:     name,
			required: required,
			ftype:    "String",
			vs:       []Validator{validator},
		},
	}
	f.vs = append(f.vs, vs...)
	return f
}
