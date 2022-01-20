package _011_Capacity_To_Ship_Packages_Within_D_Days

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "weights = [1,2,3,4,5,6,7,8,9,10], days = 5",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
			want: 15,
		},
		{
			name: "weights = [1,2,3,1,1], days = 4",
			args: args{[]int{1, 2, 3, 1, 1}, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
