package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wants []string
	}{
		{
			name:  "tree",
			args:  args{"tree"},
			wants: []string{"eetr", "eert"},
		},
		{
			name:  "cccaaa",
			args:  args{"cccaaa"},
			wants: []string{"cccaaa", "aaaccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.wants, frequencySort(tt.args.s))
			assert.Contains(t, tt.wants, frequencySort2(tt.args.s))
		})
	}
}

func Benchmark_frequencySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort(randomString(b.N))
	}
}

func Benchmark_frequencySort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySort2(randomString(b.N))
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
