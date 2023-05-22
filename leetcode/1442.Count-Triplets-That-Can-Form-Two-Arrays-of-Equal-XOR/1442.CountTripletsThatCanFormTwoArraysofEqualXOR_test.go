package leetcode

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [2,3,1,6,7]",
			args: args{[]int{2, 3, 1, 6, 7}},
			want: 4,
		},
		{
			name: "arr = [1,1,1,1,1]",
			args: args{[]int{1, 1, 1, 1, 1}},
			want: 10,
		},
		{
			name: "arr = [2,3]",
			args: args{[]int{2, 3}},
			want: 0,
		},
		{
			name: "arr = [1,3,5,7,9]",
			args: args{[]int{1, 3, 5, 7, 9}},
			want: 3,
		},
		{
			name: "arr = [7,11,12,9,5,2,7,17,22]",
			args: args{[]int{7, 11, 12, 9, 5, 2, 7, 17, 22}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
			if got := countTriplets1(tt.args.arr); got != tt.want {
				t.Errorf("countTriplets1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countTriplets([]int{7, 11, 12, 9, 5, 2, 7, 17, 22})
	}
}

func Benchmark_countTriplets1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countTriplets1([]int{7, 11, 12, 9, 5, 2, 7, 17, 22})
	}
}
