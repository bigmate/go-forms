package forms

import (
	"reflect"
	"testing"
)

func newErrors(fields []string, msgs [][]string) orderedErrs {
	if len(fields) != len(msgs) {
		panic("make sure fields and msgs lengths equal")
	}
	var e = newErrs()
	for i := 0; i < len(fields); i++ {
		e.addBulk(fields[i], msgs[i])
	}
	return e
}

func Test_errors_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		e       orderedErrs
		want    []byte
		wantErr bool
	}{
		{
			name:    "first",
			e:       newErrors([]string{"A"}, [][]string{{"B", "C", "D"}}),
			want:    []byte(`{"A":["B","C","D"]}`),
			wantErr: false,
		},
		{
			name: "second",
			e: newErrors(
				[]string{"a", "b", "c"},
				[][]string{
					{"B", "C", "D"},
					{"b", "c", "d"},
					{"E", "F", "G"},
				}),
			want:    []byte(`{"a":["B","C","D"],"b":["b","c","d"],"c":["E","F","G"]}`),
			wantErr: false,
		},
		{
			name: "third",
			e: newErrors([]string{}, [][]string{}),
			want:    []byte("{}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_errors_String(t *testing.T) {
	tests := []struct {
		name string
		e    orderedErrs
		want string
	}{
		{
			name: "First",
			e:    newErrors([]string{"A"}, [][]string{{"B", "C", "D"}}),
			want: "B\nC\nD",
		},
		{
			name: "Second",
			e:    newErrors([]string{}, [][]string{}),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
