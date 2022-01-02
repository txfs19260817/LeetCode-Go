package _968_Number_of_Laser_Beams_in_a_Bank

import "testing"

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `bank = ["011001","000000","010100","001000"]`,
			args: args{[]string{"011001", "000000", "010100", "001000"}},
			want: 8,
		},
		{
			name: `bank = ["000","111","000"]`,
			args: args{[]string{"000", "111", "000"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
