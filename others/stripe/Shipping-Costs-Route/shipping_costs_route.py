import heapq
from typing import Dict, List, Optional, Tuple

# https://www.1point3acres.com/bbs/thread-1092448-1-1.html

# A single directed route: (target_country, shipping_method, cost).
Route = Tuple[str, str, int]
# A normalized summary used by Part 2/3/4 outputs.
RouteSummary = Dict[str, object]


def parse_routes(route_string: str) -> Dict[str, List[Route]]:
    """Parse the input string into an adjacency list keyed by source country."""
    routes_by_source: Dict[str, List[Route]] = {}
    if not route_string:
        return routes_by_source

    for raw in route_string.split(","):
        raw = raw.strip()
        if not raw:
            continue
        source, target, method, cost_str = [part.strip() for part in raw.split(":")]
        # Store all outgoing routes from the same source.
        routes_by_source.setdefault(source, []).append((target, method, int(cost_str)))

    return routes_by_source


def _route_summary(path: List[str], methods: List[str], cost: int) -> RouteSummary:
    """Format path/methods/cost into a consistent summary dict."""
    return {"route": " -> ".join(path), "method": " -> ".join(methods), "cost": cost}


def _first_leg_to_target(
    route_dict: Dict[str, List[Route]], mids: List[str], target_country: str
) -> Dict[str, Tuple[str, int]]:
    """
    For each mid country, cache the first (method, cost) route that reaches target_country.
    This preserves iteration order while avoiding repeated rescans of the same mid adjacency.
    """
    first_leg: Dict[str, Tuple[str, int]] = {}
    for mid in dict.fromkeys(mids):
        for target, method, cost in route_dict.get(mid, []):
            if target == target_country:
                first_leg[mid] = (method, cost)
                break
    return first_leg


def _cheapest_leg_to_target(
    route_dict: Dict[str, List[Route]], mids: List[str], target_country: str
) -> Dict[str, Tuple[str, int]]:
    """
    For each mid country, cache the cheapest (method, cost) route to target_country.
    Ties keep the first encountered route to preserve deterministic behavior.
    """
    cheapest_leg: Dict[str, Tuple[str, int]] = {}
    for mid in dict.fromkeys(mids):
        best: Optional[Tuple[str, int]] = None
        for target, method, cost in route_dict.get(mid, []):
            if target != target_country:
                continue
            if best is None or cost < best[1]:
                best = (method, cost)
        if best is not None:
            cheapest_leg[mid] = best
    return cheapest_leg


# Part 1: direct route with specific method
def direct_cost(
    route_string: str, source_country: str, target_country: str, method: str
) -> Optional[int]:
    """
    Return the cost if a direct route exists with the specified method.
    If no such route exists, return None.
    """
    route_dict = parse_routes(route_string)

    # Scan all outgoing edges from the source for an exact match.
    for target, route_method, cost in route_dict.get(source_country, []):
        if target == target_country and route_method == method:
            return cost

    return None


# Part 2: any route with at most one intermediate country (ignore method selection)
def find_route_with_one_stop(
    route_string: str, source_country: str, target_country: str
) -> Optional[RouteSummary]:
    """
    Find any valid route with 0 or 1 stop. Methods are returned for the path found.
    Direct routes are preferred because they are checked first.
    """
    route_dict = parse_routes(route_string)
    source_routes = route_dict.get(source_country, [])

    # First check direct routes (0 stops).
    for target, method, cost in source_routes:
        if target == target_country:
            return _route_summary([source_country, target], [method], cost)

    # Then check 1-stop routes (source -> mid -> target).
    first_leg_map = _first_leg_to_target(
        route_dict,
        [mid for mid, _, _ in source_routes],
        target_country,
    )
    for mid, first_method, first_cost in source_routes:
        second_leg = first_leg_map.get(mid)
        if second_leg is None:
            continue
        second_method, second_cost = second_leg
        return _route_summary(
            [source_country, mid, target_country],
            [first_method, second_method],
            first_cost + second_cost,
        )

    return None


# Part 3: cheapest route with at most one intermediate country
def cheapest_route_with_one_stop(
    route_string: str, source_country: str, target_country: str
) -> Optional[RouteSummary]:
    """
    Among direct and 1-stop routes, return the minimum total cost path.
    If multiple routes have the same cost, the first encountered is returned.
    """
    route_dict = parse_routes(route_string)
    source_routes = route_dict.get(source_country, [])
    best_route: Optional[RouteSummary] = None

    # Direct routes: initialize best_route if possible.
    for target, method, cost in source_routes:
        if target == target_country:
            candidate = _route_summary([source_country, target], [method], cost)
            if best_route is None or candidate["cost"] < best_route["cost"]:
                best_route = candidate

    # One-stop routes: combine each source->mid edge with the cheapest mid->target edge.
    cheapest_leg_map = _cheapest_leg_to_target(
        route_dict,
        [mid for mid, _, _ in source_routes],
        target_country,
    )
    for mid, first_method, first_cost in source_routes:
        second_leg = cheapest_leg_map.get(mid)
        if second_leg is None:
            continue
        second_method, second_cost = second_leg
        total_cost = first_cost + second_cost
        candidate = _route_summary(
            [source_country, mid, target_country],
            [first_method, second_method],
            total_cost,
        )
        if best_route is None or total_cost < best_route["cost"]:
            best_route = candidate

    return best_route


# Part 4: cheapest route with any number of hops
def cheapest_route_any_hops(
    route_string: str, source_country: str, target_country: str
) -> Optional[RouteSummary]:
    """
    Part 4: find the minimum-cost route with any number of hops.

    Dijkstra is appropriate because all edge costs are non-negative.
    We also track the predecessor and method used to reconstruct the path.
    """

    # Build adjacency list from the input string.
    graph = parse_routes(route_string)

    # If source and target are the same, the cheapest route is cost 0
    # with a single-node path and no methods.
    if source_country == target_country:
        return _route_summary([source_country], [], 0)

    # Min-heap storing (total_cost_so_far, current_country).
    heap: List[Tuple[int, str]] = [(0, source_country)]

    # Best known cost to reach each country.
    distances: Dict[str, int] = {source_country: 0}

    # Predecessor map to reconstruct the cheapest path.
    # Each entry: country -> (previous_country, method_used).
    previous: Dict[str, Tuple[Optional[str], Optional[str]]] = {
        source_country: (None, None)
    }

    while heap:
        current_cost, current_country = heapq.heappop(heap)

        # Skip outdated heap entries that are no longer optimal.
        if current_cost > distances.get(current_country, current_cost):
            continue

        # With Dijkstra, the first time we pop the target, we have the minimum cost.
        if current_country == target_country:
            break

        # Explore outgoing edges from the current country.
        for neighbor, method, travel_cost in graph.get(current_country, []):
            new_cost = current_cost + travel_cost

            # Relax the edge if this path is cheaper.
            if new_cost < distances.get(neighbor, float("inf")):
                distances[neighbor] = new_cost
                previous[neighbor] = (current_country, method)
                heapq.heappush(heap, (new_cost, neighbor))

    # If target was never reached, no valid route exists.
    if target_country not in distances:
        return None

    # Reconstruct the path and methods by walking the predecessor chain.
    path: List[str] = []
    methods: List[str] = []
    current = target_country
    while current is not None:
        path.append(current)
        prev_country, method = previous.get(current, (None, None))
        if prev_country is not None and method is not None:
            methods.append(method)
        current = prev_country

    # We reconstructed from target -> source, so reverse to get source -> target.
    path.reverse()
    methods.reverse()

    return _route_summary(path, methods, distances[target_country])


if __name__ == "__main__":
    routes = (
        "US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10,AU:JP:DHL:20,US:JP:DHL:50,"
        "JP:CA:DHL:15"
    )

    assert direct_cost(routes, "US", "UK", "UPS") == 4
    assert direct_cost(routes, "US", "UK", "FedEx") is None

    route = find_route_with_one_stop(routes, "US", "CA")
    assert route == {
        "route": "US -> UK -> CA",
        "method": "UPS -> FedEx",
        "cost": 14,
    }

    cheapest = cheapest_route_with_one_stop(routes, "US", "CA")
    assert cheapest == {
        "route": "US -> UK -> CA",
        "method": "UPS -> FedEx",
        "cost": 14,
    }

    cheapest_any = cheapest_route_any_hops(routes, "US", "CA")
    assert cheapest_any == {
        "route": "US -> UK -> CA",
        "method": "UPS -> FedEx",
        "cost": 14,
    }
