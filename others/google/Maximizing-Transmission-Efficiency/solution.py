from __future__ import annotations

import heapq


def _max_product_dijkstra(n: int, start: int, graph: list[list[tuple[int, float]]]) -> list[float]:
    # Interview framing:
    # Standard Dijkstra minimizes additive distance. Here we maximize a
    # multiplicative score, but the same best-first idea still works because
    # every edge efficiency is in [0, 1], so extending a path can never improve
    # its current product.
    best = [0.0] * n
    best[start] = 1.0
    heap: list[tuple[float, int]] = [(-1.0, start)]

    while heap:
        neg_product, node = heapq.heappop(heap)
        product = -neg_product
        if product < best[node]:
            continue

        for neighbor, efficiency in graph[node]:
            new_product = product * efficiency
            if new_product > best[neighbor]:
                best[neighbor] = new_product
                heapq.heappush(heap, (-new_product, neighbor))

    # best[v] = maximum product achievable from start to v.
    return best


def max_transmission_efficiency(
    n: int,
    start: int,
    destination: int,
    graph: list[list[tuple[int, float]]],
) -> float:
    return _max_product_dijkstra(n, start, graph)[destination]


def max_transmission_efficiency_with_one_upgrade(
    n: int,
    start: int,
    destination: int,
    graph: list[list[tuple[int, float]]],
) -> float:
    # Run the base problem from both ends:
    # - best_from_start[u] = best product from start to u
    # - best_to_destination[v] = best product from v to destination
    #
    # Then try every edge u -> v as the upgraded edge. Since that one edge
    # becomes 1.0, the total product is:
    #   best_from_start[u] * 1.0 * best_to_destination[v]
    best_from_start: list[float] = _max_product_dijkstra(n, start, graph)

    reverse_graph: list[list[tuple[int, float]]] = [[] for _ in range(n)]
    for node in range(n):
        for neighbor, efficiency in graph[node]:
            reverse_graph[neighbor].append((node, efficiency))

    best_to_destination = _max_product_dijkstra(n, destination, reverse_graph)
    answer = best_from_start[destination]

    for node in range(n):
        if best_from_start[node] == 0.0:
            continue
        for neighbor, _ in graph[node]:
            if best_to_destination[neighbor] == 0.0:
                continue
            # Upgrade this directed edge to efficiency 1.0 and combine
            # the best prefix and suffix products around it.
            answer = max(answer, best_from_start[node] * best_to_destination[neighbor])

    return answer


def max_transmission_efficiency_even_edges(
    n: int,
    start: int,
    destination: int,
    graph: list[list[tuple[int, float]]],
) -> float:
    # Add one extra bit of state: parity of edge count used so far.
    # best[node][0] = best product reaching node with an even number of edges
    # best[node][1] = best product reaching node with an odd number of edges
    best = [[0.0, 0.0] for _ in range(n)]
    best[start][0] = 1.0
    heap: list[tuple[float, int, int]] = [(-1.0, start, 0)]

    while heap:
        neg_product, node, parity = heapq.heappop(heap)
        product = -neg_product
        if product < best[node][parity]:
            continue

        for neighbor, efficiency in graph[node]:
            # Every edge flips parity.
            next_parity = parity ^ 1
            new_product = product * efficiency
            if new_product > best[neighbor][next_parity]:
                best[neighbor][next_parity] = new_product
                heapq.heappush(heap, (-new_product, neighbor, next_parity))

    # We only accept paths that end with even parity.
    return best[destination][0]


def max_transmission_efficiency_with_budget(
    n: int,
    start: int,
    destination: int,
    graph: list[list[tuple[int, float, int]]],
    budget: int,
) -> float:
    # Another standard interview move: turn a constrained shortest-path problem
    # into a larger state graph. Here the state is (node, spent_cost).
    # best[node][c] = best product that reaches node using total cost exactly c.
    best = [[0.0] * (budget + 1) for _ in range(n)]
    best[start][0] = 1.0
    heap: list[tuple[float, int, int]] = [(-1.0, start, 0)]

    while heap:
        neg_product, node, spent = heapq.heappop(heap)
        product = -neg_product
        if product < best[node][spent]:
            continue

        for neighbor, efficiency, cost in graph[node]:
            next_spent = spent + cost
            if next_spent > budget:
                continue
            new_product = product * efficiency
            if new_product > best[neighbor][next_spent]:
                best[neighbor][next_spent] = new_product
                heapq.heappush(heap, (-new_product, neighbor, next_spent))

    # Destination is valid for any total cost <= budget, so take the best one.
    return max(best[destination])


if __name__ == "__main__":
    graph = [
        [(1, 0.5), (2, 0.2)],
        [(2, 0.9), (3, 0.5)],
        [(3, 0.8)],
        [],
    ]
    assert abs(max_transmission_efficiency(4, 0, 3, graph) - 0.36) < 1e-12
    assert abs(max_transmission_efficiency_with_one_upgrade(4, 0, 3, graph) - 0.8) < 1e-12
    assert abs(max_transmission_efficiency_even_edges(4, 0, 3, graph) - 0.25) < 1e-12

    budget_graph = [
        [(1, 0.5, 2), (2, 0.2, 1)],
        [(2, 0.9, 2), (3, 0.5, 2)],
        [(3, 0.8, 2)],
        [],
    ]
    assert abs(max_transmission_efficiency_with_budget(4, 0, 3, budget_graph, 5) - 0.25) < 1e-12
    assert abs(max_transmission_efficiency_with_budget(4, 0, 3, budget_graph, 6) - 0.36) < 1e-12
