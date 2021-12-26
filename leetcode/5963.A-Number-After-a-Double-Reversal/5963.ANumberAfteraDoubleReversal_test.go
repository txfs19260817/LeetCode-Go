package _963_A_Number_After_a_Double_Reversal

import "testing"

func Test_isSameAfterReversals(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "526",
			args: args{526},
			want: true,
		},
		{
			name: "1800",
			args: args{1800},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameAfterReversals(tt.args.num); got != tt.want {
				t.Errorf("isSameAfterReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}
