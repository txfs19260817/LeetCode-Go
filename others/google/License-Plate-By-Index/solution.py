from __future__ import annotations


LETTER_ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
LETTER_BASE = len(LETTER_ALPHABET)
PLATE_LENGTH = 5


def _block_size(letter_count: int) -> int:
    # With letter_count leading letters fixed-width, the remaining
    # positions are decimal digits, so this block has:
    #   26^letter_count * 10^(5 - letter_count)
    # plates in total.
    return (LETTER_BASE**letter_count) * (10 ** (PLATE_LENGTH - letter_count))


MAX_INDEX = sum(_block_size(letter_count) for letter_count in range(PLATE_LENGTH + 1)) - 1


def _encode_letters(value: int, width: int) -> str:
    # Encode the leading letter prefix in base 26 using A-Z.
    chars = ["A"] * width
    for pos in range(width - 1, -1, -1):
        value, digit = divmod(value, LETTER_BASE)
        chars[pos] = LETTER_ALPHABET[digit]
    return "".join(chars)


def f(index: int) -> str:
    if not 0 <= index <= MAX_INDEX:
        raise ValueError(f"index must be in [0, {MAX_INDEX}]")

    for letter_count in range(PLATE_LENGTH + 1):
        block_size = _block_size(letter_count)
        if index >= block_size:
            # Skip all plates from smaller-format blocks until we find
            # the block that contains this index.
            index -= block_size
            continue

        digit_count = PLATE_LENGTH - letter_count
        divisor = 10**digit_count

        # Inside the chosen block:
        # - the high part selects the letter prefix
        # - the low part selects the zero-padded decimal suffix
        letter_value, digit_value = divmod(index, divisor)
        digit_part = "" if digit_count == 0 else str(digit_value).zfill(digit_count)
        return _encode_letters(letter_value, letter_count) + digit_part

    raise AssertionError("unreachable")


if __name__ == "__main__":
    assert f(0) == "00000"
    assert f(1) == "00001"
    assert f(99999) == "99999"
    assert f(100000) == "A0000"
    assert f(359999) == "Z9999"
    assert f(360000) == "AA000"
    assert f(MAX_INDEX) == "ZZZZZ"
