package _678_Valid_Parenthesis_String

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s = "(*))"`,
			args: args{"(*))"},
			want: true,
		},
		{
			name: `s = "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"`,
			args: args{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"},
			want: false,
		},
		{
			name: `s = "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"`,
			args: args{"((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
			if got := checkValidString2(tt.args.s); got != tt.want {
				t.Errorf("checkValidString2() = %v, want %v", got, tt.want)
			}
		})
	}
}
