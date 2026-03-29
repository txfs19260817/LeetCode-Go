from __future__ import annotations

from collections import defaultdict, deque
import heapq


def recommend_movies(
    movie: str,
    ratings: dict[str, int],
    similarities: list[tuple[str, str]],
    n: int,
) -> list[str]:
    graph: dict[str, list[str]] = defaultdict(list)
    for left, right in similarities:
        graph[left].append(right)
        graph[right].append(left)

    seen = {movie}
    queue = deque([movie])
    heap: list[tuple[int, str]] = []

    while queue:
        current = queue.popleft()
        for neighbor in graph[current]:
            if neighbor in seen:
                continue
            seen.add(neighbor)
            queue.append(neighbor)
            heapq.heappush(heap, (-ratings[neighbor], neighbor))

    recommendations: list[str] = []
    while heap and len(recommendations) < n:
        _, name = heapq.heappop(heap)
        recommendations.append(name)

    return recommendations


if __name__ == "__main__":
    assert recommend_movies(
        "A",
        {"A": 6, "B": 7, "C": 8, "D": 9},
        [("A", "B"), ("B", "C")],
        1,
    ) == ["C"]

    assert recommend_movies(
        "A",
        {"A": 6, "B": 7, "C": 8, "D": 9},
        [("A", "B"), ("B", "C")],
        2,
    ) == ["C", "B"]

    assert recommend_movies(
        "A",
        {"A": 6, "B": 7, "C": 7, "D": 9},
        [("A", "B"), ("A", "C"), ("D", "C")],
        3,
    ) == ["D", "B", "C"]

    assert recommend_movies(
        "A",
        {"A": 6, "B": 7, "C": 8, "D": 9},
        [],
        2,
    ) == []
