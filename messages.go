package forms

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	invalidForm        = "Invalid form"
	unsupportedContent = "Unsupported content"
	invalidJSON        = "Invalid json"
	fieldRequired      = "Field is required"
	typeMismatch       = "Expected value type: {{.type}}"
)

var bundle = i18n.NewBundle(language.English)

func init() {
	bundle.MustAddMessages(language.English, &i18n.Message{ID: invalidForm, Other: invalidForm})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: unsupportedContent, Other: unsupportedContent})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: invalidJSON, Other: invalidJSON})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: fieldRequired, Other: fieldRequired})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: typeMismatch, Other: typeMismatch})

	bundle.MustAddMessages(language.Kirghiz, &i18n.Message{ID: invalidForm, Other: "Форма туура эмес"})
	bundle.MustAddMessages(language.Kirghiz, &i18n.Message{ID: unsupportedContent, Other: "Колдоого алынбаган контент"})
	bundle.MustAddMessages(language.Kirghiz, &i18n.Message{ID: invalidJSON, Other: "Жараксыз JSON"})
	bundle.MustAddMessages(language.Kirghiz, &i18n.Message{ID: fieldRequired, Other: "Талаа толтурулушу керек"})
	bundle.MustAddMessages(language.Kirghiz, &i18n.Message{ID: typeMismatch, Other: "Күтүлүүчү маани түрү: {{.type}}"})

	bundle.MustAddMessages(language.Russian, &i18n.Message{ID: invalidForm, Other: "Неверная форма"})
	bundle.MustAddMessages(language.Russian, &i18n.Message{ID: unsupportedContent, Other: "Неподдерживаемый контент"})
	bundle.MustAddMessages(language.Russian, &i18n.Message{ID: invalidJSON, Other: "Неверный JSON"})
	bundle.MustAddMessages(language.Russian, &i18n.Message{ID: fieldRequired, Other: "Обязательно к заполнению"})
	bundle.MustAddMessages(language.Russian, &i18n.Message{ID: typeMismatch, Other: "Ожидаемый тип значения: {{.type}}"})

	bundle.MustAddMessages(language.Azerbaijani, &i18n.Message{ID: invalidForm, Other: "Yanlış forma"})
	bundle.MustAddMessages(language.Azerbaijani, &i18n.Message{ID: unsupportedContent, Other: "Dəstəklənməyən məzmun"})
	bundle.MustAddMessages(language.Azerbaijani, &i18n.Message{ID: invalidJSON, Other: "Yanlış JSON"})
	bundle.MustAddMessages(language.Azerbaijani, &i18n.Message{ID: fieldRequired, Other: "Doldurmaq üçün tələb olunur"})
	bundle.MustAddMessages(language.Azerbaijani, &i18n.Message{ID: typeMismatch, Other: "Gözlənilən dəyər növü: {{.type}}"})
}
