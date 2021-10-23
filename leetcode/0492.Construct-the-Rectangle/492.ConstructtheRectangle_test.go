package _492_Construct_the_Rectangle

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "4",
			args: args{4},
			want: []int{2, 2},
		},
		{
			name: "2",
			args: args{2},
			want: []int{2, 1},
		},
		{
			name: "37",
			args: args{37},
			want: []int{37, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
