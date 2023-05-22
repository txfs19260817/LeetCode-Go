package leetcode

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "))((",
			args: args{"))(("},
			want: "",
		},
		{
			name: "())()(((",
			args: args{"())()((("},
			want: "()()",
		},
		{
			name: ")))t((u)",
			args: args{")))t((u)"},
			want: "t(u)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
