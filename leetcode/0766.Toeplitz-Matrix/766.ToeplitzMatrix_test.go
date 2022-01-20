package _766_Toeplitz_Matrix

import "testing"

func Test_isToeplitzMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[[41,45],[81,41],[73,81],[47,73],[0,47],[79,76]]",
			args: args{[][]int{{41, 45}, {81, 41}, {73, 81}, {47, 73}, {0, 47}, {79, 76}}},
			want: false,
		},
		{
			name: "[[83],[64],[57]]",
			args: args{[][]int{{83}, {64}, {57}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
