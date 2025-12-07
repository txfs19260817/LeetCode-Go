package leetcode

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n=1,k=1",
			args: args{n: 1, k: 1},
			want: 1, // 只有一种颜色，只有一根柱子
		},
		{
			name: "n=1,k=3",
			args: args{n: 1, k: 3},
			want: 3, // 任意一种颜色
		},
		{
			name: "n=2,k=1",
			args: args{n: 2, k: 1},
			want: 1, // 只能全是同一个颜色，长度2没超限制
		},
		{
			name: "n=2,k=2",
			args: args{n: 2, k: 2},
			want: 4, // 2^2 种，都合法
		},
		{
			name: "n=3,k=1",
			args: args{n: 3, k: 1},
			want: 0, // 只能 111，但 3 连一色非法
		},
		{
			name: "n=3,k=2",
			args: args{n: 3, k: 2},
			want: 6,
			// 所有长度3的 0/1 串除了 000 和 111，剩 6 种
		},
		{
			name: "n=3,k=3",
			args: args{n: 3, k: 3},
			want: 24,
			// 总 3^3=27，减去 3 种全同色，剩 24
		},
		{
			name: "n=4,k=2",
			args: args{n: 4, k: 2},
			want: 10,
		},
		{
			name: "n=5,k=2",
			args: args{n: 5, k: 2},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
