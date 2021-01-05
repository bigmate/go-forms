package forms

import (
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type FloatValidator func(lc *i18n.Localizer, val float64) error

type floatField struct {
	field
	value float64
	vs    []FloatValidator
}

func (f *floatField) Value() interface{} {
	return f.value
}

func (f *floatField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *floatField) set(val interface{}) error {
	switch value := val.(type) {
	case string:
		var converted, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return typeMismatchError
		}
		f.value = converted
	case float64:
		f.value = value
	default:
		return typeMismatchError
	}
	f.bound = true
	return nil
}

func (f *floatField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func FloatField(name string, required bool, vs ...FloatValidator) Field {
	return &floatField{
		field: field{
			name:     name,
			required: required,
			ftype:    "Float",
		},
		vs: vs,
	}
}

type NumValidator func(lc *i18n.Localizer, val int64) error

type numberField struct {
	field
	value int64
	vs    []NumValidator
}

func (f *numberField) Value() interface{} {
	return f.value
}

func (f *numberField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *numberField) set(val interface{}) error {
	switch value := val.(type) {
	case string:
		var converted, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return typeMismatchError
		}
		f.value = converted
	case float64:
		f.value = int64(value)
	default:
		return typeMismatchError
	}
	f.bound = true
	return nil
}

func (f *numberField) Validate(lc *i18n.Localizer, val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{MessageID: fieldRequired}))
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

func NumberField(name string, required bool, vs ...NumValidator) Field {
	return &numberField{
		field: field{
			name:     name,
			required: required,
			ftype:    "Number",
		},
		vs: vs,
	}
}
