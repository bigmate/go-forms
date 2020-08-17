package forms

import "fmt"

var (
	invalidForm        = "Invalid form"
	unsupportedContent = "Unsupported content"
	invalidJSON        = "Invalid json"
	fieldRequired      = "Field is required"
	typeMismatch       = "Expected value type: %s"
)

func t(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}

func T(msg string, args ...interface{}) error {
	return newError(t(msg, args...))
}
