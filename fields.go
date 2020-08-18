package forms

import (
	"strconv"
	"time"
)

type ftype int

const (
	char ftype = 1 << iota
	number
	boolean
	datetime
	array
	charOrNumber = char | number
)

var table = map[ftype]string{
	char:         "String",
	datetime:     "Datetime",
	number:       "Number",
	boolean:      "Boolean",
	array:        "Array",
	charOrNumber: "String or Number",
}

var layouts = []string{time.RFC3339, time.RFC822, time.RubyDate}

func (ft ftype) String() string {
	return table[ft]
}

type Field interface {
	Name() string
	Validate(val interface{}) []string
	Convert(val interface{}) interface{}
	Value() interface{}
}

type field struct {
	name     string
	required bool
	ftype    ftype
	value    interface{}
	vs       []Validator
}

func (f *field) toTime(val interface{}) interface{} {
	var s, ok = val.(string)
	if !ok {
		return val
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t
		}
	}
	return val
}

func (f *field) Value() interface{} {
	return f.value
}

func (f *field) Convert(val interface{}) interface{} {
	if s, ok := val.(string); ok {
		switch f.ftype {
		case number:
			var converted, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return val
			}
			return converted
		case boolean:
			var converted, err = strconv.ParseBool(s)
			if err != nil {
				return val
			}
			return converted
		case datetime:
			return f.toTime(val)
		}
	}
	return val
}

func (f *field) Name() string {
	return f.name
}

func (f *field) typeMatch(val interface{}) bool {
	switch val.(type) {
	case float64:
		return f.ftype == number || f.ftype == charOrNumber
	case string:
		return f.ftype == char || f.ftype == charOrNumber || f.ftype == datetime
	case bool:
		return f.ftype == boolean
	case []interface{}:
		return f.ftype == array
	default:
		return false
	}
}

func (f *field) Validate(val interface{}) []string {
	var errors = make([]string, 0)
	if !f.required && val == nil {
		return errors
	}
	if f.required && val == nil {
		errors = append(errors, t(fieldRequired))
		return errors
	}
	if !f.typeMatch(val) {
		errors = append(errors, t(typeMismatch, f.ftype))
		return errors
	}
	for _, validator := range f.vs {
		var err = validator(val)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) == 0 {
		f.value = val
	}
	return errors
}

func CharField(name string, required bool, minLen, maxLen int, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if s := val.(string); len(s) < minLen || len(s) > maxLen {
			return T("length should be between %v and %v", minLen, maxLen)
		}
		return nil
	}
	var f = &field{
		name:     name,
		required: required,
		ftype:    char,
		vs:       []Validator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}

func DateTimeField(name string, required bool, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		for _, layout := range layouts {
			if _, err := time.Parse(layout, val.(string)); err == nil {
				return nil
			}
		}
		return T("Invalid datetime format")
	}
	var f = &field{
		name:     name,
		required: required,
		ftype:    datetime,
		vs:       []Validator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}

func BoolField(name string, required bool, vs ...Validator) Field {
	return &field{
		name:     name,
		required: required,
		ftype:    boolean,
		vs:       vs,
	}
}

func NumberField(name string, required bool, min, max float64, vs ...Validator) Field {
	var validator = func(val interface{}) error {
		if v := val.(float64); v < min || v > max {
			return T("value should be between %v and %v", min, max)
		}
		return nil
	}
	var f = &field{
		name:     name,
		required: required,
		ftype:    number,
		vs:       []Validator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
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
	var f = &field{
		name:     name,
		required: required,
		ftype:    charOrNumber,
		vs:       []Validator{validator},
	}
	f.vs = append(f.vs, vs...)
	return f
}
