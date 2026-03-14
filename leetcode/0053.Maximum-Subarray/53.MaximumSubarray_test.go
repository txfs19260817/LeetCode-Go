package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "leetcode example",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "all positive",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "all negative",
			args: args{nums: []int{-8, -3, -6, -2, -5, -4}},
			want: -2,
		},
		{
			name: "zeros and negatives",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "best subarray in middle",
			args: args{nums: []int{1, -2, 3, 10, -4, 7, 2, -5}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSubArray(tt.args.nums))
			assert.Equal(t, tt.want, maxSubArray2(tt.args.nums))
		})
	}
}

func benchmarkMaxSubArrayInput() []int {
	input := make([]int, 10000)
	for i := range input {
		switch i % 5 {
		case 0:
			input[i] = 7
		case 1:
			input[i] = -3
		case 2:
			input[i] = 5
		case 3:
			input[i] = -8
		default:
			input[i] = 6
		}
	}
	return input
}

func Benchmark_maxSubArray(b *testing.B) {
	input := benchmarkMaxSubArrayInput()
	b.ResetTimer()
	for b.Loop() {
		_ = maxSubArray(input)
	}
}

func Benchmark_maxSubArray2(b *testing.B) {
	input := benchmarkMaxSubArrayInput()
	b.ResetTimer()
	for b.Loop() {
		_ = maxSubArray2(input)
	}
}
