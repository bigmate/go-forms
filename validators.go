package forms

import (
	"errors"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Validator func(lc *i18n.Localizer, val interface{}) error

func Min(v int) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		if len(val.(string)) < v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "length_greater",
					Other: "Length should be greater than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func Max(v int) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		if len(val.(string)) > v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "length_less",
					Other: "Length should be less than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func Within(l, h int) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		if s := val.(string); len(s) < l || len(s) > h {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "length_between",
					Other: "Length should be between {{.low}} and {{.high}}",
				},
				TemplateData: map[string]interface{}{
					"low":  l,
					"high": h,
				},
			}))
		}
		return nil
	}
}

func NumMin(v float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < v
		case time.Duration:
			err = float64(value) < v
		case int64:
			err = float64(value) < v
		default:
			return errors.New(
				lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
					ID: "numeric_value_expected",
				}}))
		}
		if err {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value_greater",
					Other: "value should be greater than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func NumMax(v float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value > v
		case time.Duration:
			err = float64(value) > v
		case int64:
			err = float64(value) > v
		default:
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
				ID: "numeric_value_expected",
				Other: "expected numeric value",
			}}))
		}
		if err {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value_less",
					Other: "value should be less than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func NumWithin(l, h float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < l || value > h
		case time.Duration:
			err = float64(value) < l || float64(value) > h
		case int64:
			err = float64(value) < l || float64(value) > h
		default:
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
				ID: "numeric_value_expected",
				Other: "expected numeric value",
			}}))
		}
		if err {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value_between",
					Other: "value should be between {{.low}} and {{.high}}",
				},
				TemplateData: map[string]interface{}{
					"low":  l,
					"high": h,
				},
			}))
		}
		return nil
	}
}

type FormValidator func(messenger Messenger, fields map[string]Field)
