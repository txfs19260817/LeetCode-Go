package leetcode

import (
	"reflect"
	"testing"
)

func Test_findCoins(t *testing.T) {
	type args struct {
		numWays []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{numWays: []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}},
			want: []int{2, 4, 6},
		},
		{
			name: "case 2",
			args: args{numWays: []int{1, 2, 2, 3, 4}},
			want: []int{1, 2, 5},
		},
		{
			name: "case 3",
			args: args{numWays: []int{1, 2, 3, 4, 15}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCoins(tt.args.numWays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
