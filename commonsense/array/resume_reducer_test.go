package array

import (
	"reflect"
	"testing"
)

func TestPickResume(t *testing.T) {

	tests := []struct {
		name    string
		resumes []string
		want    *string
		wantErr error
	}{
		/*
		 "doc1:doc1",
		            "doc1,doc2,doc3:doc2",
		            "doc1,doc2,doc3,doc4:doc3",
		            "doc1,doc2,doc3,doc4,doc5:doc3",
		            "doc1,doc2,doc3,doc4,doc5,doc6,doc7:doc5",
		*/
		{
			name:    "nil resume input",
			resumes: nil,
			want:    nil,
			wantErr: ErrMinimumValue,
		},
		{
			name:    "empty resume input",
			resumes: []string{},
			want:    nil,
			wantErr: ErrMinimumValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PickResume(tt.resumes)
			if err != tt.wantErr {
				t.Errorf("PickResume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickResume() got = %v, want %v", got, tt.want)
			}
		})
	}
}
