package _027_Remove_Element

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,2,2,3], val = 3",
			args: args{[]int{3, 2, 2, 3}, 3},
			want: 2,
		},
		{
			name: "nums = [2], val = 3",
			args: args{[]int{2}, 3},
			want: 1,
		},
		{
			name: "nums = [0,1,2,2,3,0,4,2], val = 2",
			args: args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
