package _135_Candy

import "testing"

var benchCase = []int{1, 2, 3, 4, 5, 3, 2, 1, 2, 6, 5, 4, 3, 3, 2, 1, 1, 3, 3, 3, 4, 2}

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ratings = [1,0,2]",
			args: args{[]int{1, 0, 2}},
			want: 5,
		},
		{
			name: "ratings = [1,2,2]",
			args: args{[]int{1, 2, 2}},
			want: 4,
		},
		{
			name: "ratings = [1,3,2,2,1]",
			args: args{[]int{1, 3, 2, 2, 1}},
			want: 7,
		},
		{
			name: "ratings = [1,3,4,5,2]",
			args: args{[]int{1, 3, 4, 5, 2}},
			want: 11,
		},
		{
			name: "ratings = [1,2,3,4,5,3,2,1,2,6,5,4,3,3,2,1,1,3,3,3,4,2]",
			args: args{[]int{1, 2, 3, 4, 5, 3, 2, 1, 2, 6, 5, 4, 3, 3, 2, 1, 1, 3, 3, 3, 4, 2}},
			want: 47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := candy1(tt.args.ratings); got != tt.want {
				t.Errorf("candy1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_candy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = candy(benchCase)
	}
}

func Benchmark_candy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = candy1(benchCase)
	}
}
