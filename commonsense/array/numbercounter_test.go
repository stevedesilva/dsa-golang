package array

import (
	"reflect"
	"testing"
)

func TestCountOnes(t *testing.T) {
	type args struct {
		array [][]int
	}
	tests := []struct {
		name    string
		args    args
		want    *CountResult
		wantErr error
	}{
		{
			name: "empty array 1",
			args: args{
				array: [][]int{},
			},
			wantErr: ErrNoInputProvided,
			want:    nil,
		},
		{
			name: "empty array 2",
			args: args{
				array: [][]int{{}, {}},
			},
			wantErr: ErrNoInputProvided,
			want:    nil,
		},
		{
			name: "first array",
			args: args{
				array: [][]int{{0, 1, 0, 0, 1}, {1, 1, 0}, {0, 0, 1}, {1}, {0}},
			},
			wantErr: nil,
			want: &CountResult{
				Value: 6,
			},
		},
		{
			name: "second array",
			args: args{
				array: [][]int{{0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {0, 0, 0, 1}, {1}},
			},
			wantErr: nil,
			want: &CountResult{
				Value: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountOnes(tt.args.array)
			if err != tt.wantErr {
				t.Errorf("CountOnes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountOnes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
