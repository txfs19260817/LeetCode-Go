package _013_Roman_to_Integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "III",
			args: args{"III"},
			want: 3,
		},
		{
			name: "IV",
			args: args{"IV"},
			want: 4,
		},
		{
			name: "XXXIX",
			args: args{"XXXIX"},
			want: 39,
		},
		{
			name: "MDCCCLXXXVIII",
			args: args{"MDCCCLXXXVIII"},
			want: 1888,
		},
		{
			name: "MCMLXXVI",
			args: args{"MCMLXXVI"},
			want: 1976,
		},
		{
			name: "MMMCMXCIX",
			args: args{"MMMCMXCIX"},
			want: 3999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
