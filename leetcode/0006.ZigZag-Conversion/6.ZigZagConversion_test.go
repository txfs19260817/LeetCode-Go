package _006_ZigZag_Conversion

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1`,
			args: args{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1},
			want: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name: `"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2`,
			args: args{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 2},
			want: "ACEGIKMOQSUWYBDFHJLNPRTVXZ",
		},
		{
			name: `"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 3`,
			args: args{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 3},
			want: "AEIMQUYBDFHJLNPRTVXZCGKOSW",
		},
		{
			name: `"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 4`,
			args: args{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 4},
			want: "AGMSYBFHLNRTXZCEIKOQUWDJPV",
		},
		{
			name: `"ABCDEFGHIJKLMNOPQRSTUVWX", 5`,
			args: args{"ABCDEFGHIJKLMNOPQRSTUVWX", 5},
			want: "AIQBHJPRXCGKOSWDFLNTVEMU",
		},
		{
			name: `"A", 3`,
			args: args{"A", 3},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
