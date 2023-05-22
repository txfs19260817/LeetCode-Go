package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `num = "123", target = 6`,
			args: args{"123", 6},
			want: []string{"1*2*3", "1+2+3"},
		},
		{
			name: `num = "00", target = 0`,
			args: args{"00", 0},
			want: []string{"0*0", "0+0", "0-0"},
		},
		{
			name: `num = "3456237490", target = 9191`,
			args: args{"3456237490", 9191},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, addOperators(tt.args.num, tt.args.target))
		})
	}
}
