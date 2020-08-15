package forms

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

var userForm = New(
	CharField("name", true, 3, 10),
	NumberField("age", true, 18, 55),
	BoolField("married", true),
	ChoiceField("country", false, []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}),
)

func Test_form_IsValid(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 500; i++ {
		go func(name, query string) {
			var req, err = http.NewRequest(http.MethodGet, "/?"+query, nil)
			if err != nil {
				return
			}
			req.Header.Set(headerContentType, mimeApplicationForm)
			var ve = userForm.Validate(req)
			if !ve.Ok() {
				fmt.Println(name, ve)
			}
			wg.Done()
		}("FIRST", "name=bekmamat&age=25&married=false")
	}
	for i := 0; i < 500; i++ {
		i := i
		go func(name, query string) {
			var req, err = http.NewRequest(http.MethodGet, "/?"+query, nil)
			if err != nil {
				wg.Done()
				return
			}
			req.Header.Set(headerContentType, mimeApplicationForm)
			var ve = userForm.Validate(req)
			if !ve.Ok() {
				fmt.Println(i)
			}
			wg.Done()
		}("SECOND", "name=be&age=17&married=false")
	}
	wg.Wait()
}
