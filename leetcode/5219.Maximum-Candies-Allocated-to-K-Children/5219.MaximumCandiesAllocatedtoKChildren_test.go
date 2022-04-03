package _219_Maximum_Candies_Allocated_to_K_Children

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "candies = [5,8,6], k = 3",
			args: args{[]int{5, 8, 6}, 3},
			want: 5,
		},
		{
			name: "candies = [2,5], k = 11",
			args: args{[]int{2, 5}, 11},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
