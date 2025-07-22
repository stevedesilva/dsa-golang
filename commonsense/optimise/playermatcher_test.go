package optimise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPlayerMatch(t *testing.T) {
	type args struct {
		playersA []PlayerMatcher
		playersB []PlayerMatcher
	}
	tests := []struct {
		name string
		args args
		want []PlayerResult
	}{
		{
			name: "Test with matching players",
			args: args{
				playersA: []PlayerMatcher{
					{firstName: "John", lastName: "Doe", team: "Team A"},
					{firstName: "Jane", lastName: "Smith", team: "Team B"},
				},
				playersB: []PlayerMatcher{
					{firstName: "John", lastName: "Doe", team: "Team A"},
					{firstName: "Jane", lastName: "Smith", team: "Team B"},
				},
			},
			want: []PlayerResult{
				{firstName: "John", lastName: "Doe"},
				{firstName: "Jane", lastName: "Smith"},
			},
		},
		{
			name: "Test with no matching players",
			args: args{
				playersA: []PlayerMatcher{
					{"Mike", "Bradley", "Team C"},
					{"Saul", "Paul", "Team C"},
				},
				playersB: []PlayerMatcher{
					{"Steve", "P", "Team C"},
				},
			},
			want: []PlayerResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsPlayerMatch(tt.args.playersA, tt.args.playersB), "IsPlayerMatch(%v, %v)", tt.args.playersA, tt.args.playersB)
		})
	}
}
