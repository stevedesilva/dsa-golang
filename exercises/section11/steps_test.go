package section11

import (
	"reflect"
	"testing"
)

func Test_executeStep(t *testing.T) {
	type args struct {
		step int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "0",
			args: args{step: 0},
			want: []string{},
		},
		{
			name: "1",
			args: args{step: 1},
			want: []string{"#\n"},
		},
		{
			name: "2",
			args: args{step: 2},
			want: []string{"# \n", "##\n"},
		},
		{
			name: "3",
			args: args{step: 3},
			want: []string{"#  \n", "## \n", "###\n"},
		},
		{
			name: "4",
			args: args{step: 4},
			want: []string{"#   \n", "##  \n", "### \n", "####\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeStep(tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
