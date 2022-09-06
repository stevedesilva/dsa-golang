package array

import (
	"reflect"
	"testing"
)

func Test_markInventory(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := markInventory(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("markInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("markInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
