package forms

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Test_numberField_Validate(t *testing.T) {
	var lc = i18n.NewLocalizer(bundle, "en", "ru")
	var f = FloatField("age", true, NumWithin(10, 30))
	if errs := f.Validate(lc, "10"); len(errs) > 0 {
		t.Error(errs)
		t.Logf("%T", f.Value())
	}
}
