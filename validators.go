package forms

import (
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Validator func(lc *i18n.Localizer, val interface{})

func Min(v int) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		if len(val.(string)) < v {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "Length should be greater than {{.Count}}",
				},
				TemplateData: v,
			})
		}
	}
}

func Max(v int) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		if len(val.(string)) > v {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "Length should be less than {{.count}}",
				},
				TemplateData: v,
			})
		}
	}
}

func Within(l, h int) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		if s := val.(string); len(s) < l || len(s) > h {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "Length should be between {{.low}} and {{.high}}",
				},
				TemplateData: map[string]interface{}{
					"low":  l,
					"high": h,
				},
			})
		}
	}
}

func NumMin(v float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < v
		case time.Duration:
			err = float64(value) < v
		case int64:
			err = float64(value) < v
		default:
			lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
				ID: "expected numeric value",
			}})
		}
		if err {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value should be greater than {{.high}}",
				},
				TemplateData: v,
			})
		}
	}
}

func NumMax(v float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value > v
		case time.Duration:
			err = float64(value) > v
		case int64:
			err = float64(value) > v
		default:
			lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
				ID: "expected numeric value",
			}})
		}
		if err {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value should be less than {{.low}}",
				},
				TemplateData: v,
			})
		}
	}
}

func NumWithin(l, h float64) Validator {
	return func(lc *i18n.Localizer, val interface{}) {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < l || value > h
		case time.Duration:
			err = float64(value) < l || float64(value) > h
		case int64:
			err = float64(value) < l || float64(value) > h
		default:
			lc.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
				ID: "expected numeric value",
			}})
		}
		if err {
			lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID: "value should be between {{.low}} and {{.high}}",
				},
				TemplateData: map[string]interface{}{
					"low":  l,
					"high": h,
				},
			})
		}
	}
}

type FormValidator func(messenger Messenger, fields map[string]Field)
