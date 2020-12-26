package forms

type choiceField struct {
	field
}

func (f *choiceField) Validate(lc *i18n.Localizer,val interface{}) []string {
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
	if f.Assign(val) != nil {
		errors = append(errors, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    typeMismatch,
			TemplateData: f.ftype,
		}))
		return errors
	}
	return f.runValidators(lc, errors)
}

func (f *choiceField) Assign(val interface{}) error {
	switch val.(type) {
	case string, float64:
		f.value = val
		f.bound = true
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
