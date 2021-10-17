package _904_Count_Number_of_Maximum_Bitwise_OR_Subsets

import "testing"

func Test_countMaxOrSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,2,1,5",
			args: args{[]int{3,2,1,5}},
			want: 6,
		},
		{
			name: "2,2,2",
			args: args{[]int{2,2,2}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
