package _965_Intervals_Between_Identical_Elements

import (
	"reflect"
	"testing"
)

func Test_getDistances(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "[2,1,3,1,2,3,3]",
			args: args{[]int{2, 1, 3, 1, 2, 3, 3}},
			want: []int64{4, 2, 7, 2, 4, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistances(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
