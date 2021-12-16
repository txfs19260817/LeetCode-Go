package _273_Integer_to_English_Words

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{0},
			want: "Zero",
		},
		{
			name: "1234567891",
			args: args{1234567891},
			want: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
		{
			name: "15",
			args: args{15},
			want: "Fifteen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
