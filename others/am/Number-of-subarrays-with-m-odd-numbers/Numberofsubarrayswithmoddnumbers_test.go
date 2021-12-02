package am

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		arr []int
		m   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[]int{2, 2, 5, 6, 9, 2, 11}, 2},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.arr, tt.args.m); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
