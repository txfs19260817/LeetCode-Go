package _667_Beautiful_Arrangement_II

import (
	"reflect"
	"testing"
)

func Test_constructArray(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n = 3, k = 1",
			args: args{3, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "n = 3, k = 2",
			args: args{3, 2},
			want: []int{1, 3, 2},
		},
		{
			name: "n = 5, k = 2",
			args: args{5, 2},
			want: []int{1, 3, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArray(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
