package array

import "testing"

func Test_findLargestProduct(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "less that 3 numbers in array",
			args: args{
				array: []int{},
			},
			want:    0,
			wantErr: ErrMinimumInputNotMet,
		},
		{
			name: "3 numbers only",
			args: args{
				array: []int{1, 2, 3},
			},
			want:    6,
			wantErr: nil,
		},
		{
			name: "4 numbers only",
			args: args{
				array: []int{1, 2, 3, 4},
			},
			want:    24,
			wantErr: nil,
		},
		{
			name: "5 numbers only",
			args: args{
				array: []int{1, 2, 3, 4, 5},
			},
			want:    60,
			wantErr: nil,
		},
		{
			name: "5 numbers - first biggest",
			args: args{
				array: []int{5, 5, 3, 4, 5},
			},
			want:    125,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLargestProduct(tt.args.array)
			if err != tt.wantErr {
				t.Errorf("findLargestProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findLargestProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
