package _470_Implement_Rand10_Using_Rand7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rand10(t *testing.T) {
	for i := 0; i < 114514; i++ {
		r := rand10()
		assert.True(t, 1 <= r && r <= 10)
	}
}

func Test_rand10_1(t *testing.T) {
	for i := 0; i < 114514; i++ {
		r := rand10()
		assert.True(t, 1 <= r && r <= 10)
	}
}

func Benchmark_rand10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand10()
	}
}

func Benchmark_rand10_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand10_1()
	}
}
