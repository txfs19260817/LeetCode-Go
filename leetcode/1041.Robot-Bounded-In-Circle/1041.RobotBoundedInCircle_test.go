package _041_Robot_Bounded_In_Circle

import "testing"

func Test_isRobotBounded(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `instructions = "GG"`,
			args: args{"GG"},
			want: false,
		},
		{
			name: `instructions = "GLRLGLLGLGRGLGL"`,
			args: args{"GLRLGLLGLGRGLGL"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRobotBounded(tt.args.instructions); got != tt.want {
				t.Errorf("isRobotBounded() = %v, want %v", got, tt.want)
			}
		})
	}
}
