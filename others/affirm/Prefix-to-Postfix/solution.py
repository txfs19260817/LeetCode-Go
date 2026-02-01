from typing import List


def prefix_to_postfix(prefix: str) -> str:
    stack: List[str] = []
    for ch in reversed(prefix):
        if ch in "+-*/":
            left = stack.pop()
            right = stack.pop()
            stack.append(left + right + ch)
        else:
            stack.append(ch)
    return stack[-1]


if __name__ == "__main__":
    tests = [
        ("+12", "12+"),
        ("-*345", "34*5-"),
        ("*+AB-CD", "AB+CD-*"),
        ("A", "A"),
        ("+A/9B", "A9B/+"),
    ]

    for prefix, expected in tests:
        result = prefix_to_postfix(prefix)
        print(prefix, "->", result)
        assert result == expected
