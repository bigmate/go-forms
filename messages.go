package forms

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	invalidForm        = "invalid_form"
	unsupportedContent = "unsupported_content"
	invalidJSON        = "invalid_json"
	fieldRequired      = "field_required"
	typeMismatch       = "type_mismatch"
)

var bundle = i18n.NewBundle(language.English)

func init() {
	bundle.MustAddMessages(language.English, &i18n.Message{ID: invalidForm, Other: "Invalid form"})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: unsupportedContent, Other: "Unsupported content"})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: invalidJSON, Other: "Invalid JSON"})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: fieldRequired, Other: "Field is required"})
	bundle.MustAddMessages(language.English, &i18n.Message{ID: typeMismatch, Other: "Expected value type: {{.}}"})

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(filepath.Dir(path), "translations")
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		match, err := filepath.Match("active.*.toml", f.Name())
		if err != nil {
			panic(err)
		}
		if match {
			bundle.MustLoadMessageFile(filepath.Join(path, f.Name()))
		}
	}
}
