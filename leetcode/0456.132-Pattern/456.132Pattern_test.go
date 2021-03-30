package _456_132_Pattern

import "testing"

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,4]",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "[3,5,0,3,4]",
			args: args{[]int{3, 5, 0, 3, 4}},
			want: true,
		},
		{
			name: "[1,0,1,-4,-3]",
			args: args{[]int{1, 0, 1, -4, -3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
