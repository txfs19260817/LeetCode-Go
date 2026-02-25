from collections import Counter
from typing import List

GREEN = 2
YELLOW = 1
GRAY = 0


def score_guess(secret: str, guess: str) -> List[int]:
    """Return WordGuess feedback for guess against secret."""
    if not isinstance(secret, str) or not isinstance(guess, str):
        raise ValueError("secret and guess must be strings")
    if not secret or len(secret) != len(guess):
        raise ValueError("secret and guess must be non-empty and same length")

    n = len(secret)
    result = [GRAY] * n
    remaining_secret = Counter()

    # Pass 1: lock exact matches; collect only unmatched secret letters.
    for i in range(n):
        if guess[i] == secret[i]:
            result[i] = GREEN
        else:
            remaining_secret[secret[i]] += 1

    # Pass 2: award YELLOW only if unmatched inventory still has that letter.
    for i in range(n):
        if result[i] == GREEN:
            continue
        ch = guess[i]
        if remaining_secret[ch] > 0:
            result[i] = YELLOW
            remaining_secret[ch] -= 1

    return result


if __name__ == "__main__":
    assert score_guess("hello", "abcde") == [0, 0, 0, 0, 1]
    assert score_guess("hello", "hable") == [2, 0, 0, 2, 1]
    assert score_guess("hello", "eabcd") == [1, 0, 0, 0, 0]

    # duplicate letters
    assert score_guess("aabbc", "caaba") == [1, 2, 1, 2, 0]

    # invalid input
    try:
        score_guess("hello", "hi")
        raise AssertionError("Expected ValueError for different lengths")
    except ValueError:
        pass

    print("All assertions passed.")
