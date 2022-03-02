package _010_Minimum_Time_to_Complete_Trips

import "testing"

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "5,10,10",
			args: args{[]int{5, 10, 10}, 9},
			want: 25,
		},
		{
			name: "1,2,3",
			args: args{[]int{1, 2, 3}, 5},
			want: 3,
		},
		{
			name: "1,5",
			args: args{[]int{1, 5}, 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.time, tt.args.totalTrips); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
