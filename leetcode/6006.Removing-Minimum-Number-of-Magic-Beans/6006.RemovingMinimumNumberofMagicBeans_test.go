package _006_Removing_Minimum_Number_of_Magic_Beans

import "testing"

func Test_minimumRemoval(t *testing.T) {
	type args struct {
		beans []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "[4,1,6,5]",
			args: args{[]int{4, 1, 6, 5}},
			want: 4,
		},
		{
			name: "[2,10,3,2]",
			args: args{[]int{2, 10, 3, 2}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRemoval(tt.args.beans); got != tt.want {
				t.Errorf("minimumRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
