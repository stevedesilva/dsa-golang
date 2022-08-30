package array

import (
	"reflect"
	"testing"
)

func Test_findStartMidEndOfArray(t *testing.T) {

	tests := []struct {
		name    string
		args    []int
		want    *Result
		wantErr error
	}{
		{
			name:    "Error when no value supplied",
			args:    []int{},
			want:    nil,
			wantErr: ErrNoInputSupplied,
		},
		{
			name:    "1",
			args:    []int{1, 1, 1},
			want:    &Result{1, 1, 1},
			wantErr: nil,
		},
		{
			name:    "1,2,3",
			args:    []int{1, 2, 3},
			want:    &Result{1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "1,2,3,4,5",
			args:    []int{1, 2, 3, 4, 5},
			want:    &Result{1, 3, 5},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findStartMidEndOfArray(tt.args)

			if err != tt.wantErr {
				t.Errorf("findStartMidEndOfArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStartMidEndOfArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}
