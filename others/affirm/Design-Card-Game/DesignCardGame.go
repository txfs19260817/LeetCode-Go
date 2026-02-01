package affirm

import (
	"math/rand"
	"time"
)

type CardGames struct {
	deck   []int
	scoreA int
	scoreB int
	played bool
	winner string
}

func Constructor() CardGames {
	deck := make([]int, 52)
	for i := range deck {
		deck[i] = i
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return CardGames{deck: deck}
}

func (cg *CardGames) DrawAndCompare() string {
	if cg.played {
		return cg.winner
	}

	for i := 0; i+1 < len(cg.deck); i += 2 {
		a := cg.deck[i]
		b := cg.deck[i+1]
		if a > b {
			cg.scoreA++
		} else if b > a {
			cg.scoreB++
		}
	}

	cg.played = true
	if cg.scoreA > cg.scoreB {
		cg.winner = "A"
	} else if cg.scoreB > cg.scoreA {
		cg.winner = "B"
	} else {
		cg.winner = "TIE"
	}

	return cg.winner
}
