package _359_Count_All_Valid_Pickup_and_Delivery_Options

import "testing"

func Test_countOrders(t *testing.T) {
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
			want: 1,
		},
		{
			name: "2",
			args: args{2},
			want: 6,
		},
		{
			name: "8",
			args: args{8},
			want: 729647433,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOrders(tt.args.n); got != tt.want {
				t.Errorf("countOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
