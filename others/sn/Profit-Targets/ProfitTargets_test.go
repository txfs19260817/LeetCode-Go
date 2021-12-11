package sn

import "testing"

func Test_profitTargets(t *testing.T) {
	type args struct {
		stocksProfit []int
		target       int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "1",
			args:    args{[]int{5, 7, 9, 13, 11, 6, 6, 3, 3}, 12},
			wantAns: 3,
		},
		{
			name:    "1ext",
			args:    args{[]int{3, 6, 3, 6, 5, 7, 7, 5, 9, 13, 11, 6, 6, 3, 3}, 12},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := profitTargets(tt.args.stocksProfit, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("profitTargets() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
