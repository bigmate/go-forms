package forms

import (
	"bytes"
	"errors"
	"strconv"
)

const errorField = "error"

var typeMismatchError = errors.New("types mismatch")

type Result interface {
	Ok() bool
	String() string
	Serialize() []byte
	MarshalJSON() ([]byte, error)
}

type Messenger interface {
	Add(field, message string)
	Has(field string) bool
}

type errs map[string][]string

func (e errs) Ok() bool {
	return e.empty()
}

func (e errs) String() string {
	var buff bytes.Buffer
	for _, messages := range e {
		for i := 0; i < len(messages); i++ {
			buff.WriteString(messages[i])
			buff.WriteByte('\n')
		}
	}
	buff.Truncate(buff.Len() - 1)
	return buff.String()
}

func (e errs) MarshalJSON() ([]byte, error) {
	return e.Serialize(), nil
}

func (e errs) Serialize() []byte {
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

func (e errs) Has(field string) bool {
	return e.has(field)
}

func (e errs) has(field string) bool {
	var _, ok = e[field]
	return ok
}

func (e errs) Add(field, message string) {
	e.add(field, message)
}

func (e errs) add(field, message string) {
	if e.has(field) {
		e[field] = append(e[field], message)
	} else {
		e[field] = []string{message}
	}
}

func (e errs) addBulk(field string, messages []string) {
	if len(messages) == 0 {
		return
	}
	if e.has(field) {
		e[field] = append(e[field], messages...)
	} else {
		e[field] = messages
	}
}

func (e errs) empty() bool {
	return len(e) == 0
}
