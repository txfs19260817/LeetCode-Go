package _443_String_Compression

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `chars = ["a","a","b","b","c","c","c"]`,
			args: args{[]byte{'a','a','b','b','c','c','c'}},
			want: 6,
		},
		{
			name: `chars = ["a"]`,
			args: args{[]byte{'a'}},
			want: 1,
		},
		{
			name: `chars = ["a","a","b","b","c","c","c"]`,
			args: args{[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}},
			want: 4,
		},
		{
			name: `chars = ["a","a","a","b","b","a","a"]`,
			args: args{[]byte{'a','a','a','b','b','a','a'}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
