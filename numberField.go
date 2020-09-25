package forms

type numberField struct {
	field
	converter Converter
}

func (f *numberField) Assign(val interface{}) error {
	var converted, err = f.converter(val)
	if err != nil {
		return err
	}
	f.value = converted
	return nil
}

func (f *numberField) Validate(val interface{}) []string {
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

func NumberField(name string, required bool, c Converter, vs ...Validator) Field {
	return &numberField{
		field{
			name:     name,
			required: required,
			ftype:    "Number",
			vs:       vs,
		},
		c,
	}
}
