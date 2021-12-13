package _954_Watering_Plants_II

import "testing"

func Test_minimumRefill(t *testing.T) {
	type args struct {
		plants    []int
		capacityA int
		capacityB int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "plants = [2,2,5,2,2], capacityA = 5, capacityB = 5",
			args: args{[]int{2, 2, 5, 2, 2}, 5, 5},
			want: 1,
		},
		{
			name: "plants = [2,2,5,2,2], capacityA = 5, capacityB = 6",
			args: args{[]int{2, 2, 5, 2, 2}, 5, 6},
			want: 1,
		},
		{
			name: "plants = [2,2,3,3], capacityA = 3, capacityB = 4",
			args: args{[]int{2, 2, 3, 3}, 3, 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRefill(tt.args.plants, tt.args.capacityA, tt.args.capacityB); got != tt.want {
				t.Errorf("minimumRefill() = %v, want %v", got, tt.want)
			}
		})
	}
}
