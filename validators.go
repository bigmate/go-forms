package forms

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
