import random
from typing import List, Tuple


class CardGames:
    def __init__(self) -> None:
        self.scoreA = 0
        self.scoreB = 0
        self._played = False
        self._winner = ""

        self.deck = list(range(52))
        random.shuffle(self.deck)

    def drawAndCompare(self) -> str:
        if self._played:
            return self._winner

        for i in range(0, len(self.deck), 2):
            a = self.deck[i]
            b = self.deck[i + 1]
            if a > b:
                self.scoreA += 1
            elif b > a:
                self.scoreB += 1

        self._played = True
        if self.scoreA > self.scoreB:
            self._winner = "A"
        elif self.scoreB > self.scoreA:
            self._winner = "B"
        else:
            self._winner = "TIE"
        return self._winner


def play_n_players(
    players: List[str], m_cards: int, seed: int | None = None
) -> Tuple[List[Tuple[str, int]], List[str]]:
    """
    Returns:
      - standings: list of (playerName, score) sorted desc
      - winners: list of winner names (can be tied)
    Rule:
      - deck: 1..m_cards
      - each player gets floor(m_cards / N) cards, remainder discarded
      - each round everyone plays one card
      - if unique max, winner gets sum(cards played this round)
      - if tie on max, no one scores this round
    """
    n = len(players)
    if n <= 0 or m_cards <= 0:
        return [], []

    deck = list(range(1, m_cards + 1))
    rng = random.Random(seed)
    rng.shuffle(deck)

    per = m_cards // n
    total_used = per * n
    deck = deck[:total_used]

    hands = [deck[i * per : (i + 1) * per] for i in range(n)]
    scores = [0] * n

    for r in range(per):
        played = [hands[i][r] for i in range(n)]
        mx = max(played)
        winners = [i for i, card in enumerate(played) if card == mx]
        if len(winners) == 1:
            scores[winners[0]] += sum(played)

    standings = sorted(zip(players, scores), key=lambda x: (-x[1], x[0]))
    top_score = standings[0][1] if standings else 0
    winners = [name for name, sc in standings if sc == top_score]
    return standings, winners


if __name__ == "__main__":
    def test_two_player(seeds: List[int]) -> None:
        for seed in seeds:
            random.seed(seed)
            game = CardGames()
            result = game.drawAndCompare()
            total = game.scoreA + game.scoreB
            print(f"Seed {seed}: result={result}, A={game.scoreA}, B={game.scoreB}")
            assert total == 26
            if result == "A":
                assert game.scoreA > game.scoreB
            elif result == "B":
                assert game.scoreB > game.scoreA
            else:
                assert game.scoreA == game.scoreB

    def test_n_players() -> None:
        standings, winners = play_n_players(["A", "B", "C"], 56, seed=42)
        print("Winners:", winners)
        print("Standings:")
        for name, score in standings:
            print(f"  {name}: {score}")
        assert standings
        assert winners
        top_score = standings[0][1]
        assert all(sc <= top_score for _, sc in standings)

    test_two_player([1, 2, 3, 4, 5])
    test_n_players()
