package leetcode

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,3,4,2,2]",
			args: args{[]int{1, 3, 4, 2, 2}},
			want: 2,
		},
		{
			name: "nums = [3,1,3,4,2]",
			args: args{[]int{3, 1, 3, 4, 2}},
			want: 3,
		},
		{
			name: "nums = [1,1]",
			args: args{[]int{1, 1}},
			want: 1,
		},
		{
			name: "nums = [1,1,2]",
			args: args{[]int{1, 1, 2}},
			want: 1,
		},
		{
			name: "nums = [8,7,1,10,17,15,18,11,16,9,19,12,5,14,3,4,2,13,18,18]",
			args: args{[]int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,3,4,2,2]",
			args: args{[]int{1, 3, 4, 2, 2}},
			want: 2,
		},
		{
			name: "nums = [3,1,3,4,2]",
			args: args{[]int{3, 1, 3, 4, 2}},
			want: 3,
		},
		{
			name: "nums = [1,1]",
			args: args{[]int{1, 1}},
			want: 1,
		},
		{
			name: "nums = [1,1,2]",
			args: args{[]int{1, 1, 2}},
			want: 1,
		},
		{
			name: "nums = [8,7,1,10,17,15,18,11,16,9,19,12,5,14,3,4,2,13,18,18]",
			args: args{[]int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate1(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicate([]int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18})
	}
}

func Benchmark_findDuplicate1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicate1([]int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18})
	}
}
