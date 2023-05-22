package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "k = 3, n = 7",
			args: args{3, 7},
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "k = 3, n = 9",
			args: args{3, 9},
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "k = 4, n = 1",
			args: args{4, 1},
			want: nil,
		},
		{
			name: "k = 3, n = 2",
			args: args{3, 2},
			want: nil,
		},
		{
			name: "k = 9, n = 45",
			args: args{9, 45},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
