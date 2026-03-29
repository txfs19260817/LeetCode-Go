from collections import defaultdict
import heapq
from typing import Dict, List, Tuple


Flight = Tuple[str, str, int, int]


class Solution:
    def canTransfer(self, origin: str, destination: str, flights: List[Flight]) -> bool:
        if origin == destination:
            return True

        graph: Dict[str, List[Tuple[int, int, str]]] = defaultdict(list)
        for departure_airport, arrival_airport, departure_time, arrival_time in flights:
            graph[departure_airport].append(
                (departure_time, arrival_time, arrival_airport)
            )

        earliest_arrival: Dict[str, int] = {origin: 0}
        min_heap: List[Tuple[int, str]] = [(0, origin)]

        while min_heap:
            current_time, airport = heapq.heappop(min_heap)
            if current_time != earliest_arrival.get(airport):
                continue

            if airport == destination:
                return True

            for departure_time, arrival_time, next_airport in graph.get(airport, []):
                if (
                    departure_time >= current_time
                    and arrival_time < earliest_arrival.get(next_airport, float("inf"))
                ):
                    earliest_arrival[next_airport] = arrival_time
                    heapq.heappush(min_heap, (arrival_time, next_airport))

        return False


if __name__ == "__main__":
    solver = Solution()

    assert solver.canTransfer(
        "NYC",
        "SFO",
        [
            ("NYC", "LAX", 0, 4),
            ("LAX", "SFO", 5, 7),
        ],
    ) is True

    assert solver.canTransfer(
        "NYC",
        "SFO",
        [
            ("NYC", "LAX", 0, 4),
            ("LAX", "SFO", 3, 5),
        ],
    ) is False

    assert solver.canTransfer(
        "A",
        "C",
        [
            ("A", "B", 2, 5),
            ("A", "C", 10, 12),
            ("B", "C", 5, 8),
        ],
    ) is True

    assert solver.canTransfer("SFO", "SFO", []) is True

    assert solver.canTransfer(
        "A",
        "D",
        [
            ("A", "B", 0, 3),
            ("B", "C", 3, 4),
            ("C", "D", 4, 6),
        ],
    ) is True
