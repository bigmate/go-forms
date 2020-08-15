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
	IsValid(r *http.Request) bool
	Messages() json.Marshaler
	Reset()
}

func New(fields ...Field) Form {
	return &form{
		fields: fields,
		errs:   make(errors),
	}
}

type form struct {
	fields []Field
	errs   errors
}

func (f *form) IsValid(r *http.Request) bool {
	var err = r.ParseForm()
	if err != nil {
		f.errs.add(errorField, invalidForm)
		return false
	}
	if r.Method == http.MethodGet {
		return f.validForm(r.Form)
	}
	var content = r.Header.Get(headerContentType)
	switch {
	case strings.HasPrefix(content, mimeApplicationJSON):
		return f.validJSON(r.Body)
	case strings.HasPrefix(content, mimeApplicationForm):
		return f.validForm(r.PostForm)
	default:
		f.errs.add(errorField, unsupportedContent)
		return false
	}
}

func (f *form) validJSON(rc io.Reader) bool {
	var dest = make(map[string]interface{})
	var err = json.NewDecoder(rc).Decode(&dest)
	if err != nil {
		f.errs.add(errorField, invalidJSON)
		return false
	}
	for _, field := range f.fields {
		f.errs.addBulk(field.Name(), field.Validate(dest[field.Name()]))
	}
	return f.errs.empty()
}

func (f *form) validForm(v url.Values) bool {
	for _, field := range f.fields {
		if _, ok := v[field.Name()]; ok {
			f.errs.addBulk(field.Name(), field.Validate(field.Convert(v.Get(field.Name()))))
		} else {
			f.errs.addBulk(field.Name(), field.Validate(nil))
		}
	}
	return f.errs.empty()
}

func (f *form) Messages() json.Marshaler {
	return f.errs
}

func (f *form) Reset() {
	f.errs = make(errors)
}
