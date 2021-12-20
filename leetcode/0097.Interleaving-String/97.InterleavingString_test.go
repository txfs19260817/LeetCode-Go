package _097_Interleaving_String

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"`,
			args: args{"aabcc", "dbbca", "aadbbcbcac"},
			want: true,
		},
		{
			name: `s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"`,
			args: args{"aabcc", "dbbca", "aadbbbaccc"},
			want: false,
		},
		{
			name: `s1 = "", s2 = "", s3 = ""`,
			args: args{"", "", ""},
			want: true,
		},
		{
			name: `s1 = "", s2 = "", s3 = "a"`,
			args: args{"", "", "a"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
			if got := isInterleave2(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave2() = %v, want %v", got, tt.want)
			}
		})
	}
}
