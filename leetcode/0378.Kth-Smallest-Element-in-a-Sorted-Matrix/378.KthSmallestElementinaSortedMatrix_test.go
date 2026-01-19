package leetcode

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example-1",
			args: args{
				matrix: [][]int{
					{1, 5, 9},
					{10, 11, 13},
					{12, 13, 15},
				},
				k: 8,
			},
			want: 13,
		},
		{
			name: "example-2",
			args: args{
				matrix: [][]int{
					{-5},
				},
				k: 1,
			},
			want: -5,
		},
		{
			name: "duplicates",
			args: args{
				matrix: [][]int{
					{1, 2, 2},
					{2, 3, 3},
					{3, 3, 4},
				},
				k: 5,
			},
			want: 3,
		},
		{
			name: "non-square-2x3",
			args: args{
				matrix: [][]int{
					{1, 3, 5},
					{6, 7, 12},
				},
				k: 4,
			},
			want: 6,
		},
		{
			name: "full-range-last",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
				k: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
