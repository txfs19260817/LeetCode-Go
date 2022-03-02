package _014_Construct_String_With_Repeat_Limit

import "testing"

func Test_repeatLimitedString(t *testing.T) {
	type args struct {
		s           string
		repeatLimit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `s = "cczazcc", repeatLimit = 3`,
			args: args{"cczazcc", 3},
			want: "zzcccac",
		},
		{
			name: `s = "aababab", repeatLimit = 2`,
			args: args{"aababab", 2},
			want: "bbabaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatLimitedString(tt.args.s, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
