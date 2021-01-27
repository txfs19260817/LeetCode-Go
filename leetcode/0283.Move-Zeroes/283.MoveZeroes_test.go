package _283_Move_Zeroes

import "testing"

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
