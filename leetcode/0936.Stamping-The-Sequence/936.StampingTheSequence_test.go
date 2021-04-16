package _936_Stamping_The_Sequence

import (
	"reflect"
	"testing"
)

func Test_movesToStamp(t *testing.T) {
	type args struct {
		stamp  string
		target string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `stamp = "abc", target = "ababc"`,
			args: args{"abc", "ababc"},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToStamp(tt.args.stamp, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movesToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
