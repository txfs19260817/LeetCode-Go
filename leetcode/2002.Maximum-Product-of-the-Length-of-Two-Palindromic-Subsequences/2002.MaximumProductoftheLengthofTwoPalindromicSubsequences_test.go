package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "leetcodecom",
			args: args{"leetcodecom"},
			want: 9,
		},
		{
			name: "dd",
			args: args{"dd"},
			want: 1,
		},
		{
			name: "accbcaxxcxx",
			args: args{"accbcaxxcxx"},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.s); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
