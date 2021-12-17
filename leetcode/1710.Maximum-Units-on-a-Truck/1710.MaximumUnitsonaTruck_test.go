package _710_Maximum_Units_on_a_Truck

import "testing"

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4",
			args: args{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
