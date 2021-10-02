package _399_Evaluate_Division

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: `equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]`,
			args: args{
				[][]string{{"a", "b"}, {"b", "c"}},
				[]float64{2.0, 3.0},
				[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}},
			want: []float64{6., 0.5, -1., 1., -1.},
		},
		{
			name: `equations = ...`,
			args: args{
				[][]string{{"a", "b"}, {"a", "c"}, {"a", "d"}, {"a", "e"}, {"a", "f"}, {"a", "g"}, {"a", "h"}, {"a", "i"}, {"a", "j"}, {"a", "k"}, {"a", "l"}, {"a", "aa"}, {"a", "aaa"}, {"a", "aaaa"}, {"a", "aaaaa"}, {"a", "bb"}, {"a", "bbb"}, {"a", "ff"}},
				[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 1.0, 1.0, 1.0, 1.0, 1.0, 3.0, 5.0},
				[][]string{{"d", "f"}, {"e", "g"}, {"e", "k"}, {"h", "a"}, {"aaa", "k"}, {"aaa", "i"}, {"aa", "e"}, {"aaa", "aa"}, {"aaa", "ff"}, {"bbb", "bb"}, {"bb", "h"}, {"bb", "i"}, {"bb", "k"}, {"aaa", "k"}, {"k", "l"}, {"x", "k"}, {"l", "ll"}},
			},
			want: []float64{6., 0.5, -1., 1., -1.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
