package _043_Multiply_Strings

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `num1 = "123", num2 = "456"`,
			args: args{"123", "456"},
			want: "56088",
		},
		{
			name: `num1 = "498828660196", num2 = "840477629533"`,
			args: args{"498828660196", "840477629533"},
			want: "419254329864656431168468",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
