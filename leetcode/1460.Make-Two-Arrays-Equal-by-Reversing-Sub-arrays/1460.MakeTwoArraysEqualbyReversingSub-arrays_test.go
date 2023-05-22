package leetcode

import "testing"

func Test_canBeEqual(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "target = [1,2,3,4], arr = [2,4,1,3]",
			args: args{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}},
			want: true,
		},
		{
			name: "target = [7], arr = [7]",
			args: args{[]int{7}, []int{7}},
			want: true,
		},
		{
			name: "target = [3,7,9], arr = [3,7,11]",
			args: args{[]int{3, 7, 9}, []int{3, 7, 11}},
			want: false,
		},
		{
			name: "target = [1,1,1,1,1], arr = [1,1,1,1,1]",
			args: args{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
			if got := canBeEqual1(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canBeEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canBeEqual([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1})
	}
}

func Benchmark_canBeEqual1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canBeEqual1([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1})
	}
}
