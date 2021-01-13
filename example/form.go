package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/bigmate/go-forms"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func UserCreateForm() forms.Form {
	return forms.New(
		forms.CharField("first_name", true, forms.Min(3), forms.Max(15)),
		forms.CharField("username", true, forms.Min(8), forms.Max(8)),
		forms.CharField("phone_number", true, PhoneValidator),
		forms.NumberField("age", true, AgeValidator(18)),
		forms.BoolField("married", false),
	)
}

var phoneReg = regexp.MustCompile(`^\+7\d{10}$`)

func PhoneValidator(lc *i18n.Localizer, val string) error {
	if !phoneReg.Match([]byte(val)) {
		return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "phone_message",
			DefaultMessage: &i18n.Message{
				ID:    "phone_message",
				Other: "Provide valid phone number",
			},
		}))
	}
	return nil
}

func AgeValidator(maxAge int64) forms.NumValidator {
	return func(lc *i18n.Localizer, val int64) error {
		if val < maxAge {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "age_message",
				DefaultMessage: &i18n.Message{
					ID:    "age_message",
					Other: "You are under " + strconv.Itoa(int(maxAge)),
				},
			}))
		}
		return nil
	}
}

const data = `{
    "first_name": "Jo",
    "username": "jon_doe",
    "phone_number": "+7999394781f",
	"age": 16,
	"married": false
}`

func getRequest() *http.Request {
	req, err := http.NewRequest("POST", "/", strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func main() {
	form := UserCreateForm()
	res := form.Validate(getRequest())
	if !res.Ok() {
		bs, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Validation error:\n%s", bs)
		return
	}
	fmt.Println("Provided data has passed validation")
}
