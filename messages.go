package forms

import "fmt"

var (
	invalidForm        = "Invalid form"
	unsupportedContent = "Unsupported content"
	invalidJSON        = "Invalid json"
	fieldRequired      = "field is required"
	typeMismatch       = "expected value type: %s"
)

type Lang int

const (
	RU Lang = 1 << iota
	KG
	AZ
)

func t(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}

func T(msg string, args ...interface{}) error {
	return newError(t(msg, args...))
}
