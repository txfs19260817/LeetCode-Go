package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeForTest(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 2, 6, 0, 0, 0},
				m:     3,
				nums2: []int{3, 4, 5},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{16},
				n:     1,
			},
			want: []int{1, 16},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{16, 0},
				m:     1,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeForTest(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeForTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeForTest1(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 2, 6, 0, 0, 0},
				m:     3,
				nums2: []int{3, 4, 5},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{16},
				n:     1,
			},
			want: []int{1, 16},
		},
		{
			name: "merge sorted arrays",
			args: args{
				nums1: []int{16, 0},
				m:     1,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeForTest1(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeForTest1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}

func Benchmark_merge1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge1([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}
