package forms

type arrayField struct {
	field
}

func (f *arrayField) Assign(val interface{}) error {
	if value, ok := val.([]interface{}); ok {
		f.value = value
		f.bound = true
		return nil
	}
	return typeMismatchError
}

func (f *arrayField) Validate(val interface{}) []string {
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

func ArrayField(name string, required bool, vs ...Validator) Field {
	return &arrayField{field{
		name:     name,
		required: required,
		ftype:    "Array",
		vs:       vs,
	}}
}
