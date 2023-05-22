package leetcode

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		startValue int
		target     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "startValue = 2, target = 3",
			args: args{2, 3},
			want: 2,
		},
		{
			name: "startValue = 1024, target = 1",
			args: args{1024, 1},
			want: 1023,
		},
		{
			name: "startValue = 1, target = 1000000000",
			args: args{1, 1000000000},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.args.startValue, tt.args.target); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
