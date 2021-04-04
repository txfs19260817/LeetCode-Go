package _781_Rabbits_in_Forest

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "answers = [1, 1, 2]",
			args: args{[]int{1, 1, 2}},
			want: 5,
		},
		{
			name: "answers = [10, 10, 10]",
			args: args{[]int{10, 10, 10}},
			want: 11,
		},
		{
			name: "answers = []",
			args: args{[]int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
