package _957_Adding_Spaces_to_a_String

import "testing"

func Test_addSpaces(t *testing.T) {
	type args struct {
		s      string
		spaces []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `s = "LeetcodeHelpsMeLearn", spaces = [8,13,15]`,
			args: args{"LeetcodeHelpsMeLearn", []int{8, 13, 15}},
			want: "Leetcode Helps Me Learn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpaces(tt.args.s, tt.args.spaces); got != tt.want {
				t.Errorf("addSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
