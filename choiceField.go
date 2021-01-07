package forms

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ChoiceValidator func(lc *i18n.Localizer, val interface{}) error

type choiceField struct {
	field
	value interface{}
	vs    []ChoiceValidator
}

func (f *choiceField) Value() interface{} {
	return f.value
}

func (f *choiceField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *choiceField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func (f *choiceField) set(val interface{}) error {
	switch val.(type) {
	case string, float64:
		f.value = val
		f.bnd = true
		return nil
	}
	return typeMismatchError
}

func ChoiceField(name string, required bool, choices []interface{}, vs ...ChoiceValidator) Field {
	var validator = func(lc *i18n.Localizer, val interface{}) error {
		for _, choice := range choices {
			if val == choice {
				return nil
			}
		}
		return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "value_one_of",
			DefaultMessage: &i18n.Message{
				ID:    "value_one_of",
				One:   "Value should be {{.}}",
				Other: "Value should be one of {{.}}",
			},
			TemplateData: choices,
		}))
	}
	var f = &choiceField{
		field: field{
			name:     name,
			required: required,
			ftype:    "String or Number",
		},
		vs: []ChoiceValidator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}
