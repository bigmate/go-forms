package forms

type charField struct {
	field
}

func (f *charField) Assign(val interface{}) error {
	if _, ok := val.(string); !ok {
		return typeMismatchError
	}
	f.value = val
	return nil
}

func (f *charField) Validate(val interface{}) []string {
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

func CharField(name string, required bool, minLen, maxLen int, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if s := val.(string); len(s) < minLen || len(s) > maxLen {
			return T("length should be between %v and %v", minLen, maxLen)
		}
		return nil
	}
	var f = &charField{
		field{
			name:     name,
			required: required,
			ftype:    "String",
			vs:       []Validator{validator},
		},
	}
	f.vs = append(f.vs, vs...)
	return f
}
