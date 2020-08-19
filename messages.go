package forms

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	invalidForm        = "Invalid form"
	unsupportedContent = "Unsupported content"
	invalidJSON        = "Invalid json"
	fieldRequired      = "Field is required"
	typeMismatch       = "Expected value type: %s"
)

func t(msg string, args ...interface{}) string {
	var pt = message.NewPrinter(language.English)
	return pt.Sprintf(msg, args...)
}

func T(msg string, args ...interface{}) error {
	return errors.New(t(msg, args...))
}

func init() {
	message.SetString(language.English, invalidForm, invalidForm)
	message.SetString(language.English, unsupportedContent, unsupportedContent)
	message.SetString(language.English, invalidJSON, invalidJSON)
	message.SetString(language.English, fieldRequired, fieldRequired)
	message.SetString(language.English, typeMismatch, typeMismatch)

	message.SetString(language.Kirghiz, invalidForm, "Форма туура эмес")
	message.SetString(language.Kirghiz, unsupportedContent, "Колдоого алынбаган контент")
	message.SetString(language.Kirghiz, invalidJSON, "Жараксыз JSON")
	message.SetString(language.Kirghiz, fieldRequired, "Талаа толтурулушу керек")
	message.SetString(language.Kirghiz, typeMismatch, "Күтүлүүчү маани түрү: %s")

	message.SetString(language.Russian, invalidForm, "Неверная форма")
	message.SetString(language.Russian, unsupportedContent, "Неподдерживаемый контент")
	message.SetString(language.Russian, invalidJSON, "Неверный JSON")
	message.SetString(language.Russian, fieldRequired, "Обязательно к заполнению")
	message.SetString(language.Russian, typeMismatch, "Ожидаемый тип значения: %s")

	message.SetString(language.Azerbaijani, invalidForm, "Yanlış forma")
	message.SetString(language.Azerbaijani, unsupportedContent, "Dəstəklənməyən məzmun")
	message.SetString(language.Azerbaijani, invalidJSON, "Yanlış JSON")
	message.SetString(language.Azerbaijani, fieldRequired, "Doldurmaq üçün tələb olunur")
	message.SetString(language.Azerbaijani, typeMismatch, "Gözlənilən dəyər növü:% s")
}
