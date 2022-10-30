package stack

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    string
		wantErr error
	}{
		{
			name:    "empty",
			data:    "",
			want:    "",
			wantErr: ErrMinimumInputNotMet,
		},
		{
			name:    "abc:cba",
			data:    "abc",
			want:    "cba",
			wantErr: nil,
		},
		{
			name:    "the:eht",
			data:    "the",
			want:    "eht",
			wantErr: nil,
		},
		{
			name:    "the abcde:edcba eht",
			data:    "the abcde",
			want:    "edcba eht",
			wantErr: nil,
		},
		{
			name:    "abcde:edcba",
			data:    "abcde",
			want:    "edcba",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := reverse(tt.data); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
			if _, err := reverse(tt.data); err != tt.wantErr {
				t.Errorf("reverse() error = %v, want Err %v", err, tt.want)
			}
		})
	}
}
