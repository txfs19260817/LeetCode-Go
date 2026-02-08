package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPath_Example1(t *testing.T) {
	assert.Equal(t, "UUURL", FindPath(5, 5, 7))
}

func TestFindPath_Example2(t *testing.T) {
	assert.Equal(t, "UUULR", FindPath(4, 8, 3))
}

func TestFindPath_Example3(t *testing.T) {
	assert.Equal(t, "UUURRRL", FindPath(5, 4, 13))
}

func TestFindPath_SameNode(t *testing.T) {
	assert.Equal(t, "", FindPath(5, 3, 3))
	assert.Equal(t, "", FindPath(3, 0, 0))
}

func TestFindPath_RootToLeaf(t *testing.T) {
	// order=5 has 15 nodes (0..14). Rightmost leaf is 14.
	// Path: 0→R(6)→R(10)→R(12)→R(14) = "RRRR"
	assert.Equal(t, "RRRR", FindPath(5, 0, 14))
}

func TestFindPath_LeafToRoot(t *testing.T) {
	assert.Equal(t, "UUUU", FindPath(5, 14, 0))
}

func TestFindPath_Order2(t *testing.T) {
	// order=2: 3 nodes. Root=0, left=1(order 0), right=2(order 1).
	assert.Equal(t, "R", FindPath(2, 0, 2))
	assert.Equal(t, "L", FindPath(2, 0, 1))
	assert.Equal(t, "UR", FindPath(2, 1, 2))
	assert.Equal(t, "UL", FindPath(2, 2, 1))
}
