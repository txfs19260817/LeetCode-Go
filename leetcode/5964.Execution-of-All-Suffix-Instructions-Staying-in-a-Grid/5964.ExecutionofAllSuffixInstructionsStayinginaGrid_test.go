package _964_Execution_of_All_Suffix_Instructions_Staying_in_a_Grid

import (
	"reflect"
	"testing"
)

func Test_executeInstructions(t *testing.T) {
	type args struct {
		n        int
		startPos []int
		s        string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `n = 3, startPos = [0,1], s = "RRDDLU"`,
			args: args{3, []int{0, 1}, "RRDDLU"},
			want: []int{1, 5, 4, 3, 1, 0},
		},
		{
			name: `n = 1, startPos = [0,0], s = "LRUD"`,
			args: args{1, []int{0, 0}, "LRUD"},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeInstructions(tt.args.n, tt.args.startPos, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
