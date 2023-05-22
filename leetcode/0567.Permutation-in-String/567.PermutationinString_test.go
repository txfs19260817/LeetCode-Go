package leetcode

import (
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1 = `ab` s2 = `eidbaooo`",
			args: args{"ab", "eidbaooo"},
			want: true,
		},
		{
			name: "s1 = `ab` s2 = `eidboaoo`",
			args: args{"ab", "eidboaoo"},
			want: false,
		},
		{
			name: "s1 = `adc` s2 = `dcda`",
			args: args{"adc", "dcda"},
			want: true,
		},
		{
			name: "s1 = `hello` s2 = `ooolleoooleh`",
			args: args{"hello", "ooolleoooleh"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInclusion1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s1 = `ab` s2 = `eidbaooo`",
			args: args{"ab", "eidbaooo"},
			want: true,
		},
		{
			name: "s1 = `ab` s2 = `eidboaoo`",
			args: args{"ab", "eidboaoo"},
			want: false,
		},
		{
			name: "s1 = `adc` s2 = `dcda`",
			args: args{"adc", "dcda"},
			want: true,
		},
		{
			name: "s1 = `hello` s2 = `ooolleoooleh`",
			args: args{"hello", "ooolleoooleh"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion1(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkInclusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkInclusion("hello", "ooolleoooleh")
	}
}

func Benchmark_checkInclusion1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkInclusion1("hello", "ooolleoooleh")
	}
}
