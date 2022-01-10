package _891_Cutting_Ribbons

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		ribbons []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ribbons = [9,7,5], k = 3",
			args: args{[]int{9, 7, 5}, 3},
			want: 5,
		},
		{
			name: "ribbons = [5,7,9], k = 22",
			args: args{[]int{5, 7, 9}, 22},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.ribbons, tt.args.k); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
