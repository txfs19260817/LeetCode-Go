package _925_Long_Pressed_Name

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "name = `alex`, typed = `aaleex`",
			args: args{"alex", "aaleex"},
			want: true,
		},
		{
			name: "name = `saeed`, typed = `ssaaedd`",
			args: args{"saeed", "ssaaedd"},
			want: false,
		},
		{
			name: "name = `leelee`, typed = `lleeelee`",
			args: args{"leelee", "lleeelee"},
			want: true,
		},
		{
			name: "name = `laiden`, typed = `laiden`",
			args: args{"laiden", "laiden"},
			want: true,
		},
		{
			name: "name = `a`, typed = `aa`",
			args: args{"a", "aa"},
			want: true,
		},
		{
			name: "name = `a`, typed = `aab`",
			args: args{"a", "aab"},
			want: false,
		},
		{
			name: "name = `pyplrz`, typed = `ppyypllr`",
			args: args{"pyplrz", "ppyypllr"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
