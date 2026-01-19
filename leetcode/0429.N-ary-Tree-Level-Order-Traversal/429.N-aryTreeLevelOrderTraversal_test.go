package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	n := func(val int, children ...*Node) *Node {
		return &Node{
			Val:      val,
			Children: children,
		}
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				root: n(1,
					n(3, n(5), n(6)),
					n(2),
					n(4),
				),
			},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			name: "example 2",
			args: args{
				root: n(1,
					n(2, n(6, n(11, n(14))), n(7, n(12))),
					n(3, n(8, n(13))),
					n(4, n(9), n(10)),
					n(5),
				),
			},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
