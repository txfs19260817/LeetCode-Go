package _565_Array_Nesting

import "testing"

func Test_arrayNesting(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [5,4,0,3,1,6,2]",
			args: args{[]int{5, 4, 0, 3, 1, 6, 2}},
			want: 4,
		},
		{
			name: "nums = [0,1,2]",
			args: args{[]int{0, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayNesting(tt.args.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
