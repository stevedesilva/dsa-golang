package optimise

import "testing"

/*
Player turn ME
Number of coins = 1 Winner = THEM
Number of coins = 2 Winner = ME
Number of coins = 3 Winner = ME
Number of coins = 4 Winner = THEM
Number of coins = 5 Winner = ME
Number of coins = 6 Winner = ME
Number of coins = 7 Winner = THEM
Number of coins = 8 Winner = ME
Number of coins = 9 Winner = ME
Number of coins = 10 Winner = THEM
*/
func Test_flipCoin(t *testing.T) {
	type args struct {
		coinNumber    int
		currentPlayer string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Number of coins = 1 Winner = THEM",
			args{
				coinNumber:    1,
				currentPlayer: "ME",
			},
			"THEM",
		},
		{
			"Number of coins = 2 Winner = ME",
			args{
				coinNumber:    2,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 3 Winner = ME",
			args{
				coinNumber:    3,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 4 Winner = THEM",
			args{
				coinNumber:    4,
				currentPlayer: "ME",
			},
			"THEM",
		},
		{
			"Number of coins = 5 Winner = ME",
			args{
				coinNumber:    5,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 6 Winner = ME",
			args{
				coinNumber:    6,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 7 Winner = THEM",
			args{
				coinNumber:    7,
				currentPlayer: "ME",
			},
			"THEM",
		},
		{
			"Number of coins = 8 Winner = ME",
			args{
				coinNumber:    8,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 9 Winner = ME",
			args{
				coinNumber:    9,
				currentPlayer: "ME",
			},
			"ME",
		},
		{
			"Number of coins = 10 Winner = THEM",
			args{
				coinNumber:    10,
				currentPlayer: "ME",
			},
			"THEM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipCoin(tt.args.coinNumber, tt.args.currentPlayer); got != tt.want {
				t.Errorf("flipCoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
