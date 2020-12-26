package forms

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Field interface {
	Name() string
	Value() interface{}
	Validate(lc *i18n.Localizer, val interface{}) []string
	Assign(val interface{}) error
	Bound() bool
}

type field struct {
	name     string
	required bool
	ftype    string
	bound    bool
	value    interface{}
	vs       []Validator
}

func (f *field) Value() interface{} {
	return f.value
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Bound() bool {
	return f.bound
}

func (f *field) runValidators(lc *i18n.Localizer, errors []string) []string {
	for _, validator := range f.vs {
		validator(lc, f.value)
	}
	return errors
}
