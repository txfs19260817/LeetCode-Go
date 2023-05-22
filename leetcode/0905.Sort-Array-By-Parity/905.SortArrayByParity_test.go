package leetcode

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [0]",
			args: args{[]int{0}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
			if got := sortArrayByParity2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_sortArrayByParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity(rand.Perm(100))
	}
}

func Benchmark_sortArrayByParity2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity2(rand.Perm(100))
	}
}
