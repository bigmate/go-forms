package forms

import (
	"reflect"
	"testing"
)

func Test_errors_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		e       errors
		want    []byte
		wantErr bool
	}{
		{
			name: "Valid",
			e: errors{
				"A": []string{"B", "C", "D"},
			},
			want:    []byte(`{"A":["B","C","D"]}`),
			wantErr: false,
		},
		{
			name: "Invalid",
			e: errors{
				"A": []string{"B", "C", "D"},
			},
			want:    []byte(`{"A":["B","C","D"]}`),
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
		e    errors
		want string
	}{
		{
			name: "First",
			e: errors{
				"A": []string{"B", "C", "D"},
			},
			want: `{"A":["B","C","D"]}`,
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
