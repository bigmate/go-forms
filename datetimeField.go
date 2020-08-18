package forms

import "time"

type datetimeField struct {
	field
}

func (f *datetimeField) Assign(val interface{}) error {
	switch val.(type) {
	case string:
		var v = val.(string)
		if v == "now" {
			f.value = time.Now()
			return nil
		}
		var t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return conversionError
		}
		f.value = t
		return nil
	case float64:
		var v = val.(float64)
		if v <= 0 {
			return conversionError
		}
		f.value = time.Unix(int64(v), 0)
		return nil
	}
	return conversionError
}

func (f *datetimeField) Validate(val interface{}) []string {
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

func DateTimeField(name string, required bool, vs ...Validator) Field {
	var f = &datetimeField{
		field{
			name:     name,
			required: required,
			ftype:    "String or Number",
			vs:       vs,
		},
	}
	return f
}
