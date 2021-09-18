package section11_test

import (
	"reflect"
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/exercises/section11"
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
			want: []string{"#"},
		},
		{
			name: "2",
			args: args{step: 2},
			want: []string{"# ", "##"},
		},
		{
			name: "3",
			args: args{step: 3},
			want: []string{"#  ", "## ", "###"},
		},
		{
			name: "4",
			args: args{step: 4},
			want: []string{"#   ", "##  ", "### ", "####"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteStep(tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeStepNaive(t *testing.T) {
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
			want: []string{"#"},
		},
		{
			name: "2",
			args: args{step: 2},
			want: []string{"# ", "##"},
		},
		{
			name: "3",
			args: args{step: 3},
			want: []string{"#  ", "## ", "###"},
		},
		{
			name: "4",
			args: args{step: 4},
			want: []string{"#   ", "##  ", "### ", "####"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteStepNaive(tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeStep() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_executeStepRecursive(t *testing.T) {
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
			want: []string{"#"},
		},
		{
			name: "2",
			args: args{step: 2},
			want: []string{"# ", "##"},
		},
		{
			name: "3",
			args: args{step: 3},
			want: []string{"#  ", "## ", "###"},
		},
		{
			name: "4",
			args: args{step: 4},
			want: []string{"#   ", "##  ", "### ", "####"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteRecursive(tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeStep() = %v, want %v", got, tt.want)
			}
		})
	}

}
