package forms

import "strconv"

type boolField struct {
	field
}

func (f *boolField) Assign(val interface{}) error {
	switch val.(type) {
	case bool:
		f.value = val
	case string:
		var v, err = strconv.ParseBool(val.(string))
		if err != nil {
			return typeMismatchError
		}
		f.value = v
	case float64:
		var v = val.(float64)
		if v <= 0 {
			f.value = false
		} else {
			f.value = true
		}
	default:
		return typeMismatchError
	}
	return nil
}

func (f *boolField) Validate(val interface{}) []string {
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

func BoolField(name string, required bool, vs ...Validator) Field {
	return &boolField{
		field{
			name:     name,
			required: required,
			ftype:    "Boolean or Number",
			vs:       vs,
		},
	}
}
