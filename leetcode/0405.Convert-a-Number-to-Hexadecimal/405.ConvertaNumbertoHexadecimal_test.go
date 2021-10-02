package _405_Convert_a_Number_to_Hexadecimal

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "num = 26",
			args: args{26},
			want: "1a",
		},
		{
			name: "num = -1",
			args: args{-1},
			want: "ffffffff",
		},
		{
			name: "num = 0",
			args: args{0},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
			if got := toHex2(tt.args.num); got != tt.want {
				t.Errorf("toHex2() = %v, want %v", got, tt.want)
			}
		})
	}
}
