package _594_Longest_Harmonious_Subsequence

import "testing"

func Test_findLHS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,4,1,3,1,-14,1,-13]",
			args: args{[]int{1, 4, 1, 3, 1, -14, 1, -13}},
			want: 2,
		},
		{
			name: "[1,2,3,4]",
			args: args{[]int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "[1,1,1,1]",
			args: args{[]int{1, 1, 1, 1}},
			want: 0,
		},
		{
			name: "[1,3,2,2,5,2,3,7]",
			args: args{[]int{1, 3, 2, 2, 5, 2, 3, 7}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
			if got := findLHS2(tt.args.nums); got != tt.want {
				t.Errorf("findLHS2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findLHS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findLHS([]int{1, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13})
	}
}

func Benchmark_findLHS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findLHS2([]int{1, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13, 4, 1, 3, 1, -14, 1, -13})
	}
}
