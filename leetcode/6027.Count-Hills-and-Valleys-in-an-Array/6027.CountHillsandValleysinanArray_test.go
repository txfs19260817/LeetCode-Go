package _027_Count_Hills_and_Valleys_in_an_Array

import "testing"

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[6,6,5,5,4,1]",
			args: args{[]int{6, 6, 5, 5, 4, 1}},
			want: 0,
		},
		{
			name: "[8,2,5,7,7,2,10,3,6,2]",
			args: args{[]int{8, 2, 5, 7, 7, 2, 10, 3, 6, 2}},
			want: 6,
		},
		{
			name: "[85,52,89,81,48,8,18,12,88,20,70,100,67,42,12,95,57,8,41,82,37,44,47,18,46]",
			args: args{[]int{85, 52, 89, 81, 48, 8, 18, 12, 88, 20, 70, 100, 67, 42, 12, 95, 57, 8, 41, 82, 37, 44, 47, 18, 46}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
