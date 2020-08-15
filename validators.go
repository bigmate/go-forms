package forms

import (
	"time"
)

type Validator func(val interface{}) error

func TimeValidator(val interface{}) error {
	var msg = T("please provide valid format")
	if t, ok := val.(string); ok {
		var _, err = time.Parse(time.RFC3339, t)
		if err != nil {
			return msg
		}
		return nil
	}
	return msg
}
