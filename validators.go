package forms

import (
	"time"
)

type Validator func(val interface{}) error

func Min(v int) Validator {
	return func(val interface{}) error {
		if len(val.(string)) < v {
			return T("Length should be greater than %v", v)
		}
		return nil
	}
}

func Max(v int) Validator {
	return func(val interface{}) error {
		if len(val.(string)) > v {
			return T("Length should be less than %v", v)
		}
		return nil
	}
}

func Within(l, h int) Validator {
	return func(val interface{}) error {
		if s := val.(string); len(s) < l || len(s) > h {
			return T("Length should be between %v and %v", l, h)
		}
		return nil
	}
}

func NumMin(v float64) Validator {
	return func(val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < v
		case time.Duration:
			err = float64(value) < v
		case int64:
			err = float64(value) < v
		default:
			return T("expected numeric value")
		}
		if err {
			return T("value should be greater than %v", v)
		}
		return nil
	}
}

func NumMax(v float64) Validator {
	return func(val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value > v
		case time.Duration:
			err = float64(value) > v
		case int64:
			err = float64(value) > v
		default:
			return T("expected numeric value")
		}
		if err {
			return T("value should be less than %v", v)
		}
		return nil
	}
}

func NumWithin(l, h float64) Validator {
	return func(val interface{}) error {
		var err bool
		switch value := val.(type) {
		case float64:
			err = value < l || value > h
		case time.Duration:
			err = float64(value) < l || float64(value) > h
		case int64:
			err = float64(value) < l || float64(value) > h
		default:
			return T("expected numeric value")
		}
		if err {
			return T("value should be between %v and %v", l, h)
		}
		return nil
	}
}

type FormValidator func(messenger Messenger, fields map[string]Field)
