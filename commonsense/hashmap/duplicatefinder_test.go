package hashmap_test

import (
	"reflect"
	"testing"

	. "github.com/stevedesilva/dsa-golang.git/commonsense/hashmap"
)

func Test_data_FirstDuplicate(t *testing.T) {
	type fields struct {
		array []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *string
		wantErr error
	}{
		{
			name:    "minimum not met",
			fields:  fields{array: []string{}},
			want:    nil,
			wantErr: ErrMinimumInputRequired,
		},
		{
			name:    "minimum not met with single value",
			fields:  fields{array: []string{"a"}},
			want:    nil,
			wantErr: ErrMinimumInputRequired,
		},
		{
			name:    "no duplicate found error",
			fields:  fields{array: []string{"a", "b", "c"}},
			want:    nil,
			wantErr: ErrNoDuplicates,
		},
		{
			name:    "minimum not met with single value",
			fields:  fields{array: []string{"a", "a"}},
			want:    stringPtr("a"),
			wantErr: nil,
		},
		{
			name:    "find duplicate not met with single value",
			fields:  fields{array: []string{"a", "b", "c", "a", "b", "d"}},
			want:    stringPtr("a"),
			wantErr: nil,
		},
		{
			name:    "find duplicate not met with single value",
			fields:  fields{array: []string{"ab", "bb", "cb", "ab", "bb", "db"}},
			want:    stringPtr("ab"),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewData(tt.fields.array)
			got, err := d.FirstDuplicate()
			if err != tt.wantErr {
				t.Errorf("FirstDuplicate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstDuplicate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringPtr(in string) *string {
	return &in
}
