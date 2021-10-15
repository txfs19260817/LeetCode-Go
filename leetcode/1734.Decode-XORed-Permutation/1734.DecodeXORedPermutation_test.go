package _734_Decode_XORed_Permutation

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		encoded []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "encoded = [3,1]",
			args: args{[]int{3, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "encoded = [6,5,4,6]",
			args: args{[]int{6, 5, 4, 6}},
			want: []int{2, 4, 1, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.encoded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
