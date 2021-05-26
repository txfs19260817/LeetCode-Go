package _738_Find_Kth_Largest_XOR_Coordinate_Value

import "testing"

func Test_kthLargestValue(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "matrix = [[5,2],[1,6]], k = 1",
			args: args{[][]int{{5, 2}, {1, 6}}, 1},
			want: 7,
		},
		{
			name: "matrix = [[5,2],[1,6]], k = 2",
			args: args{[][]int{{5, 2}, {1, 6}}, 2},
			want: 5,
		},
		{
			name: "matrix = [[5,2],[1,6]], k = 3",
			args: args{[][]int{{5, 2}, {1, 6}}, 3},
			want: 4,
		},
		{
			name: "matrix = [[5,2],[1,6]], k = 4",
			args: args{[][]int{{5, 2}, {1, 6}}, 4},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
