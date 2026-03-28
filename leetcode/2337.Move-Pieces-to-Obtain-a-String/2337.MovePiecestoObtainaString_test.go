package leetcode

import "testing"

func Test_canChange(t *testing.T) {
	type args struct {
		start  string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				start:  "_L__R__R_",
				target: "L______RR",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				start:  "R_L_",
				target: "__LR",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				start:  "_R",
				target: "R_",
			},
			want: false,
		},
		{
			name: "Already Equal",
			args: args{
				start:  "L_R_",
				target: "L_R_",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChange(tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("canChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
