package leetcode

import "testing"

func Test_mctFromLeafValues(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "arr = [6,2,4]",
			args: args{[]int{6, 2, 4}},
			want: 32,
		},
		{
			name: "arr = [4,11]",
			args: args{[]int{4, 11}},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mctFromLeafValues(tt.args.arr); got != tt.want {
				t.Errorf("mctFromLeafValues() = %v, want %v", got, tt.want)
			}
			if got := mctFromLeafValues2(tt.args.arr); got != tt.want {
				t.Errorf("mctFromLeafValues2() = %v, want %v", got, tt.want)
			}
		})
	}
}
