package _978_Longest_Turbulent_Subarray

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [9,4,2,10,7,8,8,1,9]",
			args: args{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}},
			want: 5,
		},
		{
			name: "arr = [4,8,12,16]",
			args: args{[]int{4, 8, 12, 16}},
			want: 2,
		},
		{
			name: "arr = [1]",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "arr = [1,9]",
			args: args{[]int{1, 9}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
