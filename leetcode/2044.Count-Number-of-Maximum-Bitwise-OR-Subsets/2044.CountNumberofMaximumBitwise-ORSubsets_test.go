package _044_Count_Number_of_Maximum_Bitwise_OR_Subsets

import (
	"math/rand"
	"testing"
)

func Test_countMaxOrSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,2,1,5",
			args: args{[]int{3, 2, 1, 5}},
			want: 6,
		},
		{
			name: "2,2,2",
			args: args{[]int{2, 2, 2}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
			if got := countMaxOrSubsets2(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countMaxOrSubsets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countMaxOrSubsets(rand.Perm(16))
	}
}

func Benchmark_countMaxOrSubsets2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countMaxOrSubsets2(rand.Perm(16))
	}
}
