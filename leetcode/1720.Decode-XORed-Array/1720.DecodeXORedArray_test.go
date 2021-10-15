package _720_Decode_XORed_Array

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		encoded []int
		first   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "encoded = [1,2,3], first = 1",
			args: args{[]int{1, 2, 3}, 1},
			want: []int{1, 0, 2, 1},
		},
		{
			name: "encoded = [6,2,7,3], first = 4",
			args: args{[]int{6, 2, 7, 3}, 4},
			want: []int{4, 2, 0, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.encoded, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
