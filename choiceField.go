package forms

type choiceField struct {
	field
}

func (f *choiceField) Validate(val interface{}) []string {
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

func (f *choiceField) Assign(val interface{}) error {
	switch val.(type) {
	case string, float64:
		f.value = val
		return nil
	}
	return typeMismatchError
}

func ChoiceField(name string, required bool, choices []interface{}, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		for _, choice := range choices {
			if val == choice {
				return nil
			}
		}
		return T("Value should be one of %v", choices)
	}
	var f = &choiceField{
		field{
			name:     name,
			required: required,
			ftype:    "String or Number",
			vs:       []Validator{validator},
		},
	}
	f.vs = append(f.vs, vs...)
	return f
}
