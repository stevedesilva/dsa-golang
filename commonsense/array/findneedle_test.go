package array

import "testing"

func TestSearchHaystack(t *testing.T) {
	type args struct {
		needle   string
		haystack string
	}
	type response struct {
		want bool
		err  error
	}
	tests := []struct {
		name string
		args args
		res  response
	}{
		{
			name: "thequickbrownfoxjumpedoverthelazydog",
			args: args{
				haystack: "thequickbrownfoxjumpedoverthelazydog",
				needle:   "fox",
			},
			res: response{
				want: true,
				err:  nil,
			},
		},
		{
			name: "thequickbrownfoxjumpedoverthelazydog",
			args: args{
				haystack: "thequickbrownfoxjumpedoverthelazydog",
				needle:   "notfound",
			},
			res: response{
				want: false,
				err:  nil,
			},
		},
		{
			name: "empty",
			args: args{
				needle:   "",
				haystack: "",
			},
			res: response{
				want: false,
				err:  ErrNeedleAndHaystackRequired,
			},
		},
		{
			name: "single letter matched",
			args: args{
				needle:   "a",
				haystack: "a",
			},
			res: response{
				want: true,
				err:  nil,
			},
		},
		{
			name: "single letter not matched",
			args: args{
				needle:   "a",
				haystack: "b",
			},
			res: response{
				want: false,
				err:  nil,
			},
		},
		{
			name: "single double not matched",
			args: args{
				haystack: "aa",
				needle:   "b",
			},
			res: response{
				want: false,
				err:  nil,
			},
		},
		{
			name: "matched b",
			args: args{
				haystack: "ab",
				needle:   "b",
			},
			res: response{
				want: true,
				err:  nil,
			},
		},
		{
			name: "single match",
			args: args{
				haystack: "abcde",
				needle:   "c",
			},
			res: response{
				want: true,
				err:  nil,
			},
		},
		{
			name: "abcde cd",
			args: args{
				haystack: "abcde",
				needle:   "cd",
			},
			res: response{
				want: true,
				err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchHaystack(tt.args.needle, tt.args.haystack)
			if err != tt.res.err {
				t.Errorf("SearchHaystack() error = %v\t, want error %v\n", err, tt.res.err)
			}
			if got != tt.res.want {
				t.Errorf("SearchHaystack() = res %v\t, want res %v\n", got, tt.res.want)
			}

		})
	}
}
