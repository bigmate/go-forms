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

type orderedErrs struct {
	kv map[string][]string
	ll *linkedList
}

func newErrs() orderedErrs {
	return orderedErrs{
		kv: make(map[string][]string),
		ll: newLinkedList(),
	}
}

func (e orderedErrs) Ok() bool {
	return len(e.kv) == 0
}

func (e orderedErrs) String() string {
	var buff bytes.Buffer
	var node = e.ll.head
	for node != nil {
		var messages = e.kv[node.key]
		for i := 0; i < len(messages); i++ {
			buff.WriteString(messages[i])
			buff.WriteByte('\n')
		}
		node = node.next
	}
	if buff.Len() > 0 {
		buff.Truncate(buff.Len() - 1)
	}
	return buff.String()
}

func (e orderedErrs) MarshalJSON() ([]byte, error) {
	return e.Serialize(), nil
}

func (e orderedErrs) Serialize() []byte {
	var buff bytes.Buffer
	buff.WriteByte('{')
	var node = e.ll.head
	for node != nil {
		var messages = e.kv[node.key]
		buff.WriteByte('"')
		buff.WriteString(node.key)
		buff.WriteByte('"')
		buff.WriteString(":[")
		for i := 0; i < len(messages); i++ {
			buff.WriteString(strconv.Quote(messages[i]))
			if i < len(messages)-1 {
				buff.WriteByte(',')
			}
		}
		buff.WriteByte(']')
		if node.next != nil {
			buff.WriteByte(',')
		}
		node = node.next
	}
	buff.WriteByte('}')
	return buff.Bytes()
}

func (e orderedErrs) Has(field string) bool {
	return e.has(field)
}

func (e orderedErrs) has(field string) bool {
	var _, ok = e.kv[field]
	return ok
}

func (e orderedErrs) Add(field, message string) {
	e.add(field, message)
}

func (e orderedErrs) add(field, message string) {
	if e.has(field) {
		e.kv[field] = append(e.kv[field], message)
	} else {
		e.kv[field] = []string{message}
		e.ll.append(field)
	}
}

func (e orderedErrs) addBulk(field string, messages []string) {
	if len(messages) == 0 {
		return
	}
	if e.has(field) {
		e.kv[field] = append(e.kv[field], messages...)
	} else {
		e.kv[field] = messages
		e.ll.append(field)
	}
}
