package section12

import (
	"reflect"
	"testing"
)

func TestPyramid_BuildPyramidIterative(t *testing.T) {
	type fields struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty",
			fields: fields{
				level: 0,
			},
			want: []string{},
		},
		{
			name: "1 level",
			fields: fields{
				level: 1,
			},
			want: []string{"#"},
		},
		{
			name: "2 levels",
			fields: fields{
				level: 2,
			},
			want: []string{" # ", "###"},
		},
		{
			name: "3 levels",
			fields: fields{
				level: 3,
			},
			want: []string{"  #  ", " ### ", "#####"},
		},
		{
			name: "4 levels",
			fields: fields{
				level: 4,
			},
			want: []string{"   #   ", "  ###  ", " ##### ", "#######"},
		},
		{
			name: "5 levels",
			fields: fields{
				level: 5,
			},
			want: []string{"    #    ", "   ###   ", "  #####  ", " ####### ", "#########"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pyramid{
				level: tt.fields.level,
			}
			if got := p.BuildPyramidIterative(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildPyramidIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPyramid_BuildPyramidRecursive(t *testing.T) {
	type fields struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty",
			fields: fields{
				level: 0,
			},
			want: []string{},
		},
		{
			name: "1 level",
			fields: fields{
				level: 1,
			},
			want: []string{"#"},
		},
		{
			name: "2 levels",
			fields: fields{
				level: 2,
			},
			want: []string{" # ", "###"},
		},
		{
			name: "3 levels",
			fields: fields{
				level: 3,
			},
			want: []string{"  #  ", " ### ", "#####"},
		},
		{
			name: "4 levels",
			fields: fields{
				level: 4,
			},
			want: []string{"   #   ", "  ###  ", " ##### ", "#######"},
		},
		{
			name: "5 levels",
			fields: fields{
				level: 5,
			},
			want: []string{"    #    ", "   ###   ", "  #####  ", " ####### ", "#########"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pyramid{
				level: tt.fields.level,
			}
			if got := p.BuildPyramidRecursive(tt.fields.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildPyramidRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
