package _283_Move_Zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		ans  []int
	}{
		{
			name: "nums = [0,1,0,3,12]",
			args: args{[]int{0, 1, 0, 3, 12}},
			ans:  []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp := make([]int, len(tt.args.nums))
			copy(tmp, tt.args.nums)
			moveZeroes(tmp)
			assert.Equal(t, tt.ans, tmp)
			copy(tmp, tt.args.nums)
			moveZeroes1(tmp)
			assert.Equal(t, tt.ans, tmp)
			copy(tmp, tt.args.nums)
			moveZeroes2(tmp)
			assert.Equal(t, tt.ans, tmp)
		})
	}
}

func Benchmark_moveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes([]int{0, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12})
	}
}

func Benchmark_moveZeroes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes1([]int{0, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12})
	}
}

func Benchmark_moveZeroes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes1([]int{0, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12, 1, 0, 3, 12})
	}
}
