package _907_Next_Greater_Numerically_Balanced_Number

import "testing"

func Test_nextBeautifulNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 22,
		},
		{
			name: "1000",
			args: args{1000},
			want: 1333,
		},
		{
			name: "3000",
			args: args{3000},
			want: 3133,
		},
		{
			name: "666666",
			args: args{666666},
			want: 1224444,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextBeautifulNumber(tt.args.n); got != tt.want {
				t.Errorf("nextBeautifulNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
