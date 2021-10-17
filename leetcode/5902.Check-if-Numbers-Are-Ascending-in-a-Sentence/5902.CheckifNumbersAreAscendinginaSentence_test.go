package _902_Check_if_Numbers_Are_Ascending_in_a_Sentence

import "testing"

func Test_areNumbersAscending(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			args: args{"1 box has 3 blue 4 red 6 green and 12 yellow marbles"},
			want: true,
		},
		{
			name: "hello world 5 x 5",
			args: args{"hello world 5 x 5"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areNumbersAscending(tt.args.s); got != tt.want {
				t.Errorf("areNumbersAscending() = %v, want %v", got, tt.want)
			}
		})
	}
}
