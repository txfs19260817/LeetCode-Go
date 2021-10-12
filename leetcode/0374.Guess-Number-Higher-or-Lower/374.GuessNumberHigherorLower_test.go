package _374_Guess_Number_Higher_or_Lower

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n, p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 10, pick = 6",
			args: args{10, 6},
			want: 6,
		},
		{
			name: "n = 1, pick = 1",
			args: args{1, 1},
			want: 1,
		},
		{
			name: "n = 2, pick = 1",
			args: args{2, 1},
			want: 1,
		},
		{
			name: "n = 2, pick = 2",
			args: args{2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.args.p
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
