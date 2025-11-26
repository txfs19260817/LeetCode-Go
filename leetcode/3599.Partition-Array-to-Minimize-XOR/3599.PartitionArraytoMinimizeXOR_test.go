package leetcode

import "testing"

func Test_minXor(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{nums: []int{1, 2, 3}, k: 2},
			want: 1,
		},
		{
			name: "example2",
			args: args{nums: []int{2, 3, 3, 2}, k: 3},
			want: 2,
		},
		{
			name: "example3",
			args: args{nums: []int{1, 1, 2, 3, 1}, k: 2},
			want: 0,
		},
		{
			name: "single_element",
			args: args{nums: []int{5}, k: 1},
			want: 5,
		},
		{
			name: "k_equals_1",
			args: args{nums: []int{1, 2, 3, 4}, k: 1},
			want: 4,
		},
		{
			name: "k_equals_len",
			args: args{nums: []int{4, 7, 1}, k: 3},
			want: 7,
		},
		{
			name: "all_same_even_split",
			args: args{nums: []int{2, 2, 2, 2}, k: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minXor(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
