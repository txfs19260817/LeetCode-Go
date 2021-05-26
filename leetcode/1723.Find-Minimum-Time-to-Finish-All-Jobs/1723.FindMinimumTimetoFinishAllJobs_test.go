package _723_Find_Minimum_Time_to_Finish_All_Jobs

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jobs = [3,2,3], k = 3",
			args: args{[]int{3, 2, 3}, 3},
			want: 3,
		},
		{
			name: "jobs = [1,2,4,7,8], k = 2",
			args: args{[]int{1, 2, 4, 7, 8}, 2},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
