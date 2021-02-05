package _845_Longest_Mountain_in_Array

import "testing"

func Test_longestMountain(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [2,1,4,7,3,2,5]",
			args: args{[]int{2, 1, 4, 7, 3, 2, 5}},
			want: 5,
		},
		{
			name: "arr = [2,2,2]",
			args: args{[]int{2, 2, 2}},
			want: 0,
		},
		{
			name: "arr = [875,884,239,731,723,685]",
			args: args{[]int{875, 884, 239, 731, 723, 685}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.arr); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
