package forms

import (
	"strconv"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TimeValidator func(lc *i18n.Localizer, val time.Time) error

type datetimeField struct {
	field
	value time.Time
	vs    []TimeValidator
}

func (f *datetimeField) Value() interface{} {
	return f.value
}

func (f *datetimeField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *datetimeField) set(val interface{}) error {
	switch value := val.(type) {
	case string:
		if value == "now" {
			f.value = time.Now()
		} else {
			var t, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return typeMismatchError
			}
			f.value = t
		}
	case float64:
		if value <= 0 {
			return typeMismatchError
		}
		f.value = time.Unix(int64(value), 0)
	default:
		return typeMismatchError
	}
	f.bnd = true
	return nil
}

func (f *datetimeField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func DateTimeField(name string, required bool, vs ...TimeValidator) Field {
	return &datetimeField{
		field: field{
			name:     name,
			required: required,
			ftype:    "String or Number",
		},
		vs: vs,
	}
}

type DurationValidator func(lc *i18n.Localizer, val time.Duration) error

type durationField struct {
	field
	value time.Duration
	vs    []DurationValidator
}

func (f *durationField) Value() interface{} {
	return f.value
}

func (f *durationField) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		if err := validator(lc, f.value); err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}

func (f *durationField) set(val interface{}) error {
	switch value := val.(type) {
	case string:
		var converted, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return typeMismatchError
		}
		f.value = time.Duration(converted)
	case float64:
		f.value = time.Duration(value)
	default:
		return typeMismatchError
	}
	f.bnd = true
	return nil
}

func (f *durationField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func DurationField(name string, required bool, vs ...DurationValidator) Field {
	return &durationField{
		field: field{
			name:     name,
			required: required,
			ftype:    "Number",
		},
		vs: vs,
	}
}
