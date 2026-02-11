package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper: feed all values through StreamEncoder, collect all emitted runs.
func encodeViaStream(values []int) []string {
	enc := NewStreamEncoder()
	var runs []string
	for _, v := range values {
		runs = append(runs, enc.Write(v)...)
	}
	runs = append(runs, enc.Flush()...)
	return runs
}

// helper: feed all runs through StreamDecoder, collect all decoded values.
func decodeViaStream(runs []string) []int {
	dec := NewStreamDecoder()
	var values []int
	for _, r := range runs {
		values = append(values, dec.Write(r)...)
	}
	return values
}

func TestStream_Example1(t *testing.T) {
	input := []int{5, 5, 5, 5, 5, 5, 5, 5, 1, 2, 3}
	encoded := encodeViaStream(input)
	assert.Equal(t, []string{"RLE[5,8]", "BP[1,2,3]"}, encoded)
	assert.Equal(t, input, decodeViaStream(encoded))
}

func TestStream_Example2(t *testing.T) {
	input := []int{1, 1, 1}
	encoded := encodeViaStream(input)
	assert.Equal(t, []string{"RLE[1,3]"}, encoded)
	assert.Equal(t, input, decodeViaStream(encoded))
}

func TestStream_Example3(t *testing.T) {
	input := []int{1, 1, 1, 1, 2, 3, 4, 5}
	encoded := encodeViaStream(input)
	assert.Equal(t, []string{"BP[1,1,1,1,2,3,4,5]"}, encoded)
	assert.Equal(t, input, decodeViaStream(encoded))
}

func TestStream_SingleElement(t *testing.T) {
	input := []int{42}
	encoded := encodeViaStream(input)
	assert.Equal(t, []string{"RLE[42,1]"}, encoded)
	assert.Equal(t, input, decodeViaStream(encoded))
}

func TestStream_LongMixed(t *testing.T) {
	// 10 sevens (RLE) + 3 distinct values + 9 twos (RLE)
	input := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1, 2, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	encoded := encodeViaStream(input)
	assert.Equal(t, []string{"RLE[7,10]", "BP[1,2,3]", "RLE[2,9]"}, encoded)
	assert.Equal(t, input, decodeViaStream(encoded))
}

func TestStream_Empty(t *testing.T) {
	encoded := encodeViaStream(nil)
	assert.Nil(t, encoded)
}

func TestStream_IncrementalEmission(t *testing.T) {
	// Verify that runs are emitted at the right moments, not all at Flush.
	enc := NewStreamEncoder()

	// Feed 8 fives — nothing emitted yet (run not finalized until a new value arrives).
	for i := 0; i < 8; i++ {
		assert.Nil(t, enc.Write(5))
	}
	// New value finalizes the run of 5s → RLE emitted.
	runs := enc.Write(1)
	assert.Equal(t, []string{"RLE[5,8]"}, runs)

	// Feed 2 more values, then flush.
	assert.Nil(t, enc.Write(2))
	assert.Nil(t, enc.Write(3))
	runs = enc.Flush()
	assert.Equal(t, []string{"BP[1,2,3]"}, runs)
}

func TestStream_BPFlushAt8(t *testing.T) {
	// 7 ones (below RLE threshold) + switch → fills BP to 7
	// Then 1 two fills BP to 8 → BP flushed mid-stream.
	enc := NewStreamEncoder()
	for i := 0; i < 7; i++ {
		assert.Nil(t, enc.Write(1))
	}
	// Switching to 2 finalizes the run of 1s (count=7 < 8 → goes to BP buffer, now 7 items).
	runs := enc.Write(2)
	assert.Nil(t, runs) // BP has 7, not yet 8

	// Switching to 3 finalizes the run of one 2 → BP buffer becomes 8 → flushed.
	runs = enc.Write(3)
	assert.Equal(t, []string{"BP[1,1,1,1,1,1,1,2]"}, runs)

	runs = enc.Flush()
	assert.Equal(t, []string{"RLE[3,1]"}, runs) // last-run exception: BP empty, emit RLE
}
