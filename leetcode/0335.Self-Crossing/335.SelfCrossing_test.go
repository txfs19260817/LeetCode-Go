package _335_Self_Crossing

import "testing"

func Test_isSelfCrossing(t *testing.T) {
	type args struct {
		distance []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[2,1,1,2]",
			args: args{[]int{2, 1, 1, 2}},
			want: true,
		},
		{
			name: "[1,1,1,1]",
			args: args{[]int{1, 1, 1, 1}},
			want: true,
		},
		{
			name: "[1,1,2,1,1]",
			args: args{[]int{1, 1, 2, 1, 1}},
			want: true,
		},
		{
			name: "[1,1,2,2,1,1]",
			args: args{[]int{1, 1, 2, 2, 1, 1}},
			want: true,
		},
		{
			name: "[1,2,3,4]",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSelfCrossing(tt.args.distance); got != tt.want {
				t.Errorf("isSelfCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
