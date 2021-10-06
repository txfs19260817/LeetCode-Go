package _414_Third_Maximum_Number

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,2,1]",
			args: args{[]int{3, 2, 1}},
			want: 1,
		},
		{
			name: "nums = [1,2]",
			args: args{[]int{1, 2}},
			want: 2,
		},
		{
			name: "nums = [2,2,3,1]",
			args: args{[]int{2, 2, 3, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
