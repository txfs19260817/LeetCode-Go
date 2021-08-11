package _926_Flip_String_to_Monotone_Increasing

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "00110"`,
			args: args{"00110"},
			want: 1,
		},
		{
			name: `s = "010110"`,
			args: args{"010110"},
			want: 2,
		},
		{
			name: `s = "00011000"`,
			args: args{"00011000"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
