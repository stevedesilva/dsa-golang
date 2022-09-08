package array

import (
	"reflect"
	"testing"
)

func Test_markInventory(t *testing.T) {

	tests := []struct {
		name    string
		args    []string
		want    []string
		wantErr error
	}{
		{
			name:    "Purple Shirt",
			args:    []string{"Purple Shirt"},
			want:    []string{"Purple Shirt Size: 1", "Purple Shirt Size: 2", "Purple Shirt Size: 3", "Purple Shirt Size: 4", "Purple Shirt Size: 5"},
			wantErr: nil,
		},
		//{
		//	name:    "Green Shirt",
		//	args:    []string{"Green Shirt"},
		//	want:    []string{"Green Shirt Size: 1,Green Shirt Size: 2,Green Shirt Size: 3,Green Shirt Size: 4,Green Shirt Size: 5"},
		//	wantErr: nil,
		//},
		//{
		//	name:    "Red Shirt",
		//	args:    []string{"Red Shirt"},
		//	want:    []string{"Red Shirt Size: 1,Red Shirt Size: 2,Red Shirt Size: 3,Red Shirt Size: 4,Red Shirt Size: 5"},
		//	wantErr: nil,
		//},
		//{
		//	name:    "Pink Shirt,Red Shirt-Pink",
		//	args:    []string{"Pink Shirt,Red Shirt-Pink"},
		//	want:    []string{"Pink Shirt,Red Shirt-Pink Shirt Size: 1,Pink Shirt Size: 2,Pink Shirt Size: 3,Pink Shirt Size: 4,Pink Shirt Size: 5,Red Shirt Size: 1,Red Shirt Size: 2,Red Shirt Size: 3,Red Shirt Size: 4,Red Shirt Size: 5"},
		//	wantErr: nil,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := markInventory(tt.args)
			if err != tt.wantErr {
				t.Errorf("markInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("markInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markInventoryReturnsErrorWhenNoInputGiven(t *testing.T) {
	result, err := markInventory([]string{})
	if result != nil {
		t.Errorf("No result should be returned")
	}
	if err == nil {
		t.Errorf("error should be returned")
	}

}
