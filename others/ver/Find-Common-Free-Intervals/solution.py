from __future__ import annotations

from typing import List, Tuple

Interval = Tuple[int, int]


def find_common_free_intervals(
    calendars: List[List[Interval]], day_start: int, day_end: int
) -> List[Interval]:
    if day_start >= day_end:
        raise ValueError("day_start must be smaller than day_end")

    busy: List[Interval] = []
    for person in calendars:
        for s, e in person:
            # Clip each busy interval into [day_start, day_end),
            # then keep only effective intervals.
            clipped_s = max(s, day_start)
            clipped_e = min(e, day_end)
            if clipped_s < clipped_e:
                busy.append((clipped_s, clipped_e))
    if not busy:
        return [(day_start, day_end)]

    busy.sort()

    merged: List[Interval] = []
    cur_s, cur_e = busy[0]
    for s, e in busy[1:]:
        # Merge overlap or touching intervals under [start, end) semantics.
        if s <= cur_e:
            cur_e = max(cur_e, e)
        else:
            merged.append((cur_s, cur_e))
            cur_s, cur_e = s, e
    merged.append((cur_s, cur_e))

    free: List[Interval] = []
    cursor = day_start
    for s, e in merged:
        # Complement gaps between merged busy intervals.
        if cursor < s:
            free.append((cursor, s))
        cursor = max(cursor, e)

    if cursor < day_end:
        free.append((cursor, day_end))

    return free


if __name__ == "__main__":
    calendars_1 = [
        [(9, 10), (12, 13), (16, 18)],
        [(8, 9), (11, 12), (14, 17)],
        [(9, 10), (13, 14), (15, 16)],
    ]
    assert find_common_free_intervals(calendars_1, 8, 18) == [(10, 11)]

    calendars_2 = [
        [(9, 12)],
        [(13, 15)],
        [(10, 11)],
    ]
    assert find_common_free_intervals(calendars_2, 8, 18) == [
        (8, 9),
        (12, 13),
        (15, 18),
    ]

    calendars_3 = [[], []]
    assert find_common_free_intervals(calendars_3, 9, 17) == [(9, 17)]

    try:
        find_common_free_intervals([[]], 10, 10)
        raise AssertionError("Expected ValueError")
    except ValueError:
        pass

    print("All assertions passed.")
