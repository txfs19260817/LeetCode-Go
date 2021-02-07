package _234_Replace_the_Substring_for_Balanced_String

import "testing"

func Test_balancedString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "QWER",
			args: args{"QWER"},
			want: 0,
		},
		{
			name: "QQWE",
			args: args{"QQWE"},
			want: 1,
		},
		{
			name: "QQQW",
			args: args{"QQQW"},
			want: 2,
		},
		{
			name: "QQQQ",
			args: args{"QQQQ"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedString(tt.args.s); got != tt.want {
				t.Errorf("balancedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
