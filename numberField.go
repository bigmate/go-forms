package forms

import (
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type floatField struct {
	field
}

func (f *floatField) Assign(val interface{}) error {
	switch val.(type) {
	case string:
		var converted, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return typeMismatchError
		}
		f.value = converted
	case float64:
		f.value = val
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
	if f.Assign(val) != nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errors
	}
	return f.runValidators(lc, errors)
}

func FloatField(name string, required bool, vs ...Validator) Field {
	return &floatField{
		field{
			name:     name,
			required: required,
			ftype:    "Float",
			vs:       vs,
		},
	}
}

type numberField struct {
	field
}

func (f *numberField) Assign(val interface{}) error {
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
	if f.Assign(val) != nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errors
	}
	return f.runValidators(lc, errors)
}

func NumberField(name string, required bool, vs ...Validator) Field {
	return &numberField{
		field{
			name:     name,
			required: required,
			ftype:    "Number",
			vs:       vs,
		},
	}
}
