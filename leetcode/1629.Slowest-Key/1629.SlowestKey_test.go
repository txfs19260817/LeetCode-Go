package leetcode

import "testing"

func Test_slowestKey(t *testing.T) {
	type args struct {
		releaseTimes []int
		keysPressed  string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: `releaseTimes = [12,23,36,46,62], keysPressed = "spuda"`,
			args: args{[]int{12, 23, 36, 46, 62}, "spuda"},
			want: 'a',
		},
		{
			name: `releaseTimes = [23,34,43,59,62,80,83,92,97], keysPressed = "qgkzzihfc"`,
			args: args{[]int{23, 34, 43, 59, 62, 80, 83, 92, 97}, "qgkzzihfc"},
			want: 'q',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slowestKey(tt.args.releaseTimes, tt.args.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
