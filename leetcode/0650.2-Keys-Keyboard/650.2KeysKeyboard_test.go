package leetcode

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 0,
		},
		{
			name: "2",
			args: args{2},
			want: 2,
		},
		{
			name: "3",
			args: args{3},
			want: 3,
		},
		{
			name: "1000",
			args: args{1000},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
			if got := minSteps2(tt.args.n); got != tt.want {
				t.Errorf("minSteps2() = %v, want %v", got, tt.want)
			}
		})
	}
}
