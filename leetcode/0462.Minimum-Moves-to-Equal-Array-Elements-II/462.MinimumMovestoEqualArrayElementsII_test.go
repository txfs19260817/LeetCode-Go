package _462_Minimum_Moves_to_Equal_Array_Elements_II

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: 2,
		},
		{
			name: "[1,10,2,9]",
			args: args{[]int{1, 10, 2, 9}},
			want: 16,
		},
		{
			name: "[1,0,0,8,6]",
			args: args{[]int{1, 0, 0, 8, 6}},
			want: 14,
		},
		{
			name: "[1,2]",
			args: args{[]int{1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
