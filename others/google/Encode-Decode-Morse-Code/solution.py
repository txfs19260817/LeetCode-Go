def encode(message: str, char_to_morse: dict[str, str]) -> str:
    return "".join([char_to_morse[c] for c in message])


def decode(morse: str, morse_to_char: dict[str, str]) -> list[str]:
    # Same DP shape as LeetCode 140: dp[i] holds every decoded string for
    # the prefix morse[:i].
    dp: list[list[str]] = [[] for _ in range(len(morse) + 1)]
    dp[0] = [""]

    for i in range(1, len(dp)):
        for m, ch in morse_to_char.items():
            m_len = len(m)
            if i < m_len or m != morse[i - m_len: i]:
                continue
            for prev in dp[i - m_len]:
                dp[i].append(prev + ch if prev else ch)

        dp[i].sort()

    return dp[len(morse)]


if __name__ == "__main__":
    char_to_morse = {
        "A": ".",
        "B": "-",
        "C": ".-",
    }
    morse_to_char = {
        ".": "A",
        "-": "B",
        ".-": "C",
    }

    assert encode("AB", char_to_morse) == ".-"
    assert encode("C", char_to_morse) == ".-"
    assert decode(".-", morse_to_char) == ["AB", "C"]
    assert decode(".-.", morse_to_char) == ["ABA", "CA"]
    assert decode("", morse_to_char) == [""]

    alphabet_char_to_morse = {
        "A": ".-",
        "B": "-...",
        "C": "-.-.",
        "D": "-..",
        "E": ".",
        "F": "..-.",
        "G": "--.",
        "H": "....",
        "I": "..",
        "J": ".---",
        "K": "-.-",
        "L": ".-..",
        "M": "--",
        "N": "-.",
        "O": "---",
        "P": ".--.",
        "Q": "--.-",
        "R": ".-.",
        "S": "...",
        "T": "-",
        "U": "..-",
        "V": "...-",
        "W": ".--",
        "X": "-..-",
        "Y": "-.--",
        "Z": "--..",
    }
    alphabet_morse_to_char = {
        code: ch for ch, code in alphabet_char_to_morse.items()
    }

    for ch, code in alphabet_char_to_morse.items():
        assert encode(ch, alphabet_char_to_morse) == code
        assert ch in decode(code, alphabet_morse_to_char)
