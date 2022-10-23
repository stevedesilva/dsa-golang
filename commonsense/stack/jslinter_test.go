package stack_test

import (
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/commonsense/stack"
)

func Test_braceLinter_Validate(t *testing.T) {

	tests := []struct {
		name    string
		data    string
		want    bool
		wantErr error
	}{
		{
			name:    "empty string",
			data:    "",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "simple stack ignored",
			data:    "abscefg",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "(var x = {y: [1,2,3]})",
			data:    "(var x = {y: [1,2,3]})",
			want:    true,
			wantErr: nil,
		},
		{
			name:    "(var x = 2;",
			data:    "(var x = 2;",
			want:    false,
			wantErr: ErrMissingClosingBrace,
		},
		{
			name:    "var x = 2;)",
			data:    "var x = 2;)",
			want:    false,
			wantErr: ErrMissingOpeningBrace,
		},
		{
			name:    "(var x = [1,2,3)];",
			data:    "(var x = [1,2,3)];",
			want:    false,
			wantErr: ErrBraceMismatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewLinter(tt.data)
			got, err := b.Validate()
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
