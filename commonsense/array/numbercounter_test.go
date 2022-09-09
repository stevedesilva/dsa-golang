package array

import (
	"reflect"
	"testing"
)

func TestCountOnes(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name    string
		args    args
		want    *CountResult
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountOnes(tt.args.array)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountOnes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountOnes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
