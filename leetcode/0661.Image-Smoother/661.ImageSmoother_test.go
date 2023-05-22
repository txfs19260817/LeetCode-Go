package leetcode

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "{{100,200,100},{200,50,200},{100,200,100}}",
			args: args{[][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
