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

type errs struct {
	kv map[string][]string
	ll *linkedList
}

func newErrs() errs {
	return errs{
		kv: make(map[string][]string),
		ll: newLinkedList(),
	}
}

func (e errs) Ok() bool {
	return e.empty()
}

func (e errs) String() string {
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

func (e errs) MarshalJSON() ([]byte, error) {
	return e.Serialize(), nil
}

func (e errs) Serialize() []byte {
	var buff bytes.Buffer
	buff.WriteString("{")
	var node = e.ll.head
	for node != nil {
		var messages = e.kv[node.key]
		buff.WriteString("\"")
		buff.WriteString(node.key)
		buff.WriteString("\":[")
		for i := 0; i < len(messages); i++ {
			buff.WriteString(strconv.Quote(messages[i]))
			if i < len(messages)-1 {
				buff.WriteString(",")
			}
		}
		buff.WriteString("]")
		if node.next != nil {
			buff.WriteString(",")
		}
		node = node.next
	}
	buff.WriteString("}")
	return buff.Bytes()
}

func (e errs) Has(field string) bool {
	return e.has(field)
}

func (e errs) has(field string) bool {
	var _, ok = e.kv[field]
	return ok
}

func (e errs) Add(field, message string) {
	e.add(field, message)
}

func (e errs) add(field, message string) {
	if e.has(field) {
		e.kv[field] = append(e.kv[field], message)
	} else {
		e.kv[field] = []string{message}
		e.ll.append(field)
	}
}

func (e errs) addBulk(field string, messages []string) {
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

func (e errs) empty() bool {
	return len(e.kv) == 0
}
