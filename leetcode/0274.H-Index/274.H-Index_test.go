package leetcode

import (
	"slices"
	"testing"
)

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}

	longZeros := make([]int, 10000)
	longLarge := make([]int, 10000)
	for i := range longLarge {
		longLarge[i] = 100000
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"single_zero", args{[]int{0}}, 0},
		{"single_one", args{[]int{1}}, 1},
		{"classic", args{[]int{3, 0, 6, 1, 5}}, 3},
		{"all_large", args{[]int{100, 200, 300}}, 3},
		{"sorted_asc", args{[]int{0, 1, 3, 5, 6}}, 3},
		{"sorted_desc", args{[]int{6, 5, 3, 1, 0}}, 3},
		{"duplicates", args{[]int{4, 4, 4, 4}}, 4},
		{"zeros", args{[]int{0, 0, 0}}, 0},
		{"mixed", args{[]int{1, 2, 3, 4, 5}}, 3},
		{"two", args{[]int{0, 1}}, 1},
		{"long_zeros", args{longZeros}, 0},
		{"long_large", args{longLarge}, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hIndex2(t *testing.T) {
	type args struct {
		citations []int
	}

	longZeros := make([]int, 10000)
	longLarge := make([]int, 10000)
	for i := range longLarge {
		longLarge[i] = 100000
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"single_zero", args{[]int{0}}, 0},
		{"single_one", args{[]int{1}}, 1},
		{"classic", args{[]int{3, 0, 6, 1, 5}}, 3},
		{"all_large", args{[]int{100, 200, 300}}, 3},
		{"sorted_asc", args{[]int{0, 1, 3, 5, 6}}, 3},
		{"sorted_desc", args{[]int{6, 5, 3, 1, 0}}, 3},
		{"duplicates", args{[]int{4, 4, 4, 4}}, 4},
		{"zeros", args{[]int{0, 0, 0}}, 0},
		{"mixed", args{[]int{1, 2, 3, 4, 5}}, 3},
		{"two", args{[]int{0, 1}}, 1},
		{"long_zeros", args{longZeros}, 0},
		{"long_large", args{longLarge}, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// hIndex2 sorts the slice in-place; pass a copy to avoid affecting other tests
			input := slices.Clone(tt.args.citations)
			if got := hIndex2(input); got != tt.want {
				t.Errorf("hIndex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_hIndex_longZeros(b *testing.B) {
	longZeros := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hIndex(longZeros)
	}
}

func Benchmark_hIndex_longLarge(b *testing.B) {
	longLarge := make([]int, 10000)
	for i := range longLarge {
		longLarge[i] = 100000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hIndex(longLarge)
	}
}

func Benchmark_hIndex_mixedLarge(b *testing.B) {
	mixed := make([]int, 10000)
	for i := range mixed {
		mixed[i] = i % 5000 // 一些较小一些较大的值
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hIndex(mixed)
	}
}

func Benchmark_hIndex2_longZeros(b *testing.B) {
	longZeros := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// hIndex2 会原地排序，拷贝以保持每次输入一致
		input := append([]int(nil), longZeros...)
		_ = hIndex2(input)
	}
}

func Benchmark_hIndex2_longLarge(b *testing.B) {
	longLarge := make([]int, 10000)
	for i := range longLarge {
		longLarge[i] = 100000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := append([]int(nil), longLarge...)
		_ = hIndex2(input)
	}
}

func Benchmark_hIndex2_mixedLarge(b *testing.B) {
	mixed := make([]int, 10000)
	for i := range mixed {
		mixed[i] = i % 5000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := append([]int(nil), mixed...)
		_ = hIndex2(input)
	}
}
