package _413_Minimum_Value_to_Get_Positive_Step_by_Step_Sum

import "testing"

func Test_minStartValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [-3,2,-3,4,2]",
			args: args{[]int{-3, 2, -3, 4, 2}},
			want: 5,
		},
		{
			name: "nums = [1,-2,-3]",
			args: args{[]int{1, -2, -3}},
			want: 5,
		},
		{
			name: "nums = [1,2]",
			args: args{[]int{1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartValue(tt.args.nums); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
