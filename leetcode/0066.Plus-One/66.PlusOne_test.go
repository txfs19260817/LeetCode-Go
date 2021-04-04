package _066_Plus_One

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "digits = [4,3,2,1]",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "digits = [9,9]",
			args: args{[]int{9, 9}},
			want: []int{1, 0, 0},
		},
		{
			name: "digits = [8,9,9,9]",
			args: args{[]int{8, 9, 9, 9}},
			want: []int{9, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
