package leetcode

import "testing"

func Test_discountPrices(t *testing.T) {
	type args struct {
		sentence string
		discount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "duew$11mengf $8 $1",
			args: args{"duew$11mengf $8 $1", 7},
			want: "duew$11mengf $7.44 $0.93",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := discountPrices(tt.args.sentence, tt.args.discount); got != tt.want {
				t.Errorf("discountPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
