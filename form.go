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

	"github.com/nicksnyder/go-i18n/v2/i18n"
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

type orderedFields struct {
	kv map[string]Field
	ll *linkedList
}

func (of orderedFields) add(f Field) {
	of.kv[f.Name()] = f
	of.ll.append(f.Name())
}

func newOrderedFields() orderedFields {
	return orderedFields{
		kv: make(map[string]Field),
		ll: newLinkedList(),
	}
}

func New(fields ...Field) Form {
	var form = &form{
		fields:     newOrderedFields(),
		validators: make([]FormValidator, 0),
		errs:       newErrs(),
	}
	for _, f := range fields {
		form.fields.add(f)
	}
	return form
}

type form struct {
	fields       orderedFields
	validators   []FormValidator
	errs         orderedErrs
	cachedFields []Field // bound Fields
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
		var field, ok = f.fields.kv[fType.Tag.Get("form")]
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
	var lc = i18n.NewLocalizer(bundle, r.Header.Get("Accept-Language"), "en")
	if err != nil {
		f.errs.add(errorField, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: invalidForm,
		}))
		return f.errs
	}
	if r.Method == http.MethodGet {
		f.validateForm(lc, r.Form)
		return f.errs
	}
	var content = r.Header.Get(headerContentType)
	switch {
	case strings.HasPrefix(content, mimeApplicationJSON):
		f.validateJSON(lc, r.Body)
	case strings.HasPrefix(content, mimeApplicationForm):
		f.validateForm(lc, r.PostForm)
	default:
		f.errs.add(errorField, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: unsupportedContent,
		}))
	}
	f.runFormValidators()
	return f.errs
}

func (f *form) validateJSON(lc *i18n.Localizer, rc io.Reader) {
	var dest = make(map[string]interface{})
	var err = json.NewDecoder(rc).Decode(&dest)
	if err != nil {
		f.errs.add(errorField, lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: invalidJSON,
		}))
		return
	}
	var node = f.fields.ll.head
	var fields = f.fields.kv
	for node != nil {
		var key = node.key
		f.errs.addBulk(key, fields[key].Validate(lc, dest[key]))
		node = node.next
	}
}

func (f *form) validateForm(lc *i18n.Localizer, v url.Values) {
	var node = f.fields.ll.head
	var fields = f.fields.kv
	for node != nil {
		var key = node.key
		if _, ok := v[key]; ok {
			f.errs.addBulk(
				key,
				fields[key].Validate(lc, fields[key].Validate(lc, v.Get(key))),
			)
		} else {
			f.errs.addBulk(
				key,
				fields[key].Validate(lc, fields[key].Validate(lc, nil)),
			)
		}
		node = node.next
	}
}

func (f *form) MarshalJSON() ([]byte, error) {
	var buff bytes.Buffer
	var i int
	buff.WriteByte('{')
	for _, field := range f.Fields() {
		if i > 0 {
			buff.WriteByte(',')
		}
		i++
		buff.WriteByte('"')
		buff.WriteString(field.Name())
		buff.WriteByte('"')
		buff.WriteByte(':')
		var bs, err = json.Marshal(field.Value())
		if err != nil {
			return nil, err
		}
		buff.Write(bs)
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
		formValidator(f.errs, f.fields.kv)
	}
}

func (f *form) Fields() []Field {
	if f.cachedFields != nil {
		return f.cachedFields
	}
	var fields = make([]Field, 0)
	var node = f.fields.ll.head
	var fs = f.fields.kv
	for node != nil {
		var field = fs[node.key]
		if field.bound() {
			fields = append(fields, field)
		}
		node = node.next
	}
	f.cachedFields = fields
	return fields
}
