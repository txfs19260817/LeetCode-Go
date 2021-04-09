package _640_Check_Array_Formation_Through_Concatenation

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "arr = [15,88], pieces = [[88],[15]]",
			args: args{[]int{15, 88}, [][]int{{88}, {15}}},
			want: true,
		},
		{
			name: "arr = [91,4,64,78], pieces = [[78],[4,64],[91]]",
			args: args{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}},
			want: true,
		},
		{
			name: "arr = [49,18,16], pieces = [[16,18,49]]",
			args: args{[]int{49, 18, 16}, [][]int{{16, 18, 49}}},
			want: false,
		},
		{
			name: "arr = [15,88], pieces = [[88]]",
			args: args{[]int{15, 88}, [][]int{{88}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
