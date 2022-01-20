package _249_Group_Shifted_Strings

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_groupStrings(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "abc",
			args: args{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}},
			want: [][]string{{"abc", "bcd", "xyz"}, {"acef"}, {"a", "z"}, {"az", "ba"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupStrings(tt.args.strings)
			out, want := make([][]string, len(got)), make([][]string, len(tt.want))
			for i, ss := range got {
				sort.Strings(ss)
				out[i] = ss
			}
			for i, ss := range tt.want {
				sort.Strings(ss)
				want[i] = ss
			}

			assert.ElementsMatch(t, out, want)
		})
	}
}
