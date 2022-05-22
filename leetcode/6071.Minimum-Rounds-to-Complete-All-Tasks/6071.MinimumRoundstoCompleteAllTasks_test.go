package _071_Minimum_Rounds_to_Complete_All_Tasks

import "testing"

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tasks = [2,2,3,3,2,4,4,4,4,4]",
			args: args{[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}},
			want: 4,
		},
		{
			name: "tasks = [2,3,3",
			args: args{[]int{2, 3, 3}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRounds(tt.args.tasks); got != tt.want {
				t.Errorf("minimumRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
