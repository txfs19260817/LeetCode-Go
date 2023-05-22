package leetcode

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "B0B6G0R6R0R6G9",
			args: args{"B0B6G0R6R0R6G9"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
