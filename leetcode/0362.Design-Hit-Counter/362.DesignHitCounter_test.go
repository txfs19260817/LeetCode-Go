package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	hc := Constructor()
	for i := 0; i < N; i++ {
		assert.Zero(t, hc.ts[i], "ts[%d] should start at zero", i)
		assert.Zero(t, hc.hits[i], "hits[%d] should start at zero", i)
	}
}

func TestHitCounterExample(t *testing.T) {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)

	assert.Equal(t, 3, hc.GetHits(4))

	hc.Hit(300)
	assert.Equal(t, 4, hc.GetHits(300))
	assert.Equal(t, 3, hc.GetHits(301))
}

func TestHitCounterSameTimestampAndExpiry(t *testing.T) {
	hc := Constructor()
	for i := 0; i < 5; i++ {
		hc.Hit(10)
	}

	assert.Equal(t, 5, hc.GetHits(10))
	assert.Equal(t, 5, hc.GetHits(309))
	assert.Equal(t, 0, hc.GetHits(310))
}

func TestHitCounterBucketOverwrite(t *testing.T) {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(301)

	assert.Equal(t, 1, hc.GetHits(301))
	assert.Equal(t, 1, hc.GetHits(302))
}
