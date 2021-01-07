package forms

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

func mockForm() Form {
	return New(
		CharField("name", true, Within(5, 20)),
		FloatField("pets", false),
		DateTimeField("dateofbirth", true),
		ChoiceField("language", true, []interface{}{"KG", "EN", "RU", "TR"}),
		BoolField("married", false),
	)
}

func Test_form(t *testing.T) {
	var tests = []struct {
		name        string
		contentType string
		body        string
		boundFields []string
		language    string
		result      string
	}{
		{
			"A",
			mimeApplicationJSON,
			`{"name":"Cyborg"}`,
			[]string{"name"},
			"en",
			`{"dateofbirth":["Field is required"],"language":["Field is required"]}`,
		},
		{
			"B",
			mimeApplicationJSON,
			`{"name":"Cyborg", "language": 32}`,
			[]string{"name", "language"},
			"en",
			`{"dateofbirth":["Field is required"],"language":["Value should be one of [KG EN RU TR]"]}`,
		},
		{
			"C",
			mimeApplicationJSON,
			`{"pets": 1, "dateofbirth": "1965-04-29T00:00:00.000Z", "language": "TR"}`,
			[]string{"pets", "dateofbirth", "language"},
			"ru",
			`{"name":["Обязательно к заполнению"]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var form = mockForm()
			var req, err = http.NewRequest(
				http.MethodPost, "/",
				strings.NewReader(tt.body),
			)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set(headerContentType, tt.contentType)
			req.Header.Set("Accept-Language", tt.language)
			var res = form.Validate(req)
			var exp = []byte(tt.result)
			if bytes.Compare(res.Serialize(), exp) != 0 {
				t.Errorf("Unexpected result: %s", res.Serialize())
			}
			if bytes.Compare(res.Serialize(), res.Serialize()) != 0 {
				t.Errorf("Double serialization returned mismatching results")
			}
			var boundFields = form.Fields()
			if len(boundFields) != len(tt.boundFields) {
				t.Fatalf(
					"Bound fields length mismatch; actual: %v, expected: %v",
					len(boundFields), len(tt.boundFields),
				)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mockForm()
	}
}

func BenchmarkForm_Validate(b *testing.B) {
	var tests = []struct {
		name        string
		contentType string
		body        string
		language    string
	}{
		{
			"A",
			mimeApplicationJSON,
			`{"name":"Cyborg"}`,
			"en",
		},
		{
			"B",
			mimeApplicationJSON,
			`{"name":"Cyborg", "language": 32}`,
			"en",
		},
		{
			"C",
			mimeApplicationJSON,
			`{"pets": 1, "dateofbirth": "1965-04-29T00:00:00.000Z", "language": "TR"}`,
			"ru",
		},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			var form = mockForm()
			var req, err = http.NewRequest(
				http.MethodPost, "/",
				strings.NewReader(tt.body),
			)
			if err != nil {
				panic(err)
			}
			req.Header.Set(headerContentType, tt.contentType)
			req.Header.Set("Accept-Language", tt.language)
			form.Validate(req)
		}
	}
}