from typing import List


def find_needle_in_chunked_strings(chunks: List[str], needle: str) -> tuple[int, int]:
    """Return (chunk_index, offset) for the first occurrence of needle."""
    if needle == "":
        return 0, 0

    for i in range(len(chunks)):
        for j in range(len(chunks[i])):
            x, y, k = i, j, 0
            while k < len(needle) and x < len(chunks):
                if needle[k] != chunks[x][y]:
                    break
                k += 1
                if k == len(needle):
                    return i, j
                y += 1
                while x < len(chunks) and y >= len(chunks[x]):
                    x += 1
                    y = 0

    return -1, -1


if __name__ == "__main__":
    assert find_needle_in_chunked_strings(["a", "bcd", "ea", "de"], "de") == (1, 2)
    assert find_needle_in_chunked_strings(["hello", "world"], "llo") == (0, 2)
    assert find_needle_in_chunked_strings(["ab", "", "c", "de"], "bcde") == (0, 1)
    assert find_needle_in_chunked_strings(["ab", "cd", "abc"], "abc") == (0, 0)
    assert find_needle_in_chunked_strings(["", "", "de"], "de") == (2, 0)
    assert find_needle_in_chunked_strings(["ab", "c"], "abcd") == (-1, -1)
    assert find_needle_in_chunked_strings(["ab", "cd"], "ef") == (-1, -1)
    assert find_needle_in_chunked_strings(["ab", "cd"], "") == (0, 0)
