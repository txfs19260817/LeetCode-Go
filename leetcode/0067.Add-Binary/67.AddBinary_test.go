package _067_Add_Binary

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `a = "11", b = "1"`,
			args: args{"11", "1"},
			want: "100",
		},
		{
			name: `a = "1010", b = "1011"`,
			args: args{"1010", "1011"},
			want: "10101",
		},
		{
			name: `a = "1111", b = "1111"`,
			args: args{"1111", "1111"},
			want: "11110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
