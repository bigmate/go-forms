package forms

import (
	"errors"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		val  string
		want error
		lc   string
	}{
		{
			"first",
			2,
			"a",
			errors.New("Узундугу эн аз 2 болуусу керек"),
			"ky",
		},
		{
			"second",
			3,
			"coolest thing",
			nil,
			"ky",
		},
		{
			"third",
			5,
			"col",
			errors.New("Length should be greater than 5"),
			"en",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vl := Min(tt.arg)
			lc := i18n.NewLocalizer(bundle, tt.lc)
			err := vl(lc, tt.val)
			if err != nil && tt.want != nil && err.Error() != tt.want.Error() ||
				(err == nil && tt.want != nil) || (err != nil && tt.want == nil) {
				t.Errorf("Expected: %s; got: %s\n", tt.want, err)
			}
		})
	}
}
