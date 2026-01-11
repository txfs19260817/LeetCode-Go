package snowflake

import "testing"

func Test_threeColor(t *testing.T) {
	type args struct {
		locations   []string
		adjacencies [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1: Valid coloring exists",
			args: args{
				locations: []string{"A", "B", "C", "D"},
				adjacencies: [][]string{
					{"A", "B"}, {"A", "C"}, {"B", "C"}, {"C", "D"},
				},
			},
			want: true,
		},
		{
			name: "Example 2: No valid coloring (K4 complete graph)",
			args: args{
				locations: []string{"A", "B", "C", "D"},
				adjacencies: [][]string{
					{"A", "B"}, {"A", "C"}, {"A", "D"},
					{"B", "C"}, {"B", "D"}, {"C", "D"},
				},
			},
			want: false,
		},
		{
			name: "Example 3: Valid coloring (Tree structure)",
			args: args{
				locations: []string{"A", "B", "C", "D", "E", "F"},
				adjacencies: [][]string{
					{"A", "B"}, {"A", "C"}, {"B", "D"}, {"B", "E"}, {"C", "F"},
				},
			},
			want: true,
		},
		{
			name: "Cycle 5: Valid coloring",
			args: args{
				locations: []string{"A", "B", "C", "D", "E"},
				adjacencies: [][]string{
					{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "A"},
				},
			},
			want: true,
		},
		{
			name: "Path graph: Valid coloring",
			args: args{
				locations: []string{"A", "B", "C", "D", "E"},
				adjacencies: [][]string{
					{"A", "B"}, {"B", "C"}, {"C", "D"},
				},
			},
			want: true,
		},
		{
			name: "Disconnected components: Valid coloring",
			args: args{
				locations: []string{"A", "B", "C", "D"},
				adjacencies: [][]string{
					{"A", "B"}, {"C", "D"},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeColor(tt.args.locations, tt.args.adjacencies); got != tt.want {
				t.Errorf("threeColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
