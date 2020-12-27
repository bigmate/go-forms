package forms

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

func Test_form_IsValid(t *testing.T) {
	var form = New(
		CharField("name", true, Within(5, 20)),
		FloatField("pets", false),
		DateTimeField("dateofbirth", true),
		ChoiceField("language", true, []interface{}{"KG", "EN", "RU", "TR"}),
		BoolField("married", false),
	)
	var body = strings.NewReader(`{"name":"Cyborg"}`)
	var req, err = http.NewRequest(http.MethodPost, "/", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(headerContentType, mimeApplicationJSON)
	req.Header.Set("Accept-Lang", "en")
	var res = form.Validate(req)
	if res.Ok() {
		t.Fatal("Expected error")
	}
	var exp = []byte(`{"dateofbirth":["Field is required"],"language":["Field is required"]}`)
	if bytes.Compare(res.Serialize(), exp) != 0 {
		t.Errorf("Unexpected result: %s", res.Serialize())
	}
}
