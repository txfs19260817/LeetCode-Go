package _958_Number_of_Smooth_Descent_Periods_of_a_Stock

import "testing"

func Test_getDescentPeriods(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "prices = [3,2,1,4]",
			args: args{[]int{3, 2, 1, 4}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDescentPeriods(tt.args.prices); got != tt.want {
				t.Errorf("getDescentPeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
