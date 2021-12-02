package am

import (
	"reflect"
	"testing"
)

func Test_similarities(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ababa",
			args: args{"ababa"},
			want: []int{0,0,3,0,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarities(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("similarities() = %v, want %v", got, tt.want)
			}
		})
	}
}
