package forms

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Field interface {
	Name() string
	Value() interface{}
	Validate(lc *i18n.Localizer, val interface{}) []string
	Bound() bool
	set(val interface{}) error
}

type field struct {
	name     string
	required bool
	ftype    string
	bound    bool
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Bound() bool {
	return f.bound
}
