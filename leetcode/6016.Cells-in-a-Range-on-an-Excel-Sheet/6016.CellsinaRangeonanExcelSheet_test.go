package _016_Cells_in_a_Range_on_an_Excel_Sheet

import (
	"reflect"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "K1:L2",
			args: args{"K1:L2"},
			want: []string{"K1", "K2", "L1", "L2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
