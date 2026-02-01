package affirm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeAWinsDeck() []int {
	deck := make([]int, 0, 52)
	hi, lo := 51, 0
	for lo < hi {
		deck = append(deck, hi, lo)
		hi--
		lo++
	}
	return deck
}

func makeBWinsDeck() []int {
	deck := make([]int, 0, 52)
	hi, lo := 51, 0
	for lo < hi {
		deck = append(deck, lo, hi)
		hi--
		lo++
	}
	return deck
}

func makeTieDeck() []int {
	deck := make([]int, 0, 52)
	hi, lo := 51, 0
	for i := 0; i < 13; i++ {
		deck = append(deck, hi, lo)
		hi--
		lo++
	}
	for i := 0; i < 13; i++ {
		deck = append(deck, lo, hi)
		hi--
		lo++
	}
	return deck
}

func TestCardGamesDrawAndCompare(t *testing.T) {
	t.Run("A wins all rounds", func(t *testing.T) {
		game := Constructor()
		game.deck = makeAWinsDeck()

		result := game.DrawAndCompare()

		assert.Equal(t, "A", result)
		assert.Equal(t, 26, game.scoreA)
		assert.Equal(t, 0, game.scoreB)
	})

	t.Run("B wins all rounds", func(t *testing.T) {
		game := Constructor()
		game.deck = makeBWinsDeck()

		result := game.DrawAndCompare()

		assert.Equal(t, "B", result)
		assert.Equal(t, 0, game.scoreA)
		assert.Equal(t, 26, game.scoreB)
	})

	t.Run("Tie after all rounds", func(t *testing.T) {
		game := Constructor()
		game.deck = makeTieDeck()

		result := game.DrawAndCompare()

		assert.Equal(t, "TIE", result)
		assert.Equal(t, 13, game.scoreA)
		assert.Equal(t, 13, game.scoreB)
	})

	t.Run("Random deck invariants", func(t *testing.T) {
		game := Constructor()
		result := game.DrawAndCompare()

		assert.Equal(t, 26, game.scoreA+game.scoreB)
		switch result {
		case "A":
			assert.Greater(t, game.scoreA, game.scoreB)
		case "B":
			assert.Greater(t, game.scoreB, game.scoreA)
		case "TIE":
			assert.Equal(t, game.scoreA, game.scoreB)
		default:
			t.Fatalf("unexpected result: %s", result)
		}
	})

	t.Run("Idempotent draw", func(t *testing.T) {
		game := Constructor()
		game.deck = makeAWinsDeck()

		first := game.DrawAndCompare()
		second := game.DrawAndCompare()

		assert.Equal(t, first, second)
		assert.Equal(t, 26, game.scoreA)
		assert.Equal(t, 0, game.scoreB)
	})
}
