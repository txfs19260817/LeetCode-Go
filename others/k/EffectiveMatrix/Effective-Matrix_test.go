package EffectiveMatrix

import "testing"

func Test_effectiveMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}},
			want: false,
		},
		{
			name: "2",
			args: args{[][]int{{1, 3, 2}, {2, 1, 3}, {3, 2, 1}}},
			want: true,
		},
		{
			name: "3",
			args: args{[][]int{{1, 2}, {2, 1}}},
			want: true,
		},
		{
			name: "4",
			args: args{[][]int{{1, 2}, {1, 2}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := effectiveMatrix(tt.args.mat); got != tt.want {
				t.Errorf("effectiveMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
