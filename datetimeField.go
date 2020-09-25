package forms

import (
	"strconv"
	"time"
)

type datetimeField struct {
	field
}

func (f *datetimeField) Assign(val interface{}) error {
	switch val.(type) {
	case string:
		var v = val.(string)
		if v == "now" {
			f.value = time.Now()
		} else {
			var t, err = time.Parse(time.RFC3339, v)
			if err != nil {
				return typeMismatchError
			}
			f.value = t
		}
	case float64:
		var v = val.(float64)
		if v <= 0 {
			return typeMismatchError
		}
		f.value = time.Unix(int64(v), 0)
	default:
		return typeMismatchError
	}
	return nil
}

func (f *datetimeField) Validate(val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, t(fieldRequired))
		return errors
	}
	if f.Assign(val) != nil {
		errors = append(errors, t(typeMismatch, f.ftype))
		return errors
	}
	return f.runValidators(errors)
}

func DateTimeField(name string, required bool, vs ...Validator) Field {
	return &datetimeField{
		field{
			name:     name,
			required: required,
			ftype:    "String or Number",
			vs:       vs,
		},
	}
}

type durationField struct {
	field
}

func (f *durationField) Assign(val interface{}) error {
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
	return nil
}

func (f *durationField) Validate(val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, t(fieldRequired))
		return errors
	}
	if f.Assign(val) != nil {
		errors = append(errors, t(typeMismatch, f.ftype))
		return errors
	}
	return f.runValidators(errors)
}

func DurationField(name string, required bool, vs ...Validator) Field {
	return &numberField{
		field{
			name:     name,
			required: required,
			ftype:    "Number",
			vs:       vs,
		},
		nil,
	}
}
