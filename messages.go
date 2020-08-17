package forms

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

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

func init() {
	var cat = catalog.NewBuilder()

	cat.Set(language.English, invalidForm, catalog.String(invalidForm))
	cat.Set(language.English, unsupportedContent, catalog.String(unsupportedContent))
	cat.Set(language.English, invalidJSON, catalog.String(invalidJSON))
	cat.Set(language.English, fieldRequired, catalog.String(fieldRequired))
	cat.Set(language.English, typeMismatch, catalog.String(typeMismatch))

	cat.Set(language.Kirghiz, invalidForm, catalog.String("Форма туура эмес"))
	cat.Set(language.Kirghiz, unsupportedContent, catalog.String("Колдоого алынбаган контент"))
	cat.Set(language.Kirghiz, invalidJSON, catalog.String("Жараксыз JSON"))
	cat.Set(language.Kirghiz, fieldRequired, catalog.String("Талаа толтурулушу керек"))
	cat.Set(language.Kirghiz, typeMismatch, catalog.String("Күтүлүүчү маани түрү: %s"))

	cat.Set(language.Russian, invalidForm, catalog.String("Неверная форма"))
	cat.Set(language.Russian, unsupportedContent, catalog.String("Неподдерживаемый контент"))
	cat.Set(language.Russian, invalidJSON, catalog.String("Неверный JSON"))
	cat.Set(language.Russian, fieldRequired, catalog.String("Поле, обязательное для заполнения"))
	cat.Set(language.Russian, typeMismatch, catalog.String("Ожидаемый тип значения: %s"))
}
