package leetcode

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[0,1,2,3,4,5,6,7,8,9]",
			args: args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: false,
		},
		{
			name: "[0]",
			args: args{[]int{0}},
			want: false,
		},
		{
			name: "[0,1,2,3,3,5,6,7,8,9]",
			args: args{[]int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}},
			want: false,
		},
		{
			name: "[0,3,2,1]",
			args: args{[]int{0, 3, 2, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
			if got := validMountainArray2(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_validMountainArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, i})
	}
}

func Benchmark_validMountainArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validMountainArray2([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, i})
	}
}
