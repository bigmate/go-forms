package forms

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ArrayValidator func(lc *i18n.Localizer, val []interface{}) error

type arrayField struct {
	field
	value []interface{}
	vs    []ArrayValidator
}

func (f *arrayField) Value() interface{} {
	return f.value
}

func (f *arrayField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *arrayField) set(val interface{}) error {
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
	if f.set(val) != nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errors
	}
	return f.runValidators(lc, errors)
}

func ArrayField(name string, required bool, vs ...ArrayValidator) Field {
	return &arrayField{
		field: field{
			name:     name,
			required: required,
			ftype:    "Array",
		},
		vs: vs,
	}
}
