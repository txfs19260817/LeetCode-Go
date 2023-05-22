package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "S = `ababcbacadefegdehijhklij`",
			args: args{"ababcbacadefegdehijhklij"},
			want: []int{9, 7, 8},
		},
		{
			name: "S = `eccbbbbdec`",
			args: args{"eccbbbbdec"},
			want: []int{10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
			if got := partitionLabels2(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_partitionLabels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionLabels("ababcbacadefegdehijhklij")
	}
}

func Benchmark_partitionLabels2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionLabels2("ababcbacadefegdehijhklij")
	}
}
