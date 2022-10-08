package array

import "testing"

func Test_isCreditCardValidAsNumber(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "49927398716,true",
			number: 49927398716,
			want:   true,
		},
		{
			name:   "49927398717,false",
			number: 49927398717,
			want:   false,
		},
		{
			name:   "1234567812345678,false",
			number: 1234567812345678,
			want:   false,
		},
		{
			name:   "1234567812345670,true",
			number: 1234567812345670,
			want:   true,
		},
		{
			name:   "11111666,false",
			number: 11111666,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := IsCreditCardValidAsNumber(tt.number); res != tt.want {
				t.Errorf("IsCreditCardValidAsNumber() got= %t \twant= %t ", res, tt.want)
			}
		})
	}
}

func Test_isCreditCardValidAsString(t *testing.T) {

	tests := []struct {
		name   string
		number string
		want   bool
	}{
		{
			name:   "49927398716,true",
			number: "49927398716",
			want:   true,
		},
		{
			name:   "49927398717,false",
			number: "49927398717",
			want:   false,
		},
		{
			name:   "1234567812345678,false",
			number: "1234567812345678",
			want:   false,
		},
		{
			name:   "1234567812345670,true",
			number: "1234567812345670",
			want:   true,
		},
		{
			name:   "11111666,false",
			number: "11111666",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := IsCreditCardValidAsString(tt.number); res != tt.want {
				t.Errorf("IsCreditCardValidAsString() got= %t \twant= %t ", res, tt.want)
			}
		})
	}
}
