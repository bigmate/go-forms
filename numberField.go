package forms

import (
	"strconv"
)

type numberField struct {
	field
}

func (f *numberField) Assign(val interface{}) error {
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
	return nil
}

func (f *numberField) Validate(val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, t(fieldRequired))
	}
	if f.Assign(val) != nil {
		errors = append(errors, t(typeMismatch, f.ftype))
		return errors
	}
	return f.runValidators(errors)
}

func NumberField(name string, required bool, min, max float64, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if v := val.(float64); v < min || v > max {
			return T("value should be between %v and %v", min, max)
		}
		return nil
	}
	var f = &numberField{
		field{
			name:     name,
			required: required,
			ftype:    "Number",
			vs:       []Validator{validator},
		},
	}
	f.vs = append(f.vs, vs...)
	return f
}
