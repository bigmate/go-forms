package forms

import (
	"bytes"
	"strconv"
)

const errorField = "error"

type errors map[string][]string

func (e errors) String() string {
	return string(e.serialize())
}

func (e errors) MarshalJSON() ([]byte, error) {
	return e.serialize(), nil
}

func (e errors) serialize() []byte {
	var counter int
	var buff bytes.Buffer
	buff.WriteString("{")
	for field, messages := range e {
		buff.WriteString("\"")
		buff.WriteString(field)
		buff.WriteString("\":[")
		for i := 0; i < len(messages); i++ {
			buff.WriteString(strconv.Quote(messages[i]))
			if i < len(messages)-1 {
				buff.WriteString(",")
			}
		}
		buff.WriteString("]")
		if counter < len(e)-1 {
			buff.WriteString(",")
		}
		counter++
	}
	buff.WriteString("}")
	return buff.Bytes()
}

func (e errors) has(field string) bool {
	var _, ok = e[field]
	return ok
}

func (e errors) add(field, message string) {
	if e.has(field) {
		e[field] = append(e[field], message)
	} else {
		e[field] = []string{message}
	}
}

func (e errors) addBulk(field string, messages []string) {
	if len(messages) == 0 {
		return
	}
	if e.has(field) {
		e[field] = append(e[field], messages...)
	} else {
		e[field] = messages
	}
}

func (e errors) empty() bool {
	return len(e) == 0
}

type validationError struct {
	msg string
}

func newError(msg string) error {
	return &validationError{msg: msg}
}

func (v *validationError) Error() string {
	return v.msg
}
