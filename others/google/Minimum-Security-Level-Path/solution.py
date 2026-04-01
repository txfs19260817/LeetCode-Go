from __future__ import annotations

import heapq


def min_security_level(
    n: int,
    start: int,
    destination: int,
    graph: list[list[tuple[int, int]]],
) -> int:
    if start == destination:
        return 0

    # best[node] = minimum possible "maximum edge security seen so far"
    # among all paths from start to node.
    best = [float("inf")] * n
    best[start] = 0
    heap: list[tuple[int, int]] = [(0, start)]

    while heap:
        current_security, node = heapq.heappop(heap)
        if node == destination:
            return current_security
        if current_security > best[node]:
            continue

        for neighbor, edge_security in graph[node]:
            # Extending the path updates the path cost by taking the larger
            # of the current requirement and the new edge's security level.
            next_security = max(current_security, edge_security)
            if next_security < best[neighbor]:
                best[neighbor] = next_security
                heapq.heappush(heap, (next_security, neighbor))

    return -1


if __name__ == "__main__":
    graph = [
        [(1, 3), (2, 1)],
        [(3, 5)],
        [(3, 2)],
        [],
    ]
    assert min_security_level(4, 0, 3, graph) == 2
    assert min_security_level(4, 0, 2, graph) == 1
    assert min_security_level(4, 0, 0, graph) == 0
    assert min_security_level(4, 3, 0, graph) == -1
