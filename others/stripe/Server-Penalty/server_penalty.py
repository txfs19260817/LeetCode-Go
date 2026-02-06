# https://www.1point3acres.com/bbs/thread-1092448-1-1.html
def parse_input(inputString):
    route_dict = {}

    routes = inputString.split(",")
    for route in routes:
        source, target, method, cost = route.split(":")
        if source not in route_dict:
            route_dict[source] = []
        route_dict[source].append((target, method, int(cost)))
    return route_dict


def find_direct_route(inputString, sourceCountry, targetCountry, method):
    route_dict = parse_input(inputString)
    if sourceCountry not in route_dict:
        return "No routes from source country"

    for target, method, cost in route_dict[sourceCountry]:
        if target == targetCountry and method == method:
            return cost

    return "No valid route found"


inputString = "US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10,AU:JP:DHL:20"
sourceCountry = "US"
targetCountry = "UK"

direct_route_info = find_direct_route(inputString, sourceCountry, targetCountry, "UPS")
print(direct_route_info)


def find_route(inputString, sourceCountry, targetCountry):
    route_dict = parse_input(inputString)

    if sourceCountry not in route_dict:
        return "No routes from source country"

    for intermediate_target, first_method, first_cost in route_dict[sourceCountry]:
        if intermediate_target in route_dict:
            for final_target, second_method, second_cost in route_dict[
                intermediate_target
            ]:
                if final_target == targetCountry:
                    return {
                        "route": f"{sourceCountry} -> {intermediate_target} -> {targetCountry}",
                        "method": f"{first_method} -> {second_method}",
                        "cost": first_cost + second_cost,
                    }

    return "No valid route found"


inputString = "US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10,AU:JP:DHL:20"
sourceCountry = "US"
targetCountry = "CA"
route_info = find_route(inputString, sourceCountry, targetCountry)
print(route_info)


def find_cheapest_route(inputString, sourceCountry, targetCountry):
    route_dict = parse_input(inputString)

    if sourceCountry not in route_dict:
        return "No routes from source country"

    #######
    cheapest_route = None
    min_cost = sys.maxsize
    #######

    for intermediate_target, first_method, first_cost in route_dict[sourceCountry]:
        if intermediate_target in route_dict:
            for final_target, second_method, second_cost in route_dict[
                intermediate_target
            ]:
                if final_target == targetCountry:
                    ######
                    total_cost = first_cost + second_cost
                    if total_cost < min_cost:
                        min_cost = total_cost
                        cheapest_route = {
                            "route": f"{sourceCountry} -> {intermediate_target} -> {targetCountry}",
                            "method": f"{first_method} -> {second_method}",
                            "cost": total_cost,
                        }
                    ######

    return cheapest_route if cheapest_route else "No valid route found"


import heapq


def dijkstra(inputString, sourceCountry, targetCountry):
    graph = parse_input(inputString)
    queue = [(0, sourceCountry, [], [])]  # (cost, current country, path, methods)
    visited = {sourceCountry: 0}

    while queue:
        cost, country, path, methods = heapq.heappop(queue)

        if country == targetCountry:
            return {
                "route": " -> ".join(path + [country]),
                "method": " -> ".join(methods),
                "cost": cost,
            }

        if country not in graph:
            continue

        for neighbor, method, travel_cost in graph[country]:
            new_cost = cost + travel_cost

            if neighbor not in visited or new_cost < visited[neighbor]:
                visited[neighbor] = new_cost
                heapq.heappush(
                    queue, (new_cost, neighbor, path + [country], methods + [method])
                )

    return "No valid route found"


inputString = (
    "US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10,AU:JP:DHL:20,US:JP:DHL:50,JP:CA:DHL:15"
)
sourceCountry = "US"
targetCountry = "CA"

cheapest_route_info = dijkstra(inputString, sourceCountry, targetCountry)
print(cheapest_route_info)
