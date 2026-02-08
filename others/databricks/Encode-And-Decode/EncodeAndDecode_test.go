package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode_Example1(t *testing.T) {
	input := []int{5, 5, 5, 5, 5, 5, 5, 5, 1, 2, 3}
	encoded := Encode(input)
	assert.Equal(t, []string{"RLE[5,8]", "BP[1,2,3]"}, encoded)
	assert.Equal(t, input, Decode(encoded))
}

func TestEncode_Example2(t *testing.T) {
	input := []int{1, 1, 1}
	encoded := Encode(input)
	assert.Equal(t, []string{"RLE[1,3]"}, encoded)
	assert.Equal(t, input, Decode(encoded))
}

func TestEncode_Example3(t *testing.T) {
	input := []int{1, 1, 1, 1, 2, 3, 4, 5}
	encoded := Encode(input)
	assert.Equal(t, []string{"BP[1,1,1,1,2,3,4,5]"}, encoded)
	assert.Equal(t, input, Decode(encoded))
}

func TestEncode_SingleElement(t *testing.T) {
	input := []int{42}
	encoded := Encode(input)
	assert.Equal(t, []string{"RLE[42,1]"}, encoded)
	assert.Equal(t, input, Decode(encoded))
}

func TestEncode_LongMixed(t *testing.T) {
	// 10 sevens (RLE) + 3 distinct values + 9 twos (RLE)
	input := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1, 2, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	encoded := Encode(input)
	assert.Equal(t, []string{"RLE[7,10]", "BP[1,2,3]", "RLE[2,9]"}, encoded)
	assert.Equal(t, input, Decode(encoded))
}
