package forms

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json"
	mimeApplicationForm = "application/x-www-form-urlencoded"
)

type Form interface {
	Validate(r *http.Request) Message
}

func New(fields ...Field) Form {
	return &form{fields}
}

type form struct {
	fields []Field
}

func (f *form) Validate(r *http.Request) Message {
	var errs = make(errors)
	var err = r.ParseForm()
	if err != nil {
		errs.add(errorField, t(invalidForm))
		return errs
	}
	if r.Method == http.MethodGet {
		f.validateForm(r.Form, errs)
		return errs
	}
	var content = r.Header.Get(headerContentType)
	switch {
	case strings.HasPrefix(content, mimeApplicationJSON):
		f.validateJSON(r.Body, errs)
	case strings.HasPrefix(content, mimeApplicationForm):
		f.validateForm(r.PostForm, errs)
	default:
		errs.add(errorField, t(unsupportedContent))
	}
	return errs
}

func (f *form) validateJSON(rc io.Reader, errs errors) {
	var dest = make(map[string]interface{})
	var err = json.NewDecoder(rc).Decode(&dest)
	if err != nil {
		errs.add(errorField, t(invalidJSON))
	}
	for _, field := range f.fields {
		errs.addBulk(field.Name(), field.Validate(dest[field.Name()]))
	}
}

func (f *form) validateForm(v url.Values, errs errors) {
	for _, field := range f.fields {
		if _, ok := v[field.Name()]; ok {
			errs.addBulk(field.Name(), field.Validate(field.Convert(v.Get(field.Name()))))
		} else {
			errs.addBulk(field.Name(), field.Validate(nil))
		}
	}
}
