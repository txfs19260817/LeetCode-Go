package _078_Rearrange_Characters_to_Make_Target_String

import "testing"

func Test_rearrangeCharacters(t *testing.T) {
	type args struct {
		s      string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "ilovecodingonleetcode", target = "code"`,
			args: args{"ilovecodingonleetcode", "code"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeCharacters(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("rearrangeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
