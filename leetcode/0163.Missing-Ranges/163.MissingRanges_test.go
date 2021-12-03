package _163_Missing_Ranges

import (
	"reflect"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nums = [0,1,3,50,75], lower = 0, upper = 99",
			args: args{[]int{0, 1, 3, 50, 75}, 0, 99},
			want: []string{"2", "4->49", "51->74", "76->99"},
		},
		{
			name: "nums = [], lower = 1, upper = 1",
			args: args{[]int{}, 1, 1},
			want: []string{"1"},
		},
		{
			name: "nums = [], lower = -3, upper = -1",
			args: args{[]int{}, -3, -1},
			want: []string{"-3->-1"},
		},
		{
			name: "nums = [-1], lower = -1, upper = -1",
			args: args{[]int{-1}, -1, -1},
			want: nil,
		},
		{
			name: "nums = [-1], lower = -2, upper = -1",
			args: args{[]int{-1}, -2, -1},
			want: []string{"-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
