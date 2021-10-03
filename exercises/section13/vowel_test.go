package section13

import "testing"

func TestVowel_CalculateNumberOfVowels(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Zero",
			input: "",
			want:  0,
		},
		{
			name:  "One lowercase",
			input: "a",
			want:  1,
		},
		{
			name:  "One uppercase",
			input: "A",
			want:  1,
		},
		{
			name:  "Hi there!",
			input: "Hi there!",
			want:  3,
		},
		{
			name:  "Why do you ask?",
			input: "Why do you ask?",
			want:  4,
		},
		{
			name:  "Why?",
			input: "Why?",
			want:  0,
		},
		{
			name:  "aa aa aa e e v?",
			input: "aa aa aa e e v?",
			want:  8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vowel{
				Input: tt.input,
			}
			//if got := v.CalculateNumberOfVowelsIterativeUsingMap(); got != tt.want {
			//	t.Errorf("CalculateNumberOfVowelsIterativeUsingMap() = %v, want %v", got, tt.want)
			//}
			//if got := v.CalculateNumberOfVowelsIterativeUsingContains(); got != tt.want {
			//	t.Errorf("CalculateNumberOfVowelsIterativeUsingContains() = %v, want %v", got, tt.want)
			//}
			if got := v.CalculateNumberOfVowelsRegex(); got != tt.want {
				t.Errorf("CalculateNumberOfVowelsRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}
