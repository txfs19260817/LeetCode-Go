package _392_Longest_Happy_Prefix

import "testing"

func Test_longestPrefix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "l",
			args: args{"l"},
			want: "",
		},
		{
			name: "level",
			args: args{"level"},
			want: "l",
		},
		{
			name: "ababab",
			args: args{"ababab"},
			want: "abab",
		},
		//{
		//	name: "vwantmbocxcwrqtvgzuvgrmdltfiglltaxkjfajxthcppcatddcunpkqsgpnjjgqanrwabgrtwuqbrfl(hash-collision)",
		//	args: args{"vwantmbocxcwrqtvgzuvgrmdltfiglltaxkjfajxthcppcatddcunpkqsgpnjjgqanrwabgrtwuqbrfl"},
		//	want: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrefix(tt.args.s); got != tt.want {
				t.Errorf("longestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
