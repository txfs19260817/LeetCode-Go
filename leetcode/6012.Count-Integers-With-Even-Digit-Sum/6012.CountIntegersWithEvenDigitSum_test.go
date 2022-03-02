package _012_Count_Integers_With_Even_Digit_Sum

import "testing"

func Test_countEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{4},
			want: 2,
		},
		{
			name: "30",
			args: args{30},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEven(tt.args.num); got != tt.want {
				t.Errorf("countEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
