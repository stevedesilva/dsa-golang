package leapyear

import "testing"

func TestIsLeap(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2016",
			args: args{year: 2016},
			want: true,
		},
		{
			name: "2024",
			args: args{year: 2024},
			want: true,
		},
		{
			name: "2028",
			args: args{year: 2028},
			want: true,
		},
		{
			name: "2032",
			args: args{year: 2032},
			want: true,
		},
		{
			name: "2031",
			args: args{year: 2031},
			want: false,
		},
		{
			name: "2033",
			args: args{year: 2033},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeap(tt.args.year); got != tt.want {
				t.Errorf("IsLeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
