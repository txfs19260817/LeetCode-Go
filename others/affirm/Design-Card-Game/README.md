# Design Card Game

**Votes:** +1  
**Difficulty:** Medium  
**Topics:** Design  
**Interview Stages:** Onsite  
**Frequency:** N/A  
**Asked By:** Affirm  
**Last Reported:** 1 days ago

Create a simulation of a card game. The game uses a deck of 52 unique cards numbered from `0` to `51`. There are
two players, **A** and **B**. Players take turns drawing cards randomly from the deck. After each draw, compare the
two cards:

- If Player A's card is higher, Player A earns 1 point.
- If Player B's card is higher, Player B earns 1 point.

The game continues until all cards have been drawn and compared. The player with the higher score wins. If scores
are tied at the end, return `"TIE"`.

Implement the `CardGames` class:

- `CardGames()` Initializes the card game.
- `String drawAndCompare()` Players take turns drawing and comparing card values until all cards are drawn. Returns
  the winner or `"TIE"` if the scores are equal.

## Constraints

- Only two players, A and B.
- Only 52 unique cards, from 0 to 51.
- Each card is exactly the same except for the number.
- The total score of two players is always `52 / 2 = 26`.

## Example

**Input:**
```
["CardGames", "drawAndCompare"]
[[], []]
```

**Output:**
```
[null, "A"]
```

**Explanation:**
```
CardGames game = CardGames();
game.drawAndCompare(); // Return "A", "B", or "TIE".
1st round, A and B draw cards and compare, 50 cards left
2nd round, A and B draw cards and compare, 48 cards left
...
26th round, A and B draw cards and compare, 0 cards left
```

At the end of the game, compare the scores of Player A and Player B.

## Follow-up: N players, M cards

- Use a deck of `1..M`, shuffle, and deal `floor(M / N)` cards to each player; discard any remainder.
- Each round, every player plays one card.
- If there is a unique highest card, that player earns the sum of the cards played that round.
- If multiple players tie for the highest card, no one scores that round.
