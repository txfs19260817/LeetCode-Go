def russian_peasant_multiply_iterative(a: int, b: int) -> int:
    a, b, negative = _normalize_signs(a, b)

    result = 0
    while a > 0:
        if a & 1:
            result += b
        a >>= 1
        b <<= 1

    return -result if negative else result


def russian_peasant_multiply_recursive(a: int, b: int) -> int:
    a, b, negative = _normalize_signs(a, b)
    result = _multiply_recursive_positive(a, b)
    return -result if negative else result


def _multiply_recursive_positive(a: int, b: int) -> int:
    if a == 0:
        return 0
    if a & 1:
        return b + _multiply_recursive_positive(a >> 1, b << 1)
    return _multiply_recursive_positive(a >> 1, b << 1)


def _normalize_signs(a: int, b: int) -> tuple[int, int, bool]:
    negative = False
    if a < 0:
        a = -a
        negative = not negative
    if b < 0:
        b = -b
        negative = not negative
    return a, b, negative


if __name__ == "__main__":
    cases = [
        (13, 12, 156),
        (0, 7, 0),
        (37, 0, 0),
        (-13, 12, -156),
        (13, -12, -156),
        (-13, -12, 156),
        (15, 15, 225),
        (64, 3, 192),
    ]

    for a, b, want in cases:
        assert russian_peasant_multiply_iterative(a, b) == want
        assert russian_peasant_multiply_recursive(a, b) == want
