package _482_Minimum_Number_of_Days_to_Make_m_Bouquets

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bloomDay = [1,10,3,10,2], m = 3, k = 1",
			args: args{[]int{1, 10, 3, 10, 2}, 3, 1},
			want: 3,
		},
		{
			name: "bloomDay = [1,10,3,10,2], m = 3, k = 2",
			args: args{[]int{1, 10, 3, 10, 2}, 3, 2},
			want: -1,
		},
		{
			name: "bloomDay = [7,7,7,7,12,7,7], m = 2, k = 3",
			args: args{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3},
			want: 12,
		},
		{
			name: "bloomDay = [1000000000,1000000000], m = 1, k = 1",
			args: args{[]int{1000000000, 1000000000}, 1, 1},
			want: 1000000000,
		},
		{
			name: "bloomDay = [1,10,2,9,3,8,4,7,5,6], m = 4, k = 2",
			args: args{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
