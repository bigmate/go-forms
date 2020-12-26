package forms

import (
	"regexp"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type charField struct {
	field
}

func (f *charField) Assign(val interface{}) error {
	if _, ok := val.(string); !ok {
		return typeMismatchError
	}
	f.value = val
	f.bound = true
	return nil
}

func (f *charField) Validate(lc *i18n.Localizer, val interface{}) []string {
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

func CharField(name string, required bool, vs ...Validator) Field {
	return &charField{
		field{
			name:     name,
			required: required,
			ftype:    "String",
			vs:       vs,
		},
	}
}

var emailRE = regexp.MustCompile("^\\S{4,20}@\\S{2,15}\\.\\S{2,6}$")

func EmailField(name string, required bool, vs ...Validator) Field {
	var validator = func(lc *i18n.Localizer, val interface{}) {
		if !emailRE.Match([]byte(val.(string))) {
			lc.LocalizeMessage(&i18n.Message{
				ID:    "invalid email",
				Other: "invalid email",
			})
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
