package forms

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Min(v int) CharValidator {
	return func(lc *i18n.Localizer, val string) error {
		if len(val) < v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "length_greater",
					Other: "Length should be greater than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func Max(v int) CharValidator {
	return func(lc *i18n.Localizer, val string) error {
		if len(val) > v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "length_less",
					Other: "Length should be less than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func Within(l, h int) CharValidator {
	return func(lc *i18n.Localizer, val string) error {
		if l == h && len(val) != l {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "length_exact",
					Other: "Length should be exactly {{.}}",
				},
				TemplateData: l,
			}))
		}
		if len(val) < l || len(val) > h {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "length_between",
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

func NumMin(v float64) FloatValidator {
	return func(lc *i18n.Localizer, val float64) error {
		if val < v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "value_greater",
					Other: "value should be greater than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func NumMax(v float64) FloatValidator {
	return func(lc *i18n.Localizer, val float64) error {
		if val > v {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "value_less",
					Other: "value should be less than {{.}}",
				},
				TemplateData: v,
			}))
		}
		return nil
	}
}

func NumWithin(l, h float64) FloatValidator {
	return func(lc *i18n.Localizer, val float64) error {
		if val < l || val > h {
			return errors.New(lc.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "value_between",
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
