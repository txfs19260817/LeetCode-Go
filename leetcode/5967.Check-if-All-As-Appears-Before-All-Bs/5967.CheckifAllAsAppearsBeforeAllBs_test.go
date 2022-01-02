package _967_Check_if_All_As_Appears_Before_All_Bs

import "testing"

func Test_checkString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abab",
			args: args{"abab"},
			want: false,
		},
		{
			name: "aabb",
			args: args{"aabb"},
			want: true,
		},
		{
			name: "bb",
			args: args{"bb"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkString(tt.args.s); got != tt.want {
				t.Errorf("checkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
