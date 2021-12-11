package sn

import "testing"

func Test_substringCalculator(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcde",
			args: args{"abcde"},
			want: 15,
		},
		{
			name: "banana",
			args: args{"banana"},
			want: 15,
		},
		{
			name: "ababa",
			args: args{"ababa"},
			want: 9,
		},
		{
			name: "abbbcde",
			args: args{"abbbcde"},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substringCalculator(tt.args.s); got != tt.want {
				t.Errorf("substringCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
