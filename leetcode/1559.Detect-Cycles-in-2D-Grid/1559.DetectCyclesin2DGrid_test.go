package _559_Detect_Cycles_in_2D_Grid

import "testing"

func Test_containsCycle(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `grid = [["a","a","a","a"],["a","b","b","a"],["a","b","b","a"],["a","a","a","a"]]`,
			args: args{[][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}},
			want: true,
		},
		{
			name: `grid = [["c","c","c","a"],["c","d","c","c"],["c","c","e","c"],["f","c","c","c"]]`,
			args: args{[][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}},
			want: true,
		},
		{
			name: `grid = [["a","b","b"],["b","z","b"],["b","b","a"]]`,
			args: args{[][]byte{{'a', 'a', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsCycle(tt.args.grid); got != tt.want {
				t.Errorf("containsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
