from __future__ import annotations

from typing import List, Sequence


Interval = List[float]
Sample = Sequence[float]


def motion_intervals_single(
    samples: Sequence[Sample],
    threshold: float,
    inclusive: bool = True,
) -> List[Interval]:
    """
    Return motion intervals for one camera as [start_ts, end_ts].
    By default, level >= threshold is considered motion.
    """
    if not samples:
        return []

    # Normalize to timestamp order so interval building is deterministic.
    sorted_samples = sorted(samples, key=lambda x: x[0])

    def is_motion(level: float) -> bool:
        return level >= threshold if inclusive else level > threshold

    intervals: List[Interval] = []
    in_motion = False
    start_ts = 0.0
    end_ts = 0.0

    for ts, level in sorted_samples:
        # State machine:
        # - Enter interval when motion starts.
        # - Extend interval while motion continues.
        # - Flush interval once motion drops below threshold.
        if is_motion(level):
            if not in_motion:
                in_motion = True
                start_ts = ts
            end_ts = ts
        elif in_motion:
            intervals.append([start_ts, end_ts])
            in_motion = False

    if in_motion:
        intervals.append([start_ts, end_ts])

    return intervals


def _intersect_two(a: Sequence[Interval], b: Sequence[Interval]) -> List[Interval]:
    i = 0
    j = 0
    merged: List[Interval] = []

    # Two-pointer sweep over two sorted interval lists.
    while i < len(a) and j < len(b):
        start = max(a[i][0], b[j][0])
        end = min(a[i][1], b[j][1])
        if start <= end:
            merged.append([start, end])

        # Advance the interval that ends first.
        if a[i][1] < b[j][1]:
            i += 1
        else:
            j += 1

    return merged


def motion_intervals_all_cameras(
    cameras: Sequence[Sequence[Sample]],
    threshold: float,
    inclusive: bool = True,
) -> List[Interval]:
    """
    Return time intervals where all cameras are in motion.
    """
    if not cameras:
        return []

    all_intervals = [
        motion_intervals_single(camera_samples, threshold, inclusive)
        for camera_samples in cameras
    ]
    # Intersect smaller lists first to reduce intermediate overlap size.
    all_intervals.sort(key=len)

    overlap = all_intervals[0]
    for intervals in all_intervals[1:]:
        overlap = _intersect_two(overlap, intervals)
        if not overlap:
            break

    return overlap


if __name__ == "__main__":
    single_camera = [
        [1, 0.4],
        [5, 0.2],
        [11, 0.9],
        [15, 0.9],
        [17, 0.8],
        [25, 0.5],
        [27, 0.8],
        [36, 0.9],
    ]
    assert motion_intervals_single(single_camera, 0.8) == [[11, 17], [27, 36]]
    assert motion_intervals_single(single_camera, 0.8, inclusive=False) == [[11, 15], [36, 36]]

    cam1 = single_camera
    cam2 = [
        [2, 0.7],
        [10, 0.8],
        [12, 0.9],
        [16, 0.85],
        [20, 0.1],
        [28, 0.95],
        [35, 0.82],
        [40, 0.1],
    ]
    cam3 = [
        [9, 0.81],
        [14, 0.8],
        [18, 0.2],
        [26, 0.9],
        [30, 0.8],
        [34, 0.7],
    ]

    assert motion_intervals_all_cameras([cam1, cam2, cam3], 0.8) == [[11, 14], [28, 30]]
    assert motion_intervals_all_cameras([], 0.8) == []

    print("All assertions passed.")
