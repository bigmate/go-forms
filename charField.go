package forms

import "regexp"

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
		return errors
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

var emailRE = regexp.MustCompile("^\\S{4,20}@\\S{2,15}\\.\\S{2,6}$")

func EmailField(name string, required bool, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if !emailRE.Match([]byte(val.(string))) {
			return T("invalid email")
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

var uuidRE = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func UUIDField(name string, required bool, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if !uuidRE.Match([]byte(val.(string))) {
			return T("invalid uuid")
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
