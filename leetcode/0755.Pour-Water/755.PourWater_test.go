package leetcode

import (
	"reflect"
	"testing"
)

func Test_pourWater(t *testing.T) {
	type args struct {
		heights []int
		volume  int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				heights: []int{2, 1, 1, 2, 1, 2, 2},
				volume:  4,
				k:       3,
			},
			want: []int{2, 2, 2, 3, 2, 2, 2},
		},
		{
			name: "example 2",
			args: args{
				heights: []int{1, 2, 3, 4},
				volume:  2,
				k:       2,
			},
			want: []int{2, 3, 3, 4},
		},
		{
			name: "example 3",
			args: args{
				heights: []int{3, 1, 3},
				volume:  5,
				k:       1,
			},
			want: []int{4, 4, 4},
		},
		{
			name: "symmetric valley near center with small volume",
			args: args{
				heights: []int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1},
				volume:  2,
				k:       5,
			},
			want: []int{1, 2, 3, 4, 3, 3, 2, 2, 3, 4, 3, 2, 1},
		},
		{
			name: "large volume settles left then right",
			args: args{
				heights: []int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1},
				volume:  10,
				k:       2,
			},
			want: []int{4, 4, 4, 4, 3, 3, 3, 3, 3, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pourWater(tt.args.heights, tt.args.volume, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pourWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
