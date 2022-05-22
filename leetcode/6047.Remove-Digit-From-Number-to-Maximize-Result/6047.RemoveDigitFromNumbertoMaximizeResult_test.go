package _047_Remove_Digit_From_Number_to_Maximize_Result

import "testing"

func Test_removeDigit(t *testing.T) {
	type args struct {
		number string
		digit  byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `"133235","3"`,
			args: args{"133235", '3'},
			want: "13325",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDigit(tt.args.number, tt.args.digit); got != tt.want {
				t.Errorf("removeDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
