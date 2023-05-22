package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	ticTacToe := Constructor(2)
	assert.Equal(t, 0, ticTacToe.Move(0, 0, 2))
	assert.Equal(t, 0, ticTacToe.Move(1, 1, 1))
	assert.Equal(t, 2, ticTacToe.Move(0, 1, 2))
}
