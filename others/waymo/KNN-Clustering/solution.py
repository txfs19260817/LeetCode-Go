from dataclasses import dataclass


@dataclass(frozen=True)
class Point:
    x: float
    y: float


@dataclass(frozen=True)
class LabeledPoint:
    point: Point
    label: int


def assign_clusters(train: list[LabeledPoint], queries: list[Point], k: int) -> list[int]:
    if not train or k <= 0:
        return [-1] * len(queries)
    k = min(k, len(train))

    out: list[int] = []
    for query in queries:
        neighbors = [(distance_squared(query, sample.point), sample.label) for sample in train]
        neighbors.sort(key=lambda item: (item[0], item[1]))

        stats: list[list[float | int]] = []
        for dist2, label in neighbors[:k]:
            found = False
            for stat in stats:
                if stat[0] == label:
                    stat[1] += 1
                    stat[2] += dist2
                    found = True
                    break
            if not found:
                stats.append([label, 1, dist2])

        best_label, best_votes, best_dist = stats[0]
        for label, votes, total_dist in stats[1:]:
            if votes > best_votes:
                best_label, best_votes, best_dist = label, votes, total_dist
            elif votes == best_votes:
                if total_dist < best_dist:
                    best_label, best_votes, best_dist = label, votes, total_dist
                elif total_dist == best_dist and label < best_label:
                    best_label, best_votes, best_dist = label, votes, total_dist
        out.append(int(best_label))
    return out


def distance_squared(a: Point, b: Point) -> float:
    dx = a.x - b.x
    dy = a.y - b.y
    return dx * dx + dy * dy


if __name__ == "__main__":
    train = [
        LabeledPoint(Point(0, 0), 0),
        LabeledPoint(Point(0, 1), 0),
        LabeledPoint(Point(10, 10), 1),
        LabeledPoint(Point(10, 11), 1),
    ]
    queries = [Point(0.2, 0.1), Point(9.8, 10.2), Point(5, 5)]
    assert assign_clusters(train, queries, 3) == [0, 1, 0]

    train2 = [LabeledPoint(Point(-1, 0), 5), LabeledPoint(Point(1, 0), 3)]
    assert assign_clusters(train2, [Point(0, 0)], 2) == [3]

    assert assign_clusters([], [Point(0, 0), Point(1, 1)], 3) == [-1, -1]
