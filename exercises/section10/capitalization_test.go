package section10

import "testing"

func TestCapitalizeSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single letter A",
			args: args{
				sentence: "A",
			},
			want: "A",
		},
		{
			name: "a",
			args: args{
				sentence: "a",
			},
			want: "A",
		},
		{
			name: "ab",
			args: args{
				sentence: "ab",
			},
			want: "Ab",
		},
		{
			name: "abc",
			args: args{
				sentence: "abc",
			},
			want: "Abc",
		},
		{
			name: "abcd",
			args: args{
				sentence: "abcd",
			},
			want: "Abcd",
		},
		{
			name: "abcd efg",
			args: args{
				sentence: "abcd efg",
			},
			want: "Abcd Efg",
		},
		{
			name: "a short sentence",
			args: args{
				sentence: "a short sentence",
			},
			want: "A Short Sentence",
		},
		{
			name: "look, it is not working!",
			args: args{
				sentence: "look, it is not working!",
			},
			want: "Look, It Is Not Working!",
		},
		{
			name: "look, it's not working!",
			args: args{
				sentence: "look, it's not working!",
			},
			want: "Look, It's Not Working!",
		},
		{
			name: "loOk, it's not working!",
			args: args{
				sentence: "loOk, it's not working!",
			},
			want: "LoOk, It's Not Working!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CapitalizeSentence(tt.args.sentence); got != tt.want {
				t.Errorf("CapitalizeSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
