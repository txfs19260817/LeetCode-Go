package leetcode

import "testing"

func Test_shoppingOffers(t *testing.T) {
	type args struct {
		price   []int
		special [][]int
		needs   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]",
			args: args{[]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}},
			want: 14,
		},
		{
			name: "price = [2,3,4], special = [[1,1,0,4],[2,2,1,9]], needs = [1,2,1]",
			args: args{[]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}},
			want: 11,
		},
		{
			name: "price = [9,9], special = [[1,1,1]], needs = [2,2]",
			args: args{[]int{9, 9}, [][]int{{1, 1, 1}}, []int{2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shoppingOffers(tt.args.price, tt.args.special, tt.args.needs); got != tt.want {
				t.Errorf("shoppingOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}
