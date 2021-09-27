package _162_Find_Peak_Element

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,2,1,3,5,6,4]",
			args: args{[]int{1, 2, 1, 3, 5, 6, 4}},
			want: 5,
		},
		{
			name: "nums = [1]",
			args: args{[]int{1}},
			want: 0,
		},
		{
			name: "nums = [2,1]",
			args: args{[]int{2, 1}},
			want: 0,
		},
		{
			name: "nums = [1,2]",
			args: args{[]int{1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
