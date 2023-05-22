package leetcode

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "people = [1,2], limit = 3",
			args: args{[]int{1, 2}, 3},
			want: 1,
		},
		{
			name: "people = [3,2,2,1], limit = 3",
			args: args{[]int{3, 2, 2, 1}, 3},
			want: 3,
		},
		{
			name: "people = [3,5,3,4], limit = 5",
			args: args{[]int{3, 5, 3, 4}, 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
