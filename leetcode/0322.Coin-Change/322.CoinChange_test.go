package _322_Coin_Change

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coins = [1,2,5], amount = 11",
			args: args{[]int{1, 2, 5}, 11},
			want: 3,
		},
		{
			name: "coins = [2], amount = 3",
			args: args{[]int{2}, 3},
			want: -1,
		},
		{
			name: "coins = [1], amount = 0",
			args: args{[]int{1}, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
