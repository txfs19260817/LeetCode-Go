package _189_Maximum_Number_of_Balloons

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `"nlaebolko"`,
			args: args{"nlaebolko"},
			want: 1,
		},
		{
			name: `"loonbalxballpoon"`,
			args: args{"loonbalxballpoon"},
			want: 2,
		},
		{
			name: `"leetcode"`,
			args: args{"leetcode"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
			if got := maxNumberOfBalloons2(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons2() = %v, want %v", got, tt.want)
			}
		})
	}
}
