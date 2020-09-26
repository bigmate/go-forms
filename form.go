package forms

import (
	"bytes"
	"encoding/json"
	"errors"
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
	BindValidators(validator ...FormValidator) Form

	// Available after validation
	MarshalJSON() ([]byte, error)
	Bind(s interface{}) error
	Fields() []Field
}

func New(fields ...Field) Form {
	var form = &form{
		fields:     make(map[string]Field),
		validators: make([]FormValidator, 0),
		errs:       make(errs),
	}
	for _, f := range fields {
		form.fields[f.Name()] = f
	}
	return form
}

type form struct {
	fields       map[string]Field
	validators   []FormValidator
	errs         errs
	cachedFields []Field
}

func (f *form) Bind(s interface{}) error {
	var typ = reflect.TypeOf(s)
	var val = reflect.ValueOf(s).Elem()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("pass pointer to struct")
	}
	return f.bind(val, typ)
}

func (f *form) bind(str reflect.Value, strType reflect.Type) error {
	for i := 0; i < str.NumField(); i++ {
		var fType = strType.Elem().Field(i)
		var field, ok = f.fields[fType.Tag.Get("form")]
		if !ok {
			continue
		}
		var strField = str.Field(i)
		var fv = reflect.ValueOf(field.Value())
		if !fv.IsValid() {
			continue
		}
		if !(strField.CanSet() && fv.Type().AssignableTo(strField.Type())) {
			return errors.New("impossible to assign to field " + fType.Name)
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
	f.runFormValidators()
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

func (f *form) MarshalJSON() ([]byte, error) {
	var buff bytes.Buffer
	var i int
	buff.WriteByte('{')
	for _, field := range f.fields {
		buff.WriteByte('"')
		buff.WriteString(field.Name())
		buff.WriteByte('"')
		buff.WriteByte(':')
		var bs, err = json.Marshal(field.Value())
		if err != nil {
			return nil, err
		}
		buff.Write(bs)
		if i < len(f.fields)-1 {
			buff.WriteByte(',')
		}
		i++
	}
	buff.WriteByte('}')
	return buff.Bytes(), nil
}

func (f *form) BindValidators(validators ...FormValidator) Form {
	f.validators = append(f.validators, validators...)
	return f
}

func (f *form) runFormValidators() {
	for _, formValidator := range f.validators {
		formValidator(f.errs, f.fields)
	}
}

func (f *form) Fields() []Field {
	if f.cachedFields != nil {
		return f.cachedFields
	}
	var fields = make([]Field, 0)
	for _, field := range f.fields {
		if field.Value() != nil {
			fields = append(fields, field)
		}
	}
	f.cachedFields = fields
	return fields
}
