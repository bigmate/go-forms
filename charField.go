package forms

import (
	"errors"
	"regexp"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type CharValidator func(lc *i18n.Localizer, val string) error

type charField struct {
	field
	value string
	vs    []CharValidator
}

func (f *charField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *charField) Value() interface{} {
	return f.value
}

func (f *charField) set(val interface{}) error {
	if value, ok := val.(string); ok {
		f.value = value
		f.bnd = true
		return nil
	}
	return typeMismatchError
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
	if f.set(val) != nil {
		errs = append(errs, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errs
	}
	return f.runValidators(lc, errs)
}

func CharField(name string, required bool, vs ...CharValidator) Field {
	return &charField{
		field: field{
			name:     name,
			required: required,
			ftype:    "String",
		},
		vs: vs,
	}
}

var emailRE = regexp.MustCompile("^\\S{4,20}@\\S{2,15}\\.\\S{2,6}$")

func EmailField(name string, required bool, vs ...CharValidator) Field {
	var validator = func(lc *i18n.Localizer, val string) error {
		if !emailRE.Match([]byte(val)) {
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
		field: field{
			name:     name,
			required: required,
			ftype:    "String",
		},
		vs: []CharValidator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}

var uuidRE = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func UUIDField(name string, required bool, vs ...CharValidator) Field {
	var validator = func(lc *i18n.Localizer, val string) error {
		if !uuidRE.Match([]byte(val)) {
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
		field: field{
			name:     name,
			required: required,
			ftype:    "String",
		},
		vs: []CharValidator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}
