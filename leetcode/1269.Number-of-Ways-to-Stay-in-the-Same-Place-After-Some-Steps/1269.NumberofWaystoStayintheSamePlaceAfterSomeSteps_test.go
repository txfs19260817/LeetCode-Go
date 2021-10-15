package _269_Number_of_Ways_to_Stay_in_the_Same_Place_After_Some_Steps

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		steps  int
		arrLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "steps = 3, arrLen = 2",
			args: args{3, 2},
			want: 4,
		},
		{
			name: "steps = 2, arrLen = 4",
			args: args{2, 4},
			want: 2,
		},
		{
			name: "steps = 4, arrLen = 2",
			args: args{4, 2},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.steps, tt.args.arrLen); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
