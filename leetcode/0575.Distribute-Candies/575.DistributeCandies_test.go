package _575_Distribute_Candies

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candyType []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "candyType = [1,1,2,2,3,3]",
			args: args{[]int{1, 1, 2, 2, 3, 3}},
			want: 3,
		},
		{
			name: "candyType = [1,1,2,3]",
			args: args{[]int{1, 1, 2, 3}},
			want: 2,
		},
		{
			name: "candyType = [6,6,6,6]",
			args: args{[]int{6, 6, 6, 6}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candyType); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
