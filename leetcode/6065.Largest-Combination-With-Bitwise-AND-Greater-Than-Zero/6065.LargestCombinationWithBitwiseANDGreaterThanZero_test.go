package _065_Largest_Combination_With_Bitwise_AND_Greater_Than_Zero

import "testing"

func Test_largestCombination(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "candidates = [16,17,71,62,12,24,14]",
			args: args{[]int{16, 17, 71, 62, 12, 24, 14}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestCombination(tt.args.candidates); got != tt.want {
				t.Errorf("largestCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
