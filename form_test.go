package forms

import (
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
	var res = form.Validate(req)
	if res.Ok() {
		t.Fatal("Expected error")
	}
	var exp = "{\"dateofbirth\":[\"Field is required\"],\"language\":[\"Field is required\"]}"
	if res.String() != exp {
		t.Errorf("Unexpected result: %s", res.String())
	}
}
