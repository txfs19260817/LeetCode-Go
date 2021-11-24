package _423_Reconstruct_Original_Digits_from_English

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "owoztneoer",
			args: args{"owoztneoer"},
			want: "012",
		},
		{
			name: "fviefuro",
			args: args{"fviefuro"},
			want: "45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
