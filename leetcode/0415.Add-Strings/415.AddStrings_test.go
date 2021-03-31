package _415_Add_Strings

import "testing"

func Test_addStrings(t *testing.T) {
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
			name: `num1 = "11", num2 = "123"`,
			args: args{"11", "123"},
			want: "134",
		},
		{
			name: `num1 = "0", num2 = "0"`,
			args: args{"0", "0"},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
