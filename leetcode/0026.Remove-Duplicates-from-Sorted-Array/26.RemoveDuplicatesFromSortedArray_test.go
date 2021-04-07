package _026_Remove_Duplicates_from_Sorted_Array

import "testing"

var benchCase = []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5, 5, 6, 6, 6}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,1,2]",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "nums = [0,0,1,1,1,2,2,3,3,4]",
			args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(append([]int(nil), tt.args.nums...)); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if got := removeDuplicates1(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates(append([]int(nil), benchCase...))
	}
}

func Benchmark_removeDuplicates1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates1(append([]int(nil), benchCase...))
	}
}
