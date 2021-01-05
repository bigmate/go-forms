package forms

import (
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type BoolValidator func(lc *i18n.Localizer, val bool) error

type boolField struct {
	field
	value bool
	vs    []BoolValidator
}

func (f *boolField) Value() interface{} {
	return f.value
}

func (f *boolField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *boolField) set(val interface{}) error {
	switch value := val.(type) {
	case bool:
		f.value = value
	case string:
		var v, err = strconv.ParseBool(value)
		if err != nil {
			return typeMismatchError
		}
		f.value = v
	case float64:
		if value <= 0 {
			f.value = false
		} else {
			f.value = true
		}
	default:
		return typeMismatchError
	}
	f.bound = true
	return nil
}

func (f *boolField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func BoolField(name string, required bool, vs ...BoolValidator) Field {
	return &boolField{
		field: field{
			name:     name,
			required: required,
			ftype:    "Boolean or Number",
		},
		vs: vs,
	}
}
