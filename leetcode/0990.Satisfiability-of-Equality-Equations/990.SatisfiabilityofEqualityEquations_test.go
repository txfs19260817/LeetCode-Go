package _990_Satisfiability_of_Equality_Equations

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `["a==b","b!=a"]`,
			args: args{[]string{"a==b", "b!=a"}},
			want: false,
		},
		{
			name: `["b==a","a==b"]`,
			args: args{[]string{"b==a", "a==b"}},
			want: true,
		},
		{
			name: `["a==b","b==c","a==c"]`,
			args: args{[]string{"a==b", "b==c", "a==c"}},
			want: true,
		},
		{
			name: `["a==b","b!=c","c==a"]`,
			args: args{[]string{"a==b", "b!=c", "c==a"}},
			want: false,
		},
		{
			name: `["c==c","b==d","x!=z"]`,
			args: args{[]string{"c==c", "b==d", "x!=z"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
