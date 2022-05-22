package _070_Calculate_Digit_Sum_of_a_String

import "testing"

func Test_digitSum(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11111222223,3",
			args: args{"11111222223", 3},
			want: "135",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitSum(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("digitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
