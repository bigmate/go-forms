package forms

import (
	"testing"
)

func Test_numberField_Validate(t *testing.T) {
	var f = NumberField("age", true, NumWithin(10, 30))
	if errs := f.Validate("10"); len(errs) > 0 {
		t.Error(errs)
		t.Logf("%T", f.Value())
	}
}
