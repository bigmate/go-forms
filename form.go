package forms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json"
	mimeApplicationForm = "application/x-www-form-urlencoded"
)

type Form interface {
	Validate(r *http.Request) Result
	Bind(s interface{}) error
}

func New(fields ...Field) Form {
	var form = &form{make(map[string]Field), make(errors)}
	for _, f := range fields {
		form.fields[f.Name()] = f
	}
	return form
}

type form struct {
	fields map[string]Field
	errs   errors
}

func (f *form) Bind(s interface{}) error {
	var typ = reflect.TypeOf(s)
	var val = reflect.ValueOf(s).Elem()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("pass pointer to struct")
	}
	return f.bind(val, typ)
}

func (f *form) bind(val reflect.Value, typ reflect.Type) error {
	for i := 0; i < val.NumField(); i++ {
		var fieldType = typ.Elem().Field(i)
		var field, ok = f.fields[fieldType.Tag.Get("form")]
		if !ok {
			continue
		}
		var strField = val.Field(i)
		var fv = reflect.ValueOf(field.Value())
		if !fv.IsValid() {
			continue
		}
		if !(strField.CanSet() && fv.Type().AssignableTo(strField.Type())) {
			return fmt.Errorf("imposible to assign to field %s", fieldType.Name)
		}
		strField.Set(fv)
	}
	return nil
}

func (f *form) Validate(r *http.Request) Result {
	var err = r.ParseForm()
	if err != nil {
		f.errs.add(errorField, t(invalidForm))
		return f.errs
	}
	if r.Method == http.MethodGet {
		f.validateForm(r.Form)
		return f.errs
	}
	var content = r.Header.Get(headerContentType)
	switch {
	case strings.HasPrefix(content, mimeApplicationJSON):
		f.validateJSON(r.Body)
	case strings.HasPrefix(content, mimeApplicationForm):
		f.validateForm(r.PostForm)
	default:
		f.errs.add(errorField, t(unsupportedContent))
	}
	return f.errs
}

func (f *form) validateJSON(rc io.Reader) {
	var dest = make(map[string]interface{})
	var err = json.NewDecoder(rc).Decode(&dest)
	if err != nil {
		f.errs.add(errorField, t(invalidJSON))
		return
	}
	for _, field := range f.fields {
		f.errs.addBulk(field.Name(), field.Validate(dest[field.Name()]))
	}
}

func (f *form) validateForm(v url.Values) {
	for _, field := range f.fields {
		if _, ok := v[field.Name()]; ok {
			f.errs.addBulk(field.Name(), field.Validate(v.Get(field.Name())))
		} else {
			f.errs.addBulk(field.Name(), field.Validate(nil))
		}
	}
}
