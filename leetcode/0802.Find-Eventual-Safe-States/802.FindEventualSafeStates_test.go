package _802_Find_Eventual_Safe_States

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "graph = [[1,2],[2,3],[5],[0],[5],[],[]]",
			args: args{[][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			want: []int{2, 4, 5, 6},
		},
		{
			name: "graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]",
			args: args{[][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
			if got := eventualSafeNodes2(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_eventualSafeNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})
	}
}

func Benchmark_eventualSafeNodes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eventualSafeNodes2([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})
	}
}
