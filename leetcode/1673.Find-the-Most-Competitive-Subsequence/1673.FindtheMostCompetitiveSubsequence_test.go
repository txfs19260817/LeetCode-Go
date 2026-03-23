package leetcode

import (
	"reflect"
	"testing"
)

func Test_mostCompetitive(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 5, 2, 6},
				k:    2,
			},
			want: []int{2, 6},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 4, 3, 3, 5, 4, 9, 6},
				k:    4,
			},
			want: []int{2, 3, 3, 4},
		},
		{
			name: "keeps smaller prefix opportunity",
			args: args{
				nums: []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2},
				k:    3,
			},
			want: []int{8, 80, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCompetitive(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
