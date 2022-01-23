package _151_Maximum_Good_People_Based_on_Statements

import "testing"

func Test_maximumGood(t *testing.T) {
	type args struct {
		statements [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[2,1,2],[1,2,2],[2,0,2]]",
			args: args{[][]int{{2, 1, 2}, {1, 2, 2}, {2, 0, 2}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGood(tt.args.statements); got != tt.want {
				t.Errorf("maximumGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
