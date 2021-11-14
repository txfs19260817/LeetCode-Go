package _928_Decode_the_Slanted_Ciphertext

import "testing"

func Test_decodeCiphertext(t *testing.T) {
	type args struct {
		encodedText string
		rows        int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `encodedText = "ch   ie   pr", rows = 3`,
			args: args{"ch   ie   pr",3},
			want: "cipher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCiphertext(tt.args.encodedText, tt.args.rows); got != tt.want {
				t.Errorf("decodeCiphertext() = %v, want %v", got, tt.want)
			}
		})
	}
}
