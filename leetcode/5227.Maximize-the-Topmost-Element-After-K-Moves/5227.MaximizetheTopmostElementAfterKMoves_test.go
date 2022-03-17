package _227_Maximize_the_Topmost_Element_After_K_Moves

import "testing"

func Test_maximumTop(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [2], k = 1",
			args: args{[]int{2}, 1},
			want: -1,
		},
		{
			name: "[94,23,58,...],79",
			args: args{[]int{94, 23, 58, 92, 3, 63, 68, 43, 8, 97, 54, 11, 63, 55, 73, 38, 22, 80, 45, 43, 25, 34, 67, 76, 80, 9, 30, 27, 50, 7, 57, 63, 63, 27, 46, 1, 50, 55, 99, 92, 73, 9, 57, 25, 59, 54, 100, 56, 64, 94, 99}, 79},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTop(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
