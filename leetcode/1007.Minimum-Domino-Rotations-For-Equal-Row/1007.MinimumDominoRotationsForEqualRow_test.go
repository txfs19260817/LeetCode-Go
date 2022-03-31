package _007_Minimum_Domino_Rotations_For_Equal_Row

import "testing"

func Test_minDominoRotations(t *testing.T) {
	type args struct {
		tops    []int
		bottoms []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,2,1,1,1,2,2,2",
			args: args{[]int{1, 2, 1, 1, 1, 2, 2, 2}, []int{2, 1, 2, 2, 2, 2, 2, 2}},
			want: 1,
		},
		{
			name: "3,5,1,2,3",
			args: args{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations(tt.args.tops, tt.args.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
