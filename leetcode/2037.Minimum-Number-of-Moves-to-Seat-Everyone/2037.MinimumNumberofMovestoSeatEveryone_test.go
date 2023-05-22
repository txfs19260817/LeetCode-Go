package leetcode

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "seats = [3,1,5], students = [2,7,4]",
			args: args{[]int{3, 1, 5}, []int{2, 7, 4}},
			want: 4,
		},
		{
			name: "seats = [2,2,6,6], students = [1,3,2,6]",
			args: args{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}},
			want: 4,
		},
		{
			name: "seats = [4,1,5,9], students = [1,3,2,6]",
			args: args{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("minMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
