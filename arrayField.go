package forms

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type arrayField struct {
	field
}

func (f *arrayField) Assign(val interface{}) error {
	if value, ok := val.([]interface{}); ok {
		f.value = value
		f.bound = true
		return nil
	}
	return typeMismatchError
}

func (f *arrayField) Validate(lc *i18n.Localizer, val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: fieldRequired,
		}))
		return errors
	}
	if f.Assign(val) != nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errors
	}
	return f.runValidators(lc, errors)
}

func ArrayField(name string, required bool, vs ...Validator) Field {
	return &arrayField{field{
		name:     name,
		required: required,
		ftype:    "Array",
		vs:       vs,
	}}
}
