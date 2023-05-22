package leetcode

import "testing"

func Test_largestGoodInteger(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "6777133339",
			args: args{"6777133339"},
			want: "777",
		},
		{
			name: "777",
			args: args{"777"},
			want: "777",
		},
		{
			name: "7776",
			args: args{"7776"},
			want: "777",
		},
		{
			name: "6777",
			args: args{"6777"},
			want: "777",
		},
		{
			name: "42352338",
			args: args{"42352338"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestGoodInteger(tt.args.num); got != tt.want {
				t.Errorf("largestGoodInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
