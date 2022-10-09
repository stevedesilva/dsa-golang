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
		{
			name:    "doc1:doc1",
			resumes: []string{"doc1"},
			want:    stringPtr("doc1"),
			wantErr: ErrMinimumValue,
		},
		{
			name:    "doc1,doc2,doc3:doc2",
			resumes: []string{"doc1", "doc2", "doc3"},
			want:    stringPtr("doc2"),
			wantErr: ErrMinimumValue,
		},
		{
			name:    "doc1,doc2,doc3,doc4:doc3",
			resumes: []string{"doc1", "doc2", "doc3", "doc4"},
			want:    stringPtr("doc3"),
			wantErr: ErrMinimumValue,
		},
		{
			name:    "doc1,doc2,doc3,doc4,doc5:doc3",
			resumes: []string{"doc1", "doc2", "doc3", "doc4", "doc5"},
			want:    stringPtr("doc3"),
			wantErr: ErrMinimumValue,
		},
		{
			name:    "doc1,doc2,doc3,doc4,doc5,doc6,doc7:doc5",
			resumes: []string{"doc1", "doc2", "doc3", "doc4", "doc5", "doc6", "doc7"},
			want:    stringPtr("doc5"),
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
				if got == nil {
					t.Errorf("PickResume() got = %v, want %v", got, toString(tt.want))
				} else {
					t.Errorf("PickResume() got = %v, want %v", toString(got), toString(tt.want))
				}
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func toString(s *string) string {
	return *s
}
