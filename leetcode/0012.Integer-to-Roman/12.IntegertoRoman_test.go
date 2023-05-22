package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "39",
			args: args{39},
			want: "XXXIX",
		},
		{
			name: "1888",
			args: args{1888},
			want: "MDCCCLXXXVIII",
		},
		{
			name: "1976",
			args: args{1976},
			want: "MCMLXXVI",
		},
		{
			name: "3999",
			args: args{3999},
			want: "MMMCMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
