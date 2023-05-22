package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `nums = ["01","10"]`,
			args: args{[]string{"01", "10"}},
			want: []string{"00", "11"},
		},
		{
			name: `nums = ["111","011","001"]`,
			args: args{[]string{"111", "011", "001"}},
			want: []string{"101", "110", "000", "010", "100"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.want, findDifferentBinaryString(tt.args.nums))
			assert.Contains(t, tt.want, findDifferentBinaryString2(tt.args.nums))
		})
	}
}

func Benchmark_findDifferentBinaryString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDifferentBinaryString([]string{"111", "011", "001"})
	}
}

func Benchmark_findDifferentBinaryString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDifferentBinaryString2([]string{"111", "011", "001"})
	}
}
