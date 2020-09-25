package forms

import (
	"strconv"
)

type Converter func(interface{}) (interface{}, error)

func ToInt(val interface{}) (interface{}, error) {
	var converted interface{}
	switch value := val.(type) {
	case string:
		var c, err = strconv.Atoi(value)
		if err != nil {
			return nil, typeMismatchError
		}
		converted = c
	case float64:
		converted = int(value)
	default:
		return nil, typeMismatchError
	}
	return converted, nil
}

func ToFloat(val interface{}) (interface{}, error) {
	var converted interface{}
	switch value := val.(type) {
	case string:
		var c, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, typeMismatchError
		}
		converted = c
	case float64:
		converted = value
	default:
		return nil, typeMismatchError
	}
	return converted, nil
}

func ToInt64(val interface{}) (interface{}, error) {
	var converted interface{}
	switch value := val.(type) {
	case string:
		var c, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, typeMismatchError
		}
		converted = c
	case float64:
		converted = int64(value)
	default:
		return nil, typeMismatchError
	}
	return converted, nil
}
